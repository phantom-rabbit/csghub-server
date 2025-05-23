// Code generated by mockery v2.53.0. DO NOT EDIT.

package database

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	database "opencsg.com/csghub-server/builder/store/database"
)

// MockDataviewerStore is an autogenerated mock type for the DataviewerStore type
type MockDataviewerStore struct {
	mock.Mock
}

type MockDataviewerStore_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDataviewerStore) EXPECT() *MockDataviewerStore_Expecter {
	return &MockDataviewerStore_Expecter{mock: &_m.Mock}
}

// CreateJob provides a mock function with given fields: ctx, job
func (_m *MockDataviewerStore) CreateJob(ctx context.Context, job database.DataviewerJob) error {
	ret := _m.Called(ctx, job)

	if len(ret) == 0 {
		panic("no return value specified for CreateJob")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.DataviewerJob) error); ok {
		r0 = rf(ctx, job)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDataviewerStore_CreateJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateJob'
type MockDataviewerStore_CreateJob_Call struct {
	*mock.Call
}

// CreateJob is a helper method to define mock.On call
//   - ctx context.Context
//   - job database.DataviewerJob
func (_e *MockDataviewerStore_Expecter) CreateJob(ctx interface{}, job interface{}) *MockDataviewerStore_CreateJob_Call {
	return &MockDataviewerStore_CreateJob_Call{Call: _e.mock.On("CreateJob", ctx, job)}
}

func (_c *MockDataviewerStore_CreateJob_Call) Run(run func(ctx context.Context, job database.DataviewerJob)) *MockDataviewerStore_CreateJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.DataviewerJob))
	})
	return _c
}

func (_c *MockDataviewerStore_CreateJob_Call) Return(_a0 error) *MockDataviewerStore_CreateJob_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDataviewerStore_CreateJob_Call) RunAndReturn(run func(context.Context, database.DataviewerJob) error) *MockDataviewerStore_CreateJob_Call {
	_c.Call.Return(run)
	return _c
}

// CreateViewer provides a mock function with given fields: ctx, viewer
func (_m *MockDataviewerStore) CreateViewer(ctx context.Context, viewer database.Dataviewer) error {
	ret := _m.Called(ctx, viewer)

	if len(ret) == 0 {
		panic("no return value specified for CreateViewer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.Dataviewer) error); ok {
		r0 = rf(ctx, viewer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDataviewerStore_CreateViewer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateViewer'
type MockDataviewerStore_CreateViewer_Call struct {
	*mock.Call
}

// CreateViewer is a helper method to define mock.On call
//   - ctx context.Context
//   - viewer database.Dataviewer
func (_e *MockDataviewerStore_Expecter) CreateViewer(ctx interface{}, viewer interface{}) *MockDataviewerStore_CreateViewer_Call {
	return &MockDataviewerStore_CreateViewer_Call{Call: _e.mock.On("CreateViewer", ctx, viewer)}
}

func (_c *MockDataviewerStore_CreateViewer_Call) Run(run func(ctx context.Context, viewer database.Dataviewer)) *MockDataviewerStore_CreateViewer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.Dataviewer))
	})
	return _c
}

func (_c *MockDataviewerStore_CreateViewer_Call) Return(_a0 error) *MockDataviewerStore_CreateViewer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDataviewerStore_CreateViewer_Call) RunAndReturn(run func(context.Context, database.Dataviewer) error) *MockDataviewerStore_CreateViewer_Call {
	_c.Call.Return(run)
	return _c
}

