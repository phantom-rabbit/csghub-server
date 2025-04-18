// Code generated by mockery v2.53.0. DO NOT EDIT.

package component

import (
	common "opencsg.com/csghub-server/dataviewer/common"

	context "context"

	mock "github.com/stretchr/testify/mock"

	types "opencsg.com/csghub-server/common/types"
)

// MockDatasetViewerComponent is an autogenerated mock type for the DatasetViewerComponent type
type MockDatasetViewerComponent struct {
	mock.Mock
}

type MockDatasetViewerComponent_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDatasetViewerComponent) EXPECT() *MockDatasetViewerComponent_Expecter {
	return &MockDatasetViewerComponent_Expecter{mock: &_m.Mock}
}

// GetCatalog provides a mock function with given fields: ctx, req
func (_m *MockDatasetViewerComponent) GetCatalog(ctx context.Context, req *common.ViewParquetFileReq) (*common.CataLogRespone, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for GetCatalog")
	}

	var r0 *common.CataLogRespone
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *common.ViewParquetFileReq) (*common.CataLogRespone, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *common.ViewParquetFileReq) *common.CataLogRespone); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.CataLogRespone)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *common.ViewParquetFileReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDatasetViewerComponent_GetCatalog_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCatalog'
type MockDatasetViewerComponent_GetCatalog_Call struct {
	*mock.Call
}

// GetCatalog is a helper method to define mock.On call
//   - ctx context.Context
//   - req *common.ViewParquetFileReq
func (_e *MockDatasetViewerComponent_Expecter) GetCatalog(ctx interface{}, req interface{}) *MockDatasetViewerComponent_GetCatalog_Call {
	return &MockDatasetViewerComponent_GetCatalog_Call{Call: _e.mock.On("GetCatalog", ctx, req)}
}

func (_c *MockDatasetViewerComponent_GetCatalog_Call) Run(run func(ctx context.Context, req *common.ViewParquetFileReq)) *MockDatasetViewerComponent_GetCatalog_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*common.ViewParquetFileReq))
	})
	return _c
}

func (_c *MockDatasetViewerComponent_GetCatalog_Call) Return(_a0 *common.CataLogRespone, _a1 error) *MockDatasetViewerComponent_GetCatalog_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDatasetViewerComponent_GetCatalog_Call) RunAndReturn(run func(context.Context, *common.ViewParquetFileReq) (*common.CataLogRespone, error)) *MockDatasetViewerComponent_GetCatalog_Call {
	_c.Call.Return(run)
	return _c
}

// LimitOffsetRows provides a mock function with given fields: ctx, req, viewerReq
func (_m *MockDatasetViewerComponent) LimitOffsetRows(ctx context.Context, req *common.ViewParquetFileReq, viewerReq types.DataViewerReq) (*common.ViewParquetFileResp, error) {
	ret := _m.Called(ctx, req, viewerReq)

	if len(ret) == 0 {
		panic("no return value specified for LimitOffsetRows")
	}

	var r0 *common.ViewParquetFileResp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *common.ViewParquetFileReq, types.DataViewerReq) (*common.ViewParquetFileResp, error)); ok {
		return rf(ctx, req, viewerReq)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *common.ViewParquetFileReq, types.DataViewerReq) *common.ViewParquetFileResp); ok {
		r0 = rf(ctx, req, viewerReq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.ViewParquetFileResp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *common.ViewParquetFileReq, types.DataViewerReq) error); ok {
		r1 = rf(ctx, req, viewerReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDatasetViewerComponent_LimitOffsetRows_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LimitOffsetRows'
type MockDatasetViewerComponent_LimitOffsetRows_Call struct {
	*mock.Call
}

// LimitOffsetRows is a helper method to define mock.On call
//   - ctx context.Context
//   - req *common.ViewParquetFileReq
//   - viewerReq types.DataViewerReq
func (_e *MockDatasetViewerComponent_Expecter) LimitOffsetRows(ctx interface{}, req interface{}, viewerReq interface{}) *MockDatasetViewerComponent_LimitOffsetRows_Call {
	return &MockDatasetViewerComponent_LimitOffsetRows_Call{Call: _e.mock.On("LimitOffsetRows", ctx, req, viewerReq)}
}

func (_c *MockDatasetViewerComponent_LimitOffsetRows_Call) Run(run func(ctx context.Context, req *common.ViewParquetFileReq, viewerReq types.DataViewerReq)) *MockDatasetViewerComponent_LimitOffsetRows_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*common.ViewParquetFileReq), args[2].(types.DataViewerReq))
	})
	return _c
}

