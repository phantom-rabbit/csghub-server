// Code generated by mockery v2.48.0. DO NOT EDIT.

package database

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	database "opencsg.com/csghub-server/builder/store/database"
)

// MockDatasetStore is an autogenerated mock type for the DatasetStore type
type MockDatasetStore struct {
	mock.Mock
}

type MockDatasetStore_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDatasetStore) EXPECT() *MockDatasetStore_Expecter {
	return &MockDatasetStore_Expecter{mock: &_m.Mock}
}

// ByOrgPath provides a mock function with given fields: ctx, namespace, per, page, onlyPublic
func (_m *MockDatasetStore) ByOrgPath(ctx context.Context, namespace string, per int, page int, onlyPublic bool) ([]database.Dataset, int, error) {
	ret := _m.Called(ctx, namespace, per, page, onlyPublic)

	if len(ret) == 0 {
		panic("no return value specified for ByOrgPath")
	}

	var r0 []database.Dataset
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int, bool) ([]database.Dataset, int, error)); ok {
		return rf(ctx, namespace, per, page, onlyPublic)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int, bool) []database.Dataset); ok {
		r0 = rf(ctx, namespace, per, page, onlyPublic)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]database.Dataset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int, int, bool) int); ok {
		r1 = rf(ctx, namespace, per, page, onlyPublic)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, int, int, bool) error); ok {
		r2 = rf(ctx, namespace, per, page, onlyPublic)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockDatasetStore_ByOrgPath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ByOrgPath'
type MockDatasetStore_ByOrgPath_Call struct {
	*mock.Call
}

// ByOrgPath is a helper method to define mock.On call
//   - ctx context.Context
//   - namespace string
//   - per int
//   - page int
//   - onlyPublic bool
func (_e *MockDatasetStore_Expecter) ByOrgPath(ctx interface{}, namespace interface{}, per interface{}, page interface{}, onlyPublic interface{}) *MockDatasetStore_ByOrgPath_Call {
	return &MockDatasetStore_ByOrgPath_Call{Call: _e.mock.On("ByOrgPath", ctx, namespace, per, page, onlyPublic)}
}

func (_c *MockDatasetStore_ByOrgPath_Call) Run(run func(ctx context.Context, namespace string, per int, page int, onlyPublic bool)) *MockDatasetStore_ByOrgPath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int), args[3].(int), args[4].(bool))
	})
	return _c
}

func (_c *MockDatasetStore_ByOrgPath_Call) Return(datasets []database.Dataset, total int, err error) *MockDatasetStore_ByOrgPath_Call {
	_c.Call.Return(datasets, total, err)
	return _c
}

func (_c *MockDatasetStore_ByOrgPath_Call) RunAndReturn(run func(context.Context, string, int, int, bool) ([]database.Dataset, int, error)) *MockDatasetStore_ByOrgPath_Call {
	_c.Call.Return(run)
	return _c
}

// ByRepoID provides a mock function with given fields: ctx, repoID
func (_m *MockDatasetStore) ByRepoID(ctx context.Context, repoID int64) (*database.Dataset, error) {
	ret := _m.Called(ctx, repoID)

	if len(ret) == 0 {
		panic("no return value specified for ByRepoID")
	}

	var r0 *database.Dataset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*database.Dataset, error)); ok {
		return rf(ctx, repoID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *database.Dataset); ok {
		r0 = rf(ctx, repoID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.Dataset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, repoID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDatasetStore_ByRepoID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ByRepoID'
type MockDatasetStore_ByRepoID_Call struct {
	*mock.Call
}

// ByRepoID is a helper method to define mock.On call
//   - ctx context.Context
//   - repoID int64
func (_e *MockDatasetStore_Expecter) ByRepoID(ctx interface{}, repoID interface{}) *MockDatasetStore_ByRepoID_Call {
	return &MockDatasetStore_ByRepoID_Call{Call: _e.mock.On("ByRepoID", ctx, repoID)}
}

func (_c *MockDatasetStore_ByRepoID_Call) Run(run func(ctx context.Context, repoID int64)) *MockDatasetStore_ByRepoID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockDatasetStore_ByRepoID_Call) Return(_a0 *database.Dataset, _a1 error) *MockDatasetStore_ByRepoID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDatasetStore_ByRepoID_Call) RunAndReturn(run func(context.Context, int64) (*database.Dataset, error)) *MockDatasetStore_ByRepoID_Call {
	_c.Call.Return(run)
	return _c
}

// ByRepoIDs provides a mock function with given fields: ctx, repoIDs
func (_m *MockDatasetStore) ByRepoIDs(ctx context.Context, repoIDs []int64) ([]database.Dataset, error) {
	ret := _m.Called(ctx, repoIDs)

	if len(ret) == 0 {
		panic("no return value specified for ByRepoIDs")
	}

	var r0 []database.Dataset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []int64) ([]database.Dataset, error)); ok {
		return rf(ctx, repoIDs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []int64) []database.Dataset); ok {
		r0 = rf(ctx, repoIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]database.Dataset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []int64) error); ok {
		r1 = rf(ctx, repoIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDatasetStore_ByRepoIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ByRepoIDs'
type MockDatasetStore_ByRepoIDs_Call struct {
	*mock.Call
}

// ByRepoIDs is a helper method to define mock.On call
//   - ctx context.Context
//   - repoIDs []int64
func (_e *MockDatasetStore_Expecter) ByRepoIDs(ctx interface{}, repoIDs interface{}) *MockDatasetStore_ByRepoIDs_Call {
	return &MockDatasetStore_ByRepoIDs_Call{Call: _e.mock.On("ByRepoIDs", ctx, repoIDs)}
}

func (_c *MockDatasetStore_ByRepoIDs_Call) Run(run func(ctx context.Context, repoIDs []int64)) *MockDatasetStore_ByRepoIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]int64))
	})
	return _c
}

