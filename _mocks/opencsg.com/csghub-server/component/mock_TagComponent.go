// Code generated by mockery v2.53.0. DO NOT EDIT.

package component

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	database "opencsg.com/csghub-server/builder/store/database"

	types "opencsg.com/csghub-server/common/types"
)

// MockTagComponent is an autogenerated mock type for the TagComponent type
type MockTagComponent struct {
	mock.Mock
}

type MockTagComponent_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTagComponent) EXPECT() *MockTagComponent_Expecter {
	return &MockTagComponent_Expecter{mock: &_m.Mock}
}

// AllCategories provides a mock function with given fields: ctx
func (_m *MockTagComponent) AllCategories(ctx context.Context) ([]database.TagCategory, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for AllCategories")
	}

	var r0 []database.TagCategory
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]database.TagCategory, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []database.TagCategory); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]database.TagCategory)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTagComponent_AllCategories_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllCategories'
type MockTagComponent_AllCategories_Call struct {
	*mock.Call
}

// AllCategories is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockTagComponent_Expecter) AllCategories(ctx interface{}) *MockTagComponent_AllCategories_Call {
	return &MockTagComponent_AllCategories_Call{Call: _e.mock.On("AllCategories", ctx)}
}

func (_c *MockTagComponent_AllCategories_Call) Run(run func(ctx context.Context)) *MockTagComponent_AllCategories_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockTagComponent_AllCategories_Call) Return(_a0 []database.TagCategory, _a1 error) *MockTagComponent_AllCategories_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTagComponent_AllCategories_Call) RunAndReturn(run func(context.Context) ([]database.TagCategory, error)) *MockTagComponent_AllCategories_Call {
	_c.Call.Return(run)
	return _c
}

// AllTags provides a mock function with given fields: ctx, filter
func (_m *MockTagComponent) AllTags(ctx context.Context, filter *types.TagFilter) ([]*types.RepoTag, error) {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for AllTags")
	}

	var r0 []*types.RepoTag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.TagFilter) ([]*types.RepoTag, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.TagFilter) []*types.RepoTag); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.RepoTag)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.TagFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTagComponent_AllTags_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllTags'
type MockTagComponent_AllTags_Call struct {
	*mock.Call
}

// AllTags is a helper method to define mock.On call
//   - ctx context.Context
//   - filter *types.TagFilter
func (_e *MockTagComponent_Expecter) AllTags(ctx interface{}, filter interface{}) *MockTagComponent_AllTags_Call {
	return &MockTagComponent_AllTags_Call{Call: _e.mock.On("AllTags", ctx, filter)}
}

func (_c *MockTagComponent_AllTags_Call) Run(run func(ctx context.Context, filter *types.TagFilter)) *MockTagComponent_AllTags_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*types.TagFilter))
	})
	return _c
}

func (_c *MockTagComponent_AllTags_Call) Return(_a0 []*types.RepoTag, _a1 error) *MockTagComponent_AllTags_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTagComponent_AllTags_Call) RunAndReturn(run func(context.Context, *types.TagFilter) ([]*types.RepoTag, error)) *MockTagComponent_AllTags_Call {
	_c.Call.Return(run)
	return _c
}

// ClearMetaTags provides a mock function with given fields: ctx, repoType, namespace, name
func (_m *MockTagComponent) ClearMetaTags(ctx context.Context, repoType types.RepositoryType, namespace string, name string) error {
	ret := _m.Called(ctx, repoType, namespace, name)

	if len(ret) == 0 {
		panic("no return value specified for ClearMetaTags")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.RepositoryType, string, string) error); ok {
		r0 = rf(ctx, repoType, namespace, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTagComponent_ClearMetaTags_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClearMetaTags'
type MockTagComponent_ClearMetaTags_Call struct {
	*mock.Call
}

// ClearMetaTags is a helper method to define mock.On call
//   - ctx context.Context
//   - repoType types.RepositoryType
//   - namespace string
//   - name string
func (_e *MockTagComponent_Expecter) ClearMetaTags(ctx interface{}, repoType interface{}, namespace interface{}, name interface{}) *MockTagComponent_ClearMetaTags_Call {
	return &MockTagComponent_ClearMetaTags_Call{Call: _e.mock.On("ClearMetaTags", ctx, repoType, namespace, name)}
}

func (_c *MockTagComponent_ClearMetaTags_Call) Run(run func(ctx context.Context, repoType types.RepositoryType, namespace string, name string)) *MockTagComponent_ClearMetaTags_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.RepositoryType), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockTagComponent_ClearMetaTags_Call) Return(_a0 error) *MockTagComponent_ClearMetaTags_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTagComponent_ClearMetaTags_Call) RunAndReturn(run func(context.Context, types.RepositoryType, string, string) error) *MockTagComponent_ClearMetaTags_Call {
	_c.Call.Return(run)
	return _c
}