func (_c *MockDatasetViewerComponent_LimitOffsetRows_Call) Return(_a0 *common.ViewParquetFileResp, _a1 error) *MockDatasetViewerComponent_LimitOffsetRows_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDatasetViewerComponent_LimitOffsetRows_Call) RunAndReturn(run func(context.Context, *common.ViewParquetFileReq, types.DataViewerReq) (*common.ViewParquetFileResp, error)) *MockDatasetViewerComponent_LimitOffsetRows_Call {
	_c.Call.Return(run)
	return _c
}

// Rows provides a mock function with given fields: ctx, req, viewerReq
func (_m *MockDatasetViewerComponent) Rows(ctx context.Context, req *common.ViewParquetFileReq, viewerReq types.DataViewerReq) (*common.ViewParquetFileResp, error) {
	ret := _m.Called(ctx, req, viewerReq)

	if len(ret) == 0 {
		panic("no return value specified for Rows")
	}

	var r0 *common.ViewParquetFileResp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *common.ViewParquetFileReq, types.DataViewerReq) (*common.ViewParquetFileResp, error)); ok {
		return rf(ctx, req, viewerReq)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *common.ViewParquetFileReq, types.DataViewerReq) *common.ViewParquetFileResp); ok {
		r0 = rf(ctx, req, viewerReq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.ViewParquetFileResp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *common.ViewParquetFileReq, types.DataViewerReq) error); ok {
		r1 = rf(ctx, req, viewerReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDatasetViewerComponent_Rows_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rows'
type MockDatasetViewerComponent_Rows_Call struct {
	*mock.Call
}

// Rows is a helper method to define mock.On call
//   - ctx context.Context
//   - req *common.ViewParquetFileReq
//   - viewerReq types.DataViewerReq
func (_e *MockDatasetViewerComponent_Expecter) Rows(ctx interface{}, req interface{}, viewerReq interface{}) *MockDatasetViewerComponent_Rows_Call {
	return &MockDatasetViewerComponent_Rows_Call{Call: _e.mock.On("Rows", ctx, req, viewerReq)}
}

func (_c *MockDatasetViewerComponent_Rows_Call) Run(run func(ctx context.Context, req *common.ViewParquetFileReq, viewerReq types.DataViewerReq)) *MockDatasetViewerComponent_Rows_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*common.ViewParquetFileReq), args[2].(types.DataViewerReq))
	})
	return _c
}

func (_c *MockDatasetViewerComponent_Rows_Call) Return(_a0 *common.ViewParquetFileResp, _a1 error) *MockDatasetViewerComponent_Rows_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDatasetViewerComponent_Rows_Call) RunAndReturn(run func(context.Context, *common.ViewParquetFileReq, types.DataViewerReq) (*common.ViewParquetFileResp, error)) *MockDatasetViewerComponent_Rows_Call {
	_c.Call.Return(run)
	return _c
}

// ViewParquetFile provides a mock function with given fields: ctx, req
func (_m *MockDatasetViewerComponent) ViewParquetFile(ctx context.Context, req *common.ViewParquetFileReq) (*common.ViewParquetFileResp, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for ViewParquetFile")
	}

	var r0 *common.ViewParquetFileResp
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *common.ViewParquetFileReq) (*common.ViewParquetFileResp, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *common.ViewParquetFileReq) *common.ViewParquetFileResp); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.ViewParquetFileResp)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *common.ViewParquetFileReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDatasetViewerComponent_ViewParquetFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ViewParquetFile'
type MockDatasetViewerComponent_ViewParquetFile_Call struct {
	*mock.Call
}

// ViewParquetFile is a helper method to define mock.On call
//   - ctx context.Context
//   - req *common.ViewParquetFileReq
func (_e *MockDatasetViewerComponent_Expecter) ViewParquetFile(ctx interface{}, req interface{}) *MockDatasetViewerComponent_ViewParquetFile_Call {
	return &MockDatasetViewerComponent_ViewParquetFile_Call{Call: _e.mock.On("ViewParquetFile", ctx, req)}
}

func (_c *MockDatasetViewerComponent_ViewParquetFile_Call) Run(run func(ctx context.Context, req *common.ViewParquetFileReq)) *MockDatasetViewerComponent_ViewParquetFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*common.ViewParquetFileReq))
	})
	return _c
}

func (_c *MockDatasetViewerComponent_ViewParquetFile_Call) Return(_a0 *common.ViewParquetFileResp, _a1 error) *MockDatasetViewerComponent_ViewParquetFile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDatasetViewerComponent_ViewParquetFile_Call) RunAndReturn(run func(context.Context, *common.ViewParquetFileReq) (*common.ViewParquetFileResp, error)) *MockDatasetViewerComponent_ViewParquetFile_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDatasetViewerComponent creates a new instance of MockDatasetViewerComponent. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDatasetViewerComponent(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDatasetViewerComponent {
	mock := &MockDatasetViewerComponent{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