func (_c *MockDatasetStore_ByRepoIDs_Call) Return(datasets []database.Dataset, err error) *MockDatasetStore_ByRepoIDs_Call {
	_c.Call.Return(datasets, err)
	return _c
}

func (_c *MockDatasetStore_ByRepoIDs_Call) RunAndReturn(run func(context.Context, []int64) ([]database.Dataset, error)) *MockDatasetStore_ByRepoIDs_Call {
	_c.Call.Return(run)
	return _c
}

// ByUsername provides a mock function with given fields: ctx, username, per, page, onlyPublic
func (_m *MockDatasetStore) ByUsername(ctx context.Context, username string, per int, page int, onlyPublic bool) ([]database.Dataset, int, error) {
	ret := _m.Called(ctx, username, per, page, onlyPublic)

	if len(ret) == 0 {
		panic("no return value specified for ByUsername")
	}

	var r0 []database.Dataset
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int, bool) ([]database.Dataset, int, error)); ok {
		return rf(ctx, username, per, page, onlyPublic)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int, bool) []database.Dataset); ok {
		r0 = rf(ctx, username, per, page, onlyPublic)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]database.Dataset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int, int, bool) int); ok {
		r1 = rf(ctx, username, per, page, onlyPublic)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, int, int, bool) error); ok {
		r2 = rf(ctx, username, per, page, onlyPublic)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockDatasetStore_ByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ByUsername'
type MockDatasetStore_ByUsername_Call struct {
	*mock.Call
}

// ByUsername is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
//   - per int
//   - page int
//   - onlyPublic bool
func (_e *MockDatasetStore_Expecter) ByUsername(ctx interface{}, username interface{}, per interface{}, page interface{}, onlyPublic interface{}) *MockDatasetStore_ByUsername_Call {
	return &MockDatasetStore_ByUsername_Call{Call: _e.mock.On("ByUsername", ctx, username, per, page, onlyPublic)}
}

func (_c *MockDatasetStore_ByUsername_Call) Run(run func(ctx context.Context, username string, per int, page int, onlyPublic bool)) *MockDatasetStore_ByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int), args[3].(int), args[4].(bool))
	})
	return _c
}

func (_c *MockDatasetStore_ByUsername_Call) Return(datasets []database.Dataset, total int, err error) *MockDatasetStore_ByUsername_Call {
	_c.Call.Return(datasets, total, err)
	return _c
}

func (_c *MockDatasetStore_ByUsername_Call) RunAndReturn(run func(context.Context, string, int, int, bool) ([]database.Dataset, int, error)) *MockDatasetStore_ByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, input
func (_m *MockDatasetStore) Create(ctx context.Context, input database.Dataset) (*database.Dataset, error) {
	ret := _m.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *database.Dataset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, database.Dataset) (*database.Dataset, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, database.Dataset) *database.Dataset); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.Dataset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, database.Dataset) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDatasetStore_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockDatasetStore_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - input database.Dataset
func (_e *MockDatasetStore_Expecter) Create(ctx interface{}, input interface{}) *MockDatasetStore_Create_Call {
	return &MockDatasetStore_Create_Call{Call: _e.mock.On("Create", ctx, input)}
}

func (_c *MockDatasetStore_Create_Call) Run(run func(ctx context.Context, input database.Dataset)) *MockDatasetStore_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.Dataset))
	})
	return _c
}