// CreateCategory provides a mock function with given fields: ctx, username, req
func (_m *MockTagComponent) CreateCategory(ctx context.Context, username string, req types.CreateCategory) (*database.TagCategory, error) {
	ret := _m.Called(ctx, username, req)

	if len(ret) == 0 {
		panic("no return value specified for CreateCategory")
	}

	var r0 *database.TagCategory
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.CreateCategory) (*database.TagCategory, error)); ok {
		return rf(ctx, username, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.CreateCategory) *database.TagCategory); ok {
		r0 = rf(ctx, username, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.TagCategory)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.CreateCategory) error); ok {
		r1 = rf(ctx, username, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTagComponent_CreateCategory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCategory'
type MockTagComponent_CreateCategory_Call struct {
	*mock.Call
}

// CreateCategory is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
//   - req types.CreateCategory
func (_e *MockTagComponent_Expecter) CreateCategory(ctx interface{}, username interface{}, req interface{}) *MockTagComponent_CreateCategory_Call {
	return &MockTagComponent_CreateCategory_Call{Call: _e.mock.On("CreateCategory", ctx, username, req)}
}

func (_c *MockTagComponent_CreateCategory_Call) Run(run func(ctx context.Context, username string, req types.CreateCategory)) *MockTagComponent_CreateCategory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(types.CreateCategory))
	})
	return _c
}

func (_c *MockTagComponent_CreateCategory_Call) Return(_a0 *database.TagCategory, _a1 error) *MockTagComponent_CreateCategory_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTagComponent_CreateCategory_Call) RunAndReturn(run func(context.Context, string, types.CreateCategory) (*database.TagCategory, error)) *MockTagComponent_CreateCategory_Call {
	_c.Call.Return(run)
	return _c
}

// CreateTag provides a mock function with given fields: ctx, username, req
func (_m *MockTagComponent) CreateTag(ctx context.Context, username string, req types.CreateTag) (*database.Tag, error) {
	ret := _m.Called(ctx, username, req)

	if len(ret) == 0 {
		panic("no return value specified for CreateTag")
	}

	var r0 *database.Tag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.CreateTag) (*database.Tag, error)); ok {
		return rf(ctx, username, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.CreateTag) *database.Tag); ok {
		r0 = rf(ctx, username, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.Tag)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.CreateTag) error); ok {
		r1 = rf(ctx, username, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTagComponent_CreateTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTag'
type MockTagComponent_CreateTag_Call struct {
	*mock.Call
}

// CreateTag is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
//   - req types.CreateTag
func (_e *MockTagComponent_Expecter) CreateTag(ctx interface{}, username interface{}, req interface{}) *MockTagComponent_CreateTag_Call {
	return &MockTagComponent_CreateTag_Call{Call: _e.mock.On("CreateTag", ctx, username, req)}
}

func (_c *MockTagComponent_CreateTag_Call) Run(run func(ctx context.Context, username string, req types.CreateTag)) *MockTagComponent_CreateTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(types.CreateTag))
	})
	return _c
}

func (_c *MockTagComponent_CreateTag_Call) Return(_a0 *database.Tag, _a1 error) *MockTagComponent_CreateTag_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTagComponent_CreateTag_Call) RunAndReturn(run func(context.Context, string, types.CreateTag) (*database.Tag, error)) *MockTagComponent_CreateTag_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCategory provides a mock function with given fields: ctx, username, id
func (_m *MockTagComponent) DeleteCategory(ctx context.Context, username string, id int64) error {
	ret := _m.Called(ctx, username, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCategory")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) error); ok {
		r0 = rf(ctx, username, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTagComponent_DeleteCategory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCategory'
type MockTagComponent_DeleteCategory_Call struct {
	*mock.Call
}

// DeleteCategory is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
//   - id int64
func (_e *MockTagComponent_Expecter) DeleteCategory(ctx interface{}, username interface{}, id interface{}) *MockTagComponent_DeleteCategory_Call {
	return &MockTagComponent_DeleteCategory_Call{Call: _e.mock.On("DeleteCategory", ctx, username, id)}
}

func (_c *MockTagComponent_DeleteCategory_Call) Run(run func(ctx context.Context, username string, id int64)) *MockTagComponent_DeleteCategory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int64))
	})
	return _c
}

