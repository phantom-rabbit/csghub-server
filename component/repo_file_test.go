package component

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"opencsg.com/csghub-server/builder/git/gitserver"
	"opencsg.com/csghub-server/builder/store/database"
	"opencsg.com/csghub-server/common/types"
)

func TestRepoFileComponent_GenRepoFileRecords(t *testing.T) {
	ctx := context.TODO()
	rc := initializeTestRepoFileComponent(ctx, t)

	rc.mocks.stores.RepoMock().EXPECT().FindByPath(ctx, types.ModelRepo, "ns", "n").Return(
		&database.Repository{ID: 1, Path: "foo/bar"}, nil,
	)
	rc.mocks.gitServer.EXPECT().GetRepoFileTree(mock.Anything, gitserver.GetRepoInfoByPathReq{
		Namespace: "foo",
		Name:      "bar",
	}).Return(
		[]*types.File{
			{Path: "a/b", Type: "dir"},
			{Path: "foo.go", Type: "go"},
		}, nil,
	)
	rc.mocks.gitServer.EXPECT().GetRepoFileTree(mock.Anything, gitserver.GetRepoInfoByPathReq{
		Path:      "a/b",
		Namespace: "foo",
		Name:      "bar",
	}).Return(
		[]*types.File{}, nil,
	)
	rc.mocks.stores.RepoFileMock().EXPECT().Exists(ctx, database.RepositoryFile{
		RepositoryID: 1,
		Path:         "foo.go",
		FileType:     "go",
	}).Return(false, nil)
	rc.mocks.stores.RepoFileMock().EXPECT().Create(ctx, &database.RepositoryFile{
		RepositoryID: 1,
		Path:         "foo.go",
		FileType:     "go",
	}).Return(nil)

	err := rc.GenRepoFileRecords(ctx, types.ModelRepo, "ns", "n")
	require.Nil(t, err)

}

func TestRepoFileComponent_GenRepoFileRecordsBatch(t *testing.T) {
	ctx := context.TODO()
	rc := initializeTestRepoFileComponent(ctx, t)

	rc.mocks.stores.RepoMock().EXPECT().BatchGet(ctx, types.ModelRepo, int64(1), 10).Return(
		[]database.Repository{{ID: 1, Path: "foo/bar"}}, nil,
	)
	rc.mocks.gitServer.EXPECT().GetRepoFileTree(mock.Anything, gitserver.GetRepoInfoByPathReq{
		Namespace: "foo",
		Name:      "bar",
	}).Return(
		[]*types.File{
			{Path: "a/b", Type: "dir"},
			{Path: "foo.go", Type: "go"},
		}, nil,
	)
	rc.mocks.gitServer.EXPECT().GetRepoFileTree(mock.Anything, gitserver.GetRepoInfoByPathReq{
		Path:      "a/b",
		Namespace: "foo",
		Name:      "bar",
	}).Return(
		[]*types.File{}, nil,
	)
	rc.mocks.stores.RepoFileMock().EXPECT().Exists(ctx, database.RepositoryFile{
		RepositoryID: 1,
		Path:         "foo.go",
		FileType:     "go",
	}).Return(false, nil)
	rc.mocks.stores.RepoFileMock().EXPECT().Create(ctx, &database.RepositoryFile{
		RepositoryID: 1,
		Path:         "foo.go",
		FileType:     "go",
	}).Return(nil)

	err := rc.GenRepoFileRecordsBatch(ctx, types.ModelRepo, 1, 10)
	require.Nil(t, err)
}