func (_c *MockDatasetStore_Create_Call) Return(_a0 *database.Dataset, _a1 error) *MockDatasetStore_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDatasetStore_Create_Call) RunAndReturn(run func(context.Context, database.Dataset) (*database.Dataset, error)) *MockDatasetStore_Create_Call {
	_c.Call.Return(run)
	return _c
}

// CreateIfNotExist provides a mock function with given fields: ctx, input
func (_m *MockDatasetStore) CreateIfNotExist(ctx context.Context, input database.Dataset) (*database.Dataset, error) {
	ret := _m.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for CreateIfNotExist")
	}

	var r0 *database.Dataset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, database.Dataset) (*database.Dataset, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, database.Dataset) *database.Dataset); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.Dataset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, database.Dataset) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDatasetStore_CreateIfNotExist_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateIfNotExist'
type MockDatasetStore_CreateIfNotExist_Call struct {
	*mock.Call
}

// CreateIfNotExist is a helper method to define mock.On call
//   - ctx context.Context
//   - input database.Dataset
func (_e *MockDatasetStore_Expecter) CreateIfNotExist(ctx interface{}, input interface{}) *MockDatasetStore_CreateIfNotExist_Call {
	return &MockDatasetStore_CreateIfNotExist_Call{Call: _e.mock.On("CreateIfNotExist", ctx, input)}
}

func (_c *MockDatasetStore_CreateIfNotExist_Call) Run(run func(ctx context.Context, input database.Dataset)) *MockDatasetStore_CreateIfNotExist_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.Dataset))
	})
	return _c
}

func (_c *MockDatasetStore_CreateIfNotExist_Call) Return(_a0 *database.Dataset, _a1 error) *MockDatasetStore_CreateIfNotExist_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDatasetStore_CreateIfNotExist_Call) RunAndReturn(run func(context.Context, database.Dataset) (*database.Dataset, error)) *MockDatasetStore_CreateIfNotExist_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, input
func (_m *MockDatasetStore) Delete(ctx context.Context, input database.Dataset) error {
	ret := _m.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.Dataset) error); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDatasetStore_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockDatasetStore_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - input database.Dataset
func (_e *MockDatasetStore_Expecter) Delete(ctx interface{}, input interface{}) *MockDatasetStore_Delete_Call {
	return &MockDatasetStore_Delete_Call{Call: _e.mock.On("Delete", ctx, input)}
}

func (_c *MockDatasetStore_Delete_Call) Run(run func(ctx context.Context, input database.Dataset)) *MockDatasetStore_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.Dataset))
	})
	return _c
}

func (_c *MockDatasetStore_Delete_Call) Return(_a0 error) *MockDatasetStore_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDatasetStore_Delete_Call) RunAndReturn(run func(context.Context, database.Dataset) error) *MockDatasetStore_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// FindByPath provides a mock function with given fields: ctx, namespace, repoPath
func (_m *MockDatasetStore) FindByPath(ctx context.Context, namespace string, repoPath string) (*database.Dataset, error) {
	ret := _m.Called(ctx, namespace, repoPath)

	if len(ret) == 0 {
		panic("no return value specified for FindByPath")
	}

	var r0 *database.Dataset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*database.Dataset, error)); ok {
		return rf(ctx, namespace, repoPath)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *database.Dataset); ok {
		r0 = rf(ctx, namespace, repoPath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.Dataset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, namespace, repoPath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDatasetStore_FindByPath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByPath'
type MockDatasetStore_FindByPath_Call struct {
	*mock.Call
}

// FindByPath is a helper method to define mock.On call
//   - ctx context.Context
//   - namespace string
//   - repoPath string
func (_e *MockDatasetStore_Expecter) FindByPath(ctx interface{}, namespace interface{}, repoPath interface{}) *MockDatasetStore_FindByPath_Call {
	return &MockDatasetStore_FindByPath_Call{Call: _e.mock.On("FindByPath", ctx, namespace, repoPath)}
}

func (_c *MockDatasetStore_FindByPath_Call) Run(run func(ctx context.Context, namespace string, repoPath string)) *MockDatasetStore_FindByPath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockDatasetStore_FindByPath_Call) Return(dataset *database.Dataset, err error) *MockDatasetStore_FindByPath_Call {
	_c.Call.Return(dataset, err)
	return _c
}

func (_c *MockDatasetStore_FindByPath_Call) RunAndReturn(run func(context.Context, string, string) (*database.Dataset, error)) *MockDatasetStore_FindByPath_Call {
	_c.Call.Return(run)
	return _c
}

// ListByPath provides a mock function with given fields: ctx, paths
func (_m *MockDatasetStore) ListByPath(ctx context.Context, paths []string) ([]database.Dataset, error) {
	ret := _m.Called(ctx, paths)

	if len(ret) == 0 {
		panic("no return value specified for ListByPath")
	}

	var r0 []database.Dataset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) ([]database.Dataset, error)); ok {
		return rf(ctx, paths)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) []database.Dataset); ok {
		r0 = rf(ctx, paths)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]database.Dataset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, paths)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDatasetStore_ListByPath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListByPath'
