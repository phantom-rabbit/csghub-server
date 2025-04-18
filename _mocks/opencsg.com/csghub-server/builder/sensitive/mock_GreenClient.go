// Code generated by mockery v2.53.0. DO NOT EDIT.

package sensitive

import (
	green "github.com/aliyun/alibaba-cloud-sdk-go/services/green"
	mock "github.com/stretchr/testify/mock"

	sensitive "opencsg.com/csghub-server/builder/sensitive"
)

// MockGreenClient is an autogenerated mock type for the GreenClient type
type MockGreenClient struct {
	mock.Mock
}

type MockGreenClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockGreenClient) EXPECT() *MockGreenClient_Expecter {
	return &MockGreenClient_Expecter{mock: &_m.Mock}
}

// TextScan provides a mock function with given fields: request
func (_m *MockGreenClient) TextScan(request *green.TextScanRequest) (*sensitive.TextScanResponse, error) {
	ret := _m.Called(request)

	if len(ret) == 0 {
		panic("no return value specified for TextScan")
	}

	var r0 *sensitive.TextScanResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*green.TextScanRequest) (*sensitive.TextScanResponse, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(*green.TextScanRequest) *sensitive.TextScanResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sensitive.TextScanResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*green.TextScanRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGreenClient_TextScan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TextScan'
type MockGreenClient_TextScan_Call struct {
	*mock.Call
}

// TextScan is a helper method to define mock.On call
//   - request *green.TextScanRequest
func (_e *MockGreenClient_Expecter) TextScan(request interface{}) *MockGreenClient_TextScan_Call {
	return &MockGreenClient_TextScan_Call{Call: _e.mock.On("TextScan", request)}
}

func (_c *MockGreenClient_TextScan_Call) Run(run func(request *green.TextScanRequest)) *MockGreenClient_TextScan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*green.TextScanRequest))
	})
	return _c
}

func (_c *MockGreenClient_TextScan_Call) Return(response *sensitive.TextScanResponse, err error) *MockGreenClient_TextScan_Call {
	_c.Call.Return(response, err)
	return _c
}

func (_c *MockGreenClient_TextScan_Call) RunAndReturn(run func(*green.TextScanRequest) (*sensitive.TextScanResponse, error)) *MockGreenClient_TextScan_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockGreenClient creates a new instance of MockGreenClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGreenClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGreenClient {
	mock := &MockGreenClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
