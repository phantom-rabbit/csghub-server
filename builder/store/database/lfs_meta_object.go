package database

import (
	"context"
	"fmt"
	"time"
)

type lfsMetaObjectStoreImpl struct {
	db *DB
}

type LfsMetaObjectStore interface {
	FindByOID(ctx context.Context, RepoId int64, Oid string) (*LfsMetaObject, error)
	FindByRepoID(ctx context.Context, repoID int64) ([]LfsMetaObject, error)
	Create(ctx context.Context, lfsObj LfsMetaObject) (*LfsMetaObject, error)
	RemoveByOid(ctx context.Context, oid string, repoID int64) error
	UpdateOrCreate(ctx context.Context, input LfsMetaObject) (*LfsMetaObject, error)
	BulkUpdateOrCreate(ctx context.Context, input []LfsMetaObject) error
}

func NewLfsMetaObjectStore() LfsMetaObjectStore {
	return &lfsMetaObjectStoreImpl{
		db: defaultDB,
	}
}

func NewLfsMetaObjectStoreWithDB(db *DB) LfsMetaObjectStore {
	return &lfsMetaObjectStoreImpl{
		db: db,
	}
}

type LfsMetaObject struct {
	ID           int64      `bun:",pk,autoincrement" json:"user_id"`
	Oid          string     `bun:",notnull" json:"oid"`
	Size         int64      `bun:",notnull" json:"size"`
	RepositoryID int64      `bun:",notnull" json:"repository_id"`
	Repository   Repository `bun:"rel:belongs-to,join:repository_id=id" json:"repository"`
	Existing     bool       `bun:",notnull" json:"existing"`
	times
}

func (s *lfsMetaObjectStoreImpl) FindByOID(ctx context.Context, RepoId int64, Oid string) (*LfsMetaObject, error) {
	var lfsMetaObject LfsMetaObject
	err := s.db.Operator.Core.NewSelect().
		Model(&lfsMetaObject).
		Where("oid=? and repository_id=?", Oid, RepoId).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &lfsMetaObject, nil
}

func (s *lfsMetaObjectStoreImpl) FindByRepoID(ctx context.Context, repoID int64) ([]LfsMetaObject, error) {
	var lfsMetaObjects []LfsMetaObject
	err := s.db.Operator.Core.NewSelect().
		Model(&lfsMetaObjects).
		Where("repository_id=?", repoID).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return lfsMetaObjects, nil
}

func (s *lfsMetaObjectStoreImpl) Create(ctx context.Context, lfsObj LfsMetaObject) (*LfsMetaObject, error) {
	err := s.db.Operator.Core.NewInsert().
		Model(&lfsObj).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &lfsObj, nil
}

func (s *lfsMetaObjectStoreImpl) RemoveByOid(ctx context.Context, oid string, repoID int64) error {
	_, err := s.db.Operator.Core.NewDelete().
		Model(&LfsMetaObject{}).
		Where("oid = ? and repository_id= ?", oid, repoID).
		Exec(ctx)

	return err
}

func (s *lfsMetaObjectStoreImpl) UpdateOrCreate(ctx context.Context, input LfsMetaObject) (*LfsMetaObject, error) {
	input.UpdatedAt = time.Now()
	_, err := s.db.Core.NewUpdate().
		Model(&input).
		Where("oid = ? and repository_id = ?", input.Oid, input.RepositoryID).
		Returning("*").
		Exec(ctx, &input)
	if err == nil {
		return &input, nil
	}

	res, err := s.db.Core.NewInsert().Model(&input).Exec(ctx, &input)
	if err := assertAffectedOneRow(res, err); err != nil {
		return nil, fmt.Errorf("create lfs meta object in tx failed,error:%w", err)
	}

	return &input, nil
}

func (s *lfsMetaObjectStoreImpl) BulkUpdateOrCreate(ctx context.Context, input []LfsMetaObject) error {
	if len(input) == 0 {
		return nil
	}
	_, err := s.db.Core.NewInsert().
		Model(&input).
		On("CONFLICT (oid, repository_id) DO UPDATE").
		Set("size = EXCLUDED.size, updated_at = EXCLUDED.updated_at, existing = EXCLUDED.existing").
		Exec(ctx)
	return err
}