type MockDatasetStore_ListByPath_Call struct {
	*mock.Call
}

// ListByPath is a helper method to define mock.On call
//   - ctx context.Context
//   - paths []string
func (_e *MockDatasetStore_Expecter) ListByPath(ctx interface{}, paths interface{}) *MockDatasetStore_ListByPath_Call {
	return &MockDatasetStore_ListByPath_Call{Call: _e.mock.On("ListByPath", ctx, paths)}
}

func (_c *MockDatasetStore_ListByPath_Call) Run(run func(ctx context.Context, paths []string)) *MockDatasetStore_ListByPath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string))
	})
	return _c
}

func (_c *MockDatasetStore_ListByPath_Call) Return(_a0 []database.Dataset, _a1 error) *MockDatasetStore_ListByPath_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDatasetStore_ListByPath_Call) RunAndReturn(run func(context.Context, []string) ([]database.Dataset, error)) *MockDatasetStore_ListByPath_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, input
func (_m *MockDatasetStore) Update(ctx context.Context, input database.Dataset) error {
	ret := _m.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.Dataset) error); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDatasetStore_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockDatasetStore_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - input database.Dataset
func (_e *MockDatasetStore_Expecter) Update(ctx interface{}, input interface{}) *MockDatasetStore_Update_Call {
	return &MockDatasetStore_Update_Call{Call: _e.mock.On("Update", ctx, input)}
}

func (_c *MockDatasetStore_Update_Call) Run(run func(ctx context.Context, input database.Dataset)) *MockDatasetStore_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.Dataset))
	})
	return _c
}

func (_c *MockDatasetStore_Update_Call) Return(err error) *MockDatasetStore_Update_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockDatasetStore_Update_Call) RunAndReturn(run func(context.Context, database.Dataset) error) *MockDatasetStore_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UserLikesDatasets provides a mock function with given fields: ctx, userID, per, page
func (_m *MockDatasetStore) UserLikesDatasets(ctx context.Context, userID int64, per int, page int) ([]database.Dataset, int, error) {
	ret := _m.Called(ctx, userID, per, page)

	if len(ret) == 0 {
		panic("no return value specified for UserLikesDatasets")
	}

	var r0 []database.Dataset
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int, int) ([]database.Dataset, int, error)); ok {
		return rf(ctx, userID, per, page)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int, int) []database.Dataset); ok {
		r0 = rf(ctx, userID, per, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]database.Dataset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int, int) int); ok {
		r1 = rf(ctx, userID, per, page)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, int64, int, int) error); ok {
		r2 = rf(ctx, userID, per, page)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockDatasetStore_UserLikesDatasets_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserLikesDatasets'
type MockDatasetStore_UserLikesDatasets_Call struct {
	*mock.Call
}

// UserLikesDatasets is a helper method to define mock.On call
//   - ctx context.Context
//   - userID int64
//   - per int
//   - page int
func (_e *MockDatasetStore_Expecter) UserLikesDatasets(ctx interface{}, userID interface{}, per interface{}, page interface{}) *MockDatasetStore_UserLikesDatasets_Call {
	return &MockDatasetStore_UserLikesDatasets_Call{Call: _e.mock.On("UserLikesDatasets", ctx, userID, per, page)}
}

func (_c *MockDatasetStore_UserLikesDatasets_Call) Run(run func(ctx context.Context, userID int64, per int, page int)) *MockDatasetStore_UserLikesDatasets_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(int), args[3].(int))
	})
	return _c
}

func (_c *MockDatasetStore_UserLikesDatasets_Call) Return(datasets []database.Dataset, total int, err error) *MockDatasetStore_UserLikesDatasets_Call {
	_c.Call.Return(datasets, total, err)
	return _c
}

func (_c *MockDatasetStore_UserLikesDatasets_Call) RunAndReturn(run func(context.Context, int64, int, int) ([]database.Dataset, int, error)) *MockDatasetStore_UserLikesDatasets_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDatasetStore creates a new instance of MockDatasetStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDatasetStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDatasetStore {
	mock := &MockDatasetStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