func (_c *MockTagComponent_DeleteCategory_Call) Return(_a0 error) *MockTagComponent_DeleteCategory_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTagComponent_DeleteCategory_Call) RunAndReturn(run func(context.Context, string, int64) error) *MockTagComponent_DeleteCategory_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteTag provides a mock function with given fields: ctx, username, id
func (_m *MockTagComponent) DeleteTag(ctx context.Context, username string, id int64) error {
	ret := _m.Called(ctx, username, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTag")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) error); ok {
		r0 = rf(ctx, username, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTagComponent_DeleteTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteTag'
type MockTagComponent_DeleteTag_Call struct {
	*mock.Call
}

// DeleteTag is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
//   - id int64
func (_e *MockTagComponent_Expecter) DeleteTag(ctx interface{}, username interface{}, id interface{}) *MockTagComponent_DeleteTag_Call {
	return &MockTagComponent_DeleteTag_Call{Call: _e.mock.On("DeleteTag", ctx, username, id)}
}

func (_c *MockTagComponent_DeleteTag_Call) Run(run func(ctx context.Context, username string, id int64)) *MockTagComponent_DeleteTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int64))
	})
	return _c
}

func (_c *MockTagComponent_DeleteTag_Call) Return(_a0 error) *MockTagComponent_DeleteTag_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTagComponent_DeleteTag_Call) RunAndReturn(run func(context.Context, string, int64) error) *MockTagComponent_DeleteTag_Call {
	_c.Call.Return(run)
	return _c
}

// GetTagByID provides a mock function with given fields: ctx, username, id
func (_m *MockTagComponent) GetTagByID(ctx context.Context, username string, id int64) (*database.Tag, error) {
	ret := _m.Called(ctx, username, id)

	if len(ret) == 0 {
		panic("no return value specified for GetTagByID")
	}

	var r0 *database.Tag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) (*database.Tag, error)); ok {
		return rf(ctx, username, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) *database.Tag); ok {
		r0 = rf(ctx, username, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.Tag)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(ctx, username, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTagComponent_GetTagByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTagByID'
type MockTagComponent_GetTagByID_Call struct {
	*mock.Call
}

// GetTagByID is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
//   - id int64
func (_e *MockTagComponent_Expecter) GetTagByID(ctx interface{}, username interface{}, id interface{}) *MockTagComponent_GetTagByID_Call {
	return &MockTagComponent_GetTagByID_Call{Call: _e.mock.On("GetTagByID", ctx, username, id)}
}

func (_c *MockTagComponent_GetTagByID_Call) Run(run func(ctx context.Context, username string, id int64)) *MockTagComponent_GetTagByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int64))
	})
	return _c
}

func (_c *MockTagComponent_GetTagByID_Call) Return(_a0 *database.Tag, _a1 error) *MockTagComponent_GetTagByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTagComponent_GetTagByID_Call) RunAndReturn(run func(context.Context, string, int64) (*database.Tag, error)) *MockTagComponent_GetTagByID_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateCategory provides a mock function with given fields: ctx, username, req, id
func (_m *MockTagComponent) UpdateCategory(ctx context.Context, username string, req types.UpdateCategory, id int64) (*database.TagCategory, error) {
	ret := _m.Called(ctx, username, req, id)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCategory")
	}

	var r0 *database.TagCategory
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.UpdateCategory, int64) (*database.TagCategory, error)); ok {
		return rf(ctx, username, req, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, types.UpdateCategory, int64) *database.TagCategory); ok {
		r0 = rf(ctx, username, req, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.TagCategory)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, types.UpdateCategory, int64) error); ok {
		r1 = rf(ctx, username, req, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTagComponent_UpdateCategory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateCategory'
type MockTagComponent_UpdateCategory_Call struct {
	*mock.Call
}

// UpdateCategory is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
//   - req types.UpdateCategory
//   - id int64
func (_e *MockTagComponent_Expecter) UpdateCategory(ctx interface{}, username interface{}, req interface{}, id interface{}) *MockTagComponent_UpdateCategory_Call {
	return &MockTagComponent_UpdateCategory_Call{Call: _e.mock.On("UpdateCategory", ctx, username, req, id)}
}

func (_c *MockTagComponent_UpdateCategory_Call) Run(run func(ctx context.Context, username string, req types.UpdateCategory, id int64)) *MockTagComponent_UpdateCategory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(types.UpdateCategory), args[3].(int64))
	})
	return _c
}