// GetJob provides a mock function with given fields: ctx, workflowID
func (_m *MockDataviewerStore) GetJob(ctx context.Context, workflowID string) (*database.DataviewerJob, error) {
	ret := _m.Called(ctx, workflowID)

	if len(ret) == 0 {
		panic("no return value specified for GetJob")
	}

	var r0 *database.DataviewerJob
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*database.DataviewerJob, error)); ok {
		return rf(ctx, workflowID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *database.DataviewerJob); ok {
		r0 = rf(ctx, workflowID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.DataviewerJob)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, workflowID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDataviewerStore_GetJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetJob'
type MockDataviewerStore_GetJob_Call struct {
	*mock.Call
}

// GetJob is a helper method to define mock.On call
//   - ctx context.Context
//   - workflowID string
func (_e *MockDataviewerStore_Expecter) GetJob(ctx interface{}, workflowID interface{}) *MockDataviewerStore_GetJob_Call {
	return &MockDataviewerStore_GetJob_Call{Call: _e.mock.On("GetJob", ctx, workflowID)}
}

func (_c *MockDataviewerStore_GetJob_Call) Run(run func(ctx context.Context, workflowID string)) *MockDataviewerStore_GetJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockDataviewerStore_GetJob_Call) Return(_a0 *database.DataviewerJob, _a1 error) *MockDataviewerStore_GetJob_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDataviewerStore_GetJob_Call) RunAndReturn(run func(context.Context, string) (*database.DataviewerJob, error)) *MockDataviewerStore_GetJob_Call {
	_c.Call.Return(run)
	return _c
}

// GetRunningJobsByRepoID provides a mock function with given fields: ctx, repoID
func (_m *MockDataviewerStore) GetRunningJobsByRepoID(ctx context.Context, repoID int64) ([]database.DataviewerJob, error) {
	ret := _m.Called(ctx, repoID)

	if len(ret) == 0 {
		panic("no return value specified for GetRunningJobsByRepoID")
	}

	var r0 []database.DataviewerJob
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]database.DataviewerJob, error)); ok {
		return rf(ctx, repoID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []database.DataviewerJob); ok {
		r0 = rf(ctx, repoID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]database.DataviewerJob)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, repoID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDataviewerStore_GetRunningJobsByRepoID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRunningJobsByRepoID'
type MockDataviewerStore_GetRunningJobsByRepoID_Call struct {
	*mock.Call
}

// GetRunningJobsByRepoID is a helper method to define mock.On call
//   - ctx context.Context
//   - repoID int64
func (_e *MockDataviewerStore_Expecter) GetRunningJobsByRepoID(ctx interface{}, repoID interface{}) *MockDataviewerStore_GetRunningJobsByRepoID_Call {
	return &MockDataviewerStore_GetRunningJobsByRepoID_Call{Call: _e.mock.On("GetRunningJobsByRepoID", ctx, repoID)}
}

func (_c *MockDataviewerStore_GetRunningJobsByRepoID_Call) Run(run func(ctx context.Context, repoID int64)) *MockDataviewerStore_GetRunningJobsByRepoID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockDataviewerStore_GetRunningJobsByRepoID_Call) Return(_a0 []database.DataviewerJob, _a1 error) *MockDataviewerStore_GetRunningJobsByRepoID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDataviewerStore_GetRunningJobsByRepoID_Call) RunAndReturn(run func(context.Context, int64) ([]database.DataviewerJob, error)) *MockDataviewerStore_GetRunningJobsByRepoID_Call {
	_c.Call.Return(run)
	return _c
}