func (_c *MockTagComponent_UpdateCategory_Call) Return(_a0 *database.TagCategory, _a1 error) *MockTagComponent_UpdateCategory_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTagComponent_UpdateCategory_Call) RunAndReturn(run func(context.Context, string, types.UpdateCategory, int64) (*database.TagCategory, error)) *MockTagComponent_UpdateCategory_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateLibraryTags provides a mock function with given fields: ctx, tagScope, namespace, name, oldFilePath, newFilePath
func (_m *MockTagComponent) UpdateLibraryTags(ctx context.Context, tagScope types.TagScope, namespace string, name string, oldFilePath string, newFilePath string) error {
	ret := _m.Called(ctx, tagScope, namespace, name, oldFilePath, newFilePath)

	if len(ret) == 0 {
		panic("no return value specified for UpdateLibraryTags")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.TagScope, string, string, string, string) error); ok {
		r0 = rf(ctx, tagScope, namespace, name, oldFilePath, newFilePath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTagComponent_UpdateLibraryTags_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateLibraryTags'
type MockTagComponent_UpdateLibraryTags_Call struct {
	*mock.Call
}

// UpdateLibraryTags is a helper method to define mock.On call
//   - ctx context.Context
//   - tagScope types.TagScope
//   - namespace string
//   - name string
//   - oldFilePath string
//   - newFilePath string
func (_e *MockTagComponent_Expecter) UpdateLibraryTags(ctx interface{}, tagScope interface{}, namespace interface{}, name interface{}, oldFilePath interface{}, newFilePath interface{}) *MockTagComponent_UpdateLibraryTags_Call {
	return &MockTagComponent_UpdateLibraryTags_Call{Call: _e.mock.On("UpdateLibraryTags", ctx, tagScope, namespace, name, oldFilePath, newFilePath)}
}

func (_c *MockTagComponent_UpdateLibraryTags_Call) Run(run func(ctx context.Context, tagScope types.TagScope, namespace string, name string, oldFilePath string, newFilePath string)) *MockTagComponent_UpdateLibraryTags_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.TagScope), args[2].(string), args[3].(string), args[4].(string), args[5].(string))
	})
	return _c
}

func (_c *MockTagComponent_UpdateLibraryTags_Call) Return(_a0 error) *MockTagComponent_UpdateLibraryTags_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTagComponent_UpdateLibraryTags_Call) RunAndReturn(run func(context.Context, types.TagScope, string, string, string, string) error) *MockTagComponent_UpdateLibraryTags_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateMetaTags provides a mock function with given fields: ctx, tagScope, namespace, name, content
func (_m *MockTagComponent) UpdateMetaTags(ctx context.Context, tagScope types.TagScope, namespace string, name string, content string) ([]*database.RepositoryTag, error) {
	ret := _m.Called(ctx, tagScope, namespace, name, content)

	if len(ret) == 0 {
		panic("no return value specified for UpdateMetaTags")
	}

	var r0 []*database.RepositoryTag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.TagScope, string, string, string) ([]*database.RepositoryTag, error)); ok {
		return rf(ctx, tagScope, namespace, name, content)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.TagScope, string, string, string) []*database.RepositoryTag); ok {
		r0 = rf(ctx, tagScope, namespace, name, content)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*database.RepositoryTag)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.TagScope, string, string, string) error); ok {
		r1 = rf(ctx, tagScope, namespace, name, content)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTagComponent_UpdateMetaTags_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateMetaTags'
type MockTagComponent_UpdateMetaTags_Call struct {
	*mock.Call
}

// UpdateMetaTags is a helper method to define mock.On call
//   - ctx context.Context
//   - tagScope types.TagScope
//   - namespace string
//   - name string
//   - content string
func (_e *MockTagComponent_Expecter) UpdateMetaTags(ctx interface{}, tagScope interface{}, namespace interface{}, name interface{}, content interface{}) *MockTagComponent_UpdateMetaTags_Call {
	return &MockTagComponent_UpdateMetaTags_Call{Call: _e.mock.On("UpdateMetaTags", ctx, tagScope, namespace, name, content)}
}

func (_c *MockTagComponent_UpdateMetaTags_Call) Run(run func(ctx context.Context, tagScope types.TagScope, namespace string, name string, content string)) *MockTagComponent_UpdateMetaTags_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.TagScope), args[2].(string), args[3].(string), args[4].(string))
	})
	return _c
}

func (_c *MockTagComponent_UpdateMetaTags_Call) Return(_a0 []*database.RepositoryTag, _a1 error) *MockTagComponent_UpdateMetaTags_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTagComponent_UpdateMetaTags_Call) RunAndReturn(run func(context.Context, types.TagScope, string, string, string) ([]*database.RepositoryTag, error)) *MockTagComponent_UpdateMetaTags_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateRepoTagsByCategory provides a mock function with given fields: ctx, tagScope, repoID, category, tagNames
func (_m *MockTagComponent) UpdateRepoTagsByCategory(ctx context.Context, tagScope types.TagScope, repoID int64, category string, tagNames []string) error {
	ret := _m.Called(ctx, tagScope, repoID, category, tagNames)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRepoTagsByCategory")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.TagScope, int64, string, []string) error); ok {
		r0 = rf(ctx, tagScope, repoID, category, tagNames)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTagComponent_UpdateRepoTagsByCategory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateRepoTagsByCategory'
type MockTagComponent_UpdateRepoTagsByCategory_Call struct {
	*mock.Call
}

// UpdateRepoTagsByCategory is a helper method to define mock.On call
//   - ctx context.Context
//   - tagScope types.TagScope
//   - repoID int64
//   - category string
//   - tagNames []string
func (_e *MockTagComponent_Expecter) UpdateRepoTagsByCategory(ctx interface{}, tagScope interface{}, repoID interface{}, category interface{}, tagNames interface{}) *MockTagComponent_UpdateRepoTagsByCategory_Call {
	return &MockTagComponent_UpdateRepoTagsByCategory_Call{Call: _e.mock.On("UpdateRepoTagsByCategory", ctx, tagScope, repoID, category, tagNames)}
}

func (_c *MockTagComponent_UpdateRepoTagsByCategory_Call) Run(run func(ctx context.Context, tagScope types.TagScope, repoID int64, category string, tagNames []string)) *MockTagComponent_UpdateRepoTagsByCategory_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.TagScope), args[2].(int64), args[3].(string), args[4].([]string))
	})
	return _c
}

func (_c *MockTagComponent_UpdateRepoTagsByCategory_Call) Return(_a0 error) *MockTagComponent_UpdateRepoTagsByCategory_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTagComponent_UpdateRepoTagsByCategory_Call) RunAndReturn(run func(context.Context, types.TagScope, int64, string, []string) error) *MockTagComponent_UpdateRepoTagsByCategory_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateTag provides a mock function with given fields: ctx, username, id, req
func (_m *MockTagComponent) UpdateTag(ctx context.Context, username string, id int64, req types.UpdateTag) (*database.Tag, error) {
	ret := _m.Called(ctx, username, id, req)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTag")
	}

	var r0 *database.Tag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, types.UpdateTag) (*database.Tag, error)); ok {
		return rf(ctx, username, id, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, types.UpdateTag) *database.Tag); ok {
		r0 = rf(ctx, username, id, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.Tag)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64, types.UpdateTag) error); ok {
		r1 = rf(ctx, username, id, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTagComponent_UpdateTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateTag'
type MockTagComponent_UpdateTag_Call struct {
	*mock.Call
}

// UpdateTag is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
//   - id int64
//   - req types.UpdateTag
func (_e *MockTagComponent_Expecter) UpdateTag(ctx interface{}, username interface{}, id interface{}, req interface{}) *MockTagComponent_UpdateTag_Call {
	return &MockTagComponent_UpdateTag_Call{Call: _e.mock.On("UpdateTag", ctx, username, id, req)}
}

func (_c *MockTagComponent_UpdateTag_Call) Run(run func(ctx context.Context, username string, id int64, req types.UpdateTag)) *MockTagComponent_UpdateTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int64), args[3].(types.UpdateTag))
	})
	return _c
}

func (_c *MockTagComponent_UpdateTag_Call) Return(_a0 *database.Tag, _a1 error) *MockTagComponent_UpdateTag_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTagComponent_UpdateTag_Call) RunAndReturn(run func(context.Context, string, int64, types.UpdateTag) (*database.Tag, error)) *MockTagComponent_UpdateTag_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTagComponent creates a new instance of MockTagComponent. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTagComponent(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTagComponent {
	mock := &MockTagComponent{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