// GetViewerByRepoID provides a mock function with given fields: ctx, repoID
func (_m *MockDataviewerStore) GetViewerByRepoID(ctx context.Context, repoID int64) (*database.Dataviewer, error) {
	ret := _m.Called(ctx, repoID)

	if len(ret) == 0 {
		panic("no return value specified for GetViewerByRepoID")
	}

	var r0 *database.Dataviewer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*database.Dataviewer, error)); ok {
		return rf(ctx, repoID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *database.Dataviewer); ok {
		r0 = rf(ctx, repoID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.Dataviewer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, repoID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDataviewerStore_GetViewerByRepoID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetViewerByRepoID'
type MockDataviewerStore_GetViewerByRepoID_Call struct {
	*mock.Call
}

// GetViewerByRepoID is a helper method to define mock.On call
//   - ctx context.Context
//   - repoID int64
func (_e *MockDataviewerStore_Expecter) GetViewerByRepoID(ctx interface{}, repoID interface{}) *MockDataviewerStore_GetViewerByRepoID_Call {
	return &MockDataviewerStore_GetViewerByRepoID_Call{Call: _e.mock.On("GetViewerByRepoID", ctx, repoID)}
}

func (_c *MockDataviewerStore_GetViewerByRepoID_Call) Run(run func(ctx context.Context, repoID int64)) *MockDataviewerStore_GetViewerByRepoID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockDataviewerStore_GetViewerByRepoID_Call) Return(_a0 *database.Dataviewer, _a1 error) *MockDataviewerStore_GetViewerByRepoID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDataviewerStore_GetViewerByRepoID_Call) RunAndReturn(run func(context.Context, int64) (*database.Dataviewer, error)) *MockDataviewerStore_GetViewerByRepoID_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateJob provides a mock function with given fields: ctx, job
func (_m *MockDataviewerStore) UpdateJob(ctx context.Context, job database.DataviewerJob) (*database.DataviewerJob, error) {
	ret := _m.Called(ctx, job)

	if len(ret) == 0 {
		panic("no return value specified for UpdateJob")
	}

	var r0 *database.DataviewerJob
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, database.DataviewerJob) (*database.DataviewerJob, error)); ok {
		return rf(ctx, job)
	}
	if rf, ok := ret.Get(0).(func(context.Context, database.DataviewerJob) *database.DataviewerJob); ok {
		r0 = rf(ctx, job)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.DataviewerJob)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, database.DataviewerJob) error); ok {
		r1 = rf(ctx, job)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDataviewerStore_UpdateJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateJob'
type MockDataviewerStore_UpdateJob_Call struct {
	*mock.Call
}

// UpdateJob is a helper method to define mock.On call
//   - ctx context.Context
//   - job database.DataviewerJob
func (_e *MockDataviewerStore_Expecter) UpdateJob(ctx interface{}, job interface{}) *MockDataviewerStore_UpdateJob_Call {
	return &MockDataviewerStore_UpdateJob_Call{Call: _e.mock.On("UpdateJob", ctx, job)}
}

func (_c *MockDataviewerStore_UpdateJob_Call) Run(run func(ctx context.Context, job database.DataviewerJob)) *MockDataviewerStore_UpdateJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.DataviewerJob))
	})
	return _c
}

func (_c *MockDataviewerStore_UpdateJob_Call) Return(_a0 *database.DataviewerJob, _a1 error) *MockDataviewerStore_UpdateJob_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDataviewerStore_UpdateJob_Call) RunAndReturn(run func(context.Context, database.DataviewerJob) (*database.DataviewerJob, error)) *MockDataviewerStore_UpdateJob_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateViewer provides a mock function with given fields: ctx, viewer
func (_m *MockDataviewerStore) UpdateViewer(ctx context.Context, viewer database.Dataviewer) (*database.Dataviewer, error) {
	ret := _m.Called(ctx, viewer)

	if len(ret) == 0 {
		panic("no return value specified for UpdateViewer")
	}

	var r0 *database.Dataviewer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, database.Dataviewer) (*database.Dataviewer, error)); ok {
		return rf(ctx, viewer)
	}
	if rf, ok := ret.Get(0).(func(context.Context, database.Dataviewer) *database.Dataviewer); ok {
		r0 = rf(ctx, viewer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.Dataviewer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, database.Dataviewer) error); ok {
		r1 = rf(ctx, viewer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDataviewerStore_UpdateViewer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateViewer'
type MockDataviewerStore_UpdateViewer_Call struct {
	*mock.Call
}

// UpdateViewer is a helper method to define mock.On call
//   - ctx context.Context
//   - viewer database.Dataviewer
func (_e *MockDataviewerStore_Expecter) UpdateViewer(ctx interface{}, viewer interface{}) *MockDataviewerStore_UpdateViewer_Call {
	return &MockDataviewerStore_UpdateViewer_Call{Call: _e.mock.On("UpdateViewer", ctx, viewer)}
}

func (_c *MockDataviewerStore_UpdateViewer_Call) Run(run func(ctx context.Context, viewer database.Dataviewer)) *MockDataviewerStore_UpdateViewer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.Dataviewer))
	})
	return _c
}

func (_c *MockDataviewerStore_UpdateViewer_Call) Return(_a0 *database.Dataviewer, _a1 error) *MockDataviewerStore_UpdateViewer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDataviewerStore_UpdateViewer_Call) RunAndReturn(run func(context.Context, database.Dataviewer) (*database.Dataviewer, error)) *MockDataviewerStore_UpdateViewer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDataviewerStore creates a new instance of MockDataviewerStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDataviewerStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDataviewerStore {
	mock := &MockDataviewerStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
