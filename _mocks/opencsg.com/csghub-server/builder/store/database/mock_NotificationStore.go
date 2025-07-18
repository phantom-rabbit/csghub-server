// Code generated by mockery v2.53.0. DO NOT EDIT.

package database

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	database "opencsg.com/csghub-server/builder/store/database"
)

// MockNotificationStore is an autogenerated mock type for the NotificationStore type
type MockNotificationStore struct {
	mock.Mock
}

type MockNotificationStore_Expecter struct {
	mock *mock.Mock
}

func (_m *MockNotificationStore) EXPECT() *MockNotificationStore_Expecter {
	return &MockNotificationStore_Expecter{mock: &_m.Mock}
}

// CreateNotificationMessage provides a mock function with given fields: ctx, msg
func (_m *MockNotificationStore) CreateNotificationMessage(ctx context.Context, msg *database.NotificationMessage) error {
	ret := _m.Called(ctx, msg)

	if len(ret) == 0 {
		panic("no return value specified for CreateNotificationMessage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *database.NotificationMessage) error); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNotificationStore_CreateNotificationMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateNotificationMessage'
type MockNotificationStore_CreateNotificationMessage_Call struct {
	*mock.Call
}

// CreateNotificationMessage is a helper method to define mock.On call
//   - ctx context.Context
//   - msg *database.NotificationMessage
func (_e *MockNotificationStore_Expecter) CreateNotificationMessage(ctx interface{}, msg interface{}) *MockNotificationStore_CreateNotificationMessage_Call {
	return &MockNotificationStore_CreateNotificationMessage_Call{Call: _e.mock.On("CreateNotificationMessage", ctx, msg)}
}

func (_c *MockNotificationStore_CreateNotificationMessage_Call) Run(run func(ctx context.Context, msg *database.NotificationMessage)) *MockNotificationStore_CreateNotificationMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*database.NotificationMessage))
	})
	return _c
}

func (_c *MockNotificationStore_CreateNotificationMessage_Call) Return(_a0 error) *MockNotificationStore_CreateNotificationMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNotificationStore_CreateNotificationMessage_Call) RunAndReturn(run func(context.Context, *database.NotificationMessage) error) *MockNotificationStore_CreateNotificationMessage_Call {
	_c.Call.Return(run)
	return _c
}

// CreateNotificationMessageForUsers provides a mock function with given fields: ctx, msg, userUUIDs
func (_m *MockNotificationStore) CreateNotificationMessageForUsers(ctx context.Context, msg *database.NotificationMessage, userUUIDs []string) error {
	ret := _m.Called(ctx, msg, userUUIDs)

	if len(ret) == 0 {
		panic("no return value specified for CreateNotificationMessageForUsers")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *database.NotificationMessage, []string) error); ok {
		r0 = rf(ctx, msg, userUUIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNotificationStore_CreateNotificationMessageForUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateNotificationMessageForUsers'
type MockNotificationStore_CreateNotificationMessageForUsers_Call struct {
	*mock.Call
}

// CreateNotificationMessageForUsers is a helper method to define mock.On call
//   - ctx context.Context
//   - msg *database.NotificationMessage
//   - userUUIDs []string
func (_e *MockNotificationStore_Expecter) CreateNotificationMessageForUsers(ctx interface{}, msg interface{}, userUUIDs interface{}) *MockNotificationStore_CreateNotificationMessageForUsers_Call {
	return &MockNotificationStore_CreateNotificationMessageForUsers_Call{Call: _e.mock.On("CreateNotificationMessageForUsers", ctx, msg, userUUIDs)}
}

func (_c *MockNotificationStore_CreateNotificationMessageForUsers_Call) Run(run func(ctx context.Context, msg *database.NotificationMessage, userUUIDs []string)) *MockNotificationStore_CreateNotificationMessageForUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*database.NotificationMessage), args[2].([]string))
	})
	return _c
}

func (_c *MockNotificationStore_CreateNotificationMessageForUsers_Call) Return(_a0 error) *MockNotificationStore_CreateNotificationMessageForUsers_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNotificationStore_CreateNotificationMessageForUsers_Call) RunAndReturn(run func(context.Context, *database.NotificationMessage, []string) error) *MockNotificationStore_CreateNotificationMessageForUsers_Call {
	_c.Call.Return(run)
	return _c
}

// CreateSetting provides a mock function with given fields: ctx, setting
func (_m *MockNotificationStore) CreateSetting(ctx context.Context, setting *database.NotificationSetting) error {
	ret := _m.Called(ctx, setting)

	if len(ret) == 0 {
		panic("no return value specified for CreateSetting")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *database.NotificationSetting) error); ok {
		r0 = rf(ctx, setting)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNotificationStore_CreateSetting_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSetting'
type MockNotificationStore_CreateSetting_Call struct {
	*mock.Call
}

// CreateSetting is a helper method to define mock.On call
//   - ctx context.Context
//   - setting *database.NotificationSetting
func (_e *MockNotificationStore_Expecter) CreateSetting(ctx interface{}, setting interface{}) *MockNotificationStore_CreateSetting_Call {
	return &MockNotificationStore_CreateSetting_Call{Call: _e.mock.On("CreateSetting", ctx, setting)}
}

func (_c *MockNotificationStore_CreateSetting_Call) Run(run func(ctx context.Context, setting *database.NotificationSetting)) *MockNotificationStore_CreateSetting_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*database.NotificationSetting))
	})
	return _c
}

func (_c *MockNotificationStore_CreateSetting_Call) Return(_a0 error) *MockNotificationStore_CreateSetting_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNotificationStore_CreateSetting_Call) RunAndReturn(run func(context.Context, *database.NotificationSetting) error) *MockNotificationStore_CreateSetting_Call {
	_c.Call.Return(run)
	return _c
}

// CreateUserMessageErrorLog provides a mock function with given fields: ctx, msgUUID, userUUID, errorMsg
func (_m *MockNotificationStore) CreateUserMessageErrorLog(ctx context.Context, msgUUID string, userUUID string, errorMsg string) error {
	ret := _m.Called(ctx, msgUUID, userUUID, errorMsg)

	if len(ret) == 0 {
		panic("no return value specified for CreateUserMessageErrorLog")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, msgUUID, userUUID, errorMsg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNotificationStore_CreateUserMessageErrorLog_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUserMessageErrorLog'
type MockNotificationStore_CreateUserMessageErrorLog_Call struct {
	*mock.Call
}

// CreateUserMessageErrorLog is a helper method to define mock.On call
//   - ctx context.Context
//   - msgUUID string
//   - userUUID string
//   - errorMsg string
func (_e *MockNotificationStore_Expecter) CreateUserMessageErrorLog(ctx interface{}, msgUUID interface{}, userUUID interface{}, errorMsg interface{}) *MockNotificationStore_CreateUserMessageErrorLog_Call {
	return &MockNotificationStore_CreateUserMessageErrorLog_Call{Call: _e.mock.On("CreateUserMessageErrorLog", ctx, msgUUID, userUUID, errorMsg)}
}

func (_c *MockNotificationStore_CreateUserMessageErrorLog_Call) Run(run func(ctx context.Context, msgUUID string, userUUID string, errorMsg string)) *MockNotificationStore_CreateUserMessageErrorLog_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockNotificationStore_CreateUserMessageErrorLog_Call) Return(_a0 error) *MockNotificationStore_CreateUserMessageErrorLog_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNotificationStore_CreateUserMessageErrorLog_Call) RunAndReturn(run func(context.Context, string, string, string) error) *MockNotificationStore_CreateUserMessageErrorLog_Call {
	_c.Call.Return(run)
	return _c
}

// CreateUserMessages provides a mock function with given fields: ctx, msgUUID, userUUIDs
func (_m *MockNotificationStore) CreateUserMessages(ctx context.Context, msgUUID string, userUUIDs []string) ([]database.NotificationUserMessageErrorLog, error) {
	ret := _m.Called(ctx, msgUUID, userUUIDs)

	if len(ret) == 0 {
		panic("no return value specified for CreateUserMessages")
	}

	var r0 []database.NotificationUserMessageErrorLog
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) ([]database.NotificationUserMessageErrorLog, error)); ok {
		return rf(ctx, msgUUID, userUUIDs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) []database.NotificationUserMessageErrorLog); ok {
		r0 = rf(ctx, msgUUID, userUUIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]database.NotificationUserMessageErrorLog)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []string) error); ok {
		r1 = rf(ctx, msgUUID, userUUIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNotificationStore_CreateUserMessages_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUserMessages'
type MockNotificationStore_CreateUserMessages_Call struct {
	*mock.Call
}

// CreateUserMessages is a helper method to define mock.On call
//   - ctx context.Context
//   - msgUUID string
//   - userUUIDs []string
func (_e *MockNotificationStore_Expecter) CreateUserMessages(ctx interface{}, msgUUID interface{}, userUUIDs interface{}) *MockNotificationStore_CreateUserMessages_Call {
	return &MockNotificationStore_CreateUserMessages_Call{Call: _e.mock.On("CreateUserMessages", ctx, msgUUID, userUUIDs)}
}

func (_c *MockNotificationStore_CreateUserMessages_Call) Run(run func(ctx context.Context, msgUUID string, userUUIDs []string)) *MockNotificationStore_CreateUserMessages_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]string))
	})
	return _c
}

func (_c *MockNotificationStore_CreateUserMessages_Call) Return(_a0 []database.NotificationUserMessageErrorLog, _a1 error) *MockNotificationStore_CreateUserMessages_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNotificationStore_CreateUserMessages_Call) RunAndReturn(run func(context.Context, string, []string) ([]database.NotificationUserMessageErrorLog, error)) *MockNotificationStore_CreateUserMessages_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllSettings provides a mock function with given fields: ctx
func (_m *MockNotificationStore) GetAllSettings(ctx context.Context) ([]*database.NotificationSetting, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllSettings")
	}

	var r0 []*database.NotificationSetting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*database.NotificationSetting, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*database.NotificationSetting); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*database.NotificationSetting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNotificationStore_GetAllSettings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllSettings'
type MockNotificationStore_GetAllSettings_Call struct {
	*mock.Call
}

// GetAllSettings is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockNotificationStore_Expecter) GetAllSettings(ctx interface{}) *MockNotificationStore_GetAllSettings_Call {
	return &MockNotificationStore_GetAllSettings_Call{Call: _e.mock.On("GetAllSettings", ctx)}
}

func (_c *MockNotificationStore_GetAllSettings_Call) Run(run func(ctx context.Context)) *MockNotificationStore_GetAllSettings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockNotificationStore_GetAllSettings_Call) Return(_a0 []*database.NotificationSetting, _a1 error) *MockNotificationStore_GetAllSettings_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNotificationStore_GetAllSettings_Call) RunAndReturn(run func(context.Context) ([]*database.NotificationSetting, error)) *MockNotificationStore_GetAllSettings_Call {
	_c.Call.Return(run)
	return _c
}

// GetEmails provides a mock function with given fields: ctx, per, page
func (_m *MockNotificationStore) GetEmails(ctx context.Context, per int, page int) ([]string, int, error) {
	ret := _m.Called(ctx, per, page)

	if len(ret) == 0 {
		panic("no return value specified for GetEmails")
	}

	var r0 []string
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) ([]string, int, error)); ok {
		return rf(ctx, per, page)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []string); ok {
		r0 = rf(ctx, per, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) int); ok {
		r1 = rf(ctx, per, page)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, int, int) error); ok {
		r2 = rf(ctx, per, page)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockNotificationStore_GetEmails_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEmails'
type MockNotificationStore_GetEmails_Call struct {
	*mock.Call
}

// GetEmails is a helper method to define mock.On call
//   - ctx context.Context
//   - per int
//   - page int
func (_e *MockNotificationStore_Expecter) GetEmails(ctx interface{}, per interface{}, page interface{}) *MockNotificationStore_GetEmails_Call {
	return &MockNotificationStore_GetEmails_Call{Call: _e.mock.On("GetEmails", ctx, per, page)}
}

func (_c *MockNotificationStore_GetEmails_Call) Run(run func(ctx context.Context, per int, page int)) *MockNotificationStore_GetEmails_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(int))
	})
	return _c
}

func (_c *MockNotificationStore_GetEmails_Call) Return(_a0 []string, _a1 int, _a2 error) *MockNotificationStore_GetEmails_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockNotificationStore_GetEmails_Call) RunAndReturn(run func(context.Context, int, int) ([]string, int, error)) *MockNotificationStore_GetEmails_Call {
	_c.Call.Return(run)
	return _c
}

// GetSetting provides a mock function with given fields: ctx, uid
func (_m *MockNotificationStore) GetSetting(ctx context.Context, uid string) (*database.NotificationSetting, error) {
	ret := _m.Called(ctx, uid)

	if len(ret) == 0 {
		panic("no return value specified for GetSetting")
	}

	var r0 *database.NotificationSetting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*database.NotificationSetting, error)); ok {
		return rf(ctx, uid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *database.NotificationSetting); ok {
		r0 = rf(ctx, uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.NotificationSetting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNotificationStore_GetSetting_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSetting'
type MockNotificationStore_GetSetting_Call struct {
	*mock.Call
}

// GetSetting is a helper method to define mock.On call
//   - ctx context.Context
//   - uid string
func (_e *MockNotificationStore_Expecter) GetSetting(ctx interface{}, uid interface{}) *MockNotificationStore_GetSetting_Call {
	return &MockNotificationStore_GetSetting_Call{Call: _e.mock.On("GetSetting", ctx, uid)}
}

func (_c *MockNotificationStore_GetSetting_Call) Run(run func(ctx context.Context, uid string)) *MockNotificationStore_GetSetting_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockNotificationStore_GetSetting_Call) Return(_a0 *database.NotificationSetting, _a1 error) *MockNotificationStore_GetSetting_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNotificationStore_GetSetting_Call) RunAndReturn(run func(context.Context, string) (*database.NotificationSetting, error)) *MockNotificationStore_GetSetting_Call {
	_c.Call.Return(run)
	return _c
}

// GetSettingByUserUUIDs provides a mock function with given fields: ctx, userUUID
func (_m *MockNotificationStore) GetSettingByUserUUIDs(ctx context.Context, userUUID []string) ([]*database.NotificationSetting, error) {
	ret := _m.Called(ctx, userUUID)

	if len(ret) == 0 {
		panic("no return value specified for GetSettingByUserUUIDs")
	}

	var r0 []*database.NotificationSetting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) ([]*database.NotificationSetting, error)); ok {
		return rf(ctx, userUUID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) []*database.NotificationSetting); ok {
		r0 = rf(ctx, userUUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*database.NotificationSetting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, userUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNotificationStore_GetSettingByUserUUIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSettingByUserUUIDs'
type MockNotificationStore_GetSettingByUserUUIDs_Call struct {
	*mock.Call
}

// GetSettingByUserUUIDs is a helper method to define mock.On call
//   - ctx context.Context
//   - userUUID []string
func (_e *MockNotificationStore_Expecter) GetSettingByUserUUIDs(ctx interface{}, userUUID interface{}) *MockNotificationStore_GetSettingByUserUUIDs_Call {
	return &MockNotificationStore_GetSettingByUserUUIDs_Call{Call: _e.mock.On("GetSettingByUserUUIDs", ctx, userUUID)}
}

func (_c *MockNotificationStore_GetSettingByUserUUIDs_Call) Run(run func(ctx context.Context, userUUID []string)) *MockNotificationStore_GetSettingByUserUUIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string))
	})
	return _c
}

func (_c *MockNotificationStore_GetSettingByUserUUIDs_Call) Return(_a0 []*database.NotificationSetting, _a1 error) *MockNotificationStore_GetSettingByUserUUIDs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNotificationStore_GetSettingByUserUUIDs_Call) RunAndReturn(run func(context.Context, []string) ([]*database.NotificationSetting, error)) *MockNotificationStore_GetSettingByUserUUIDs_Call {
	_c.Call.Return(run)
	return _c
}

// GetUnNotifiedMessages provides a mock function with given fields: ctx, userUUID, limit
func (_m *MockNotificationStore) GetUnNotifiedMessages(ctx context.Context, userUUID string, limit int) ([]database.NotificationUserMessageView, int, error) {
	ret := _m.Called(ctx, userUUID, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetUnNotifiedMessages")
	}

	var r0 []database.NotificationUserMessageView
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int) ([]database.NotificationUserMessageView, int, error)); ok {
		return rf(ctx, userUUID, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int) []database.NotificationUserMessageView); ok {
		r0 = rf(ctx, userUUID, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]database.NotificationUserMessageView)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int) int); ok {
		r1 = rf(ctx, userUUID, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, int) error); ok {
		r2 = rf(ctx, userUUID, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockNotificationStore_GetUnNotifiedMessages_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUnNotifiedMessages'
type MockNotificationStore_GetUnNotifiedMessages_Call struct {
	*mock.Call
}

// GetUnNotifiedMessages is a helper method to define mock.On call
//   - ctx context.Context
//   - userUUID string
//   - limit int
func (_e *MockNotificationStore_Expecter) GetUnNotifiedMessages(ctx interface{}, userUUID interface{}, limit interface{}) *MockNotificationStore_GetUnNotifiedMessages_Call {
	return &MockNotificationStore_GetUnNotifiedMessages_Call{Call: _e.mock.On("GetUnNotifiedMessages", ctx, userUUID, limit)}
}

func (_c *MockNotificationStore_GetUnNotifiedMessages_Call) Run(run func(ctx context.Context, userUUID string, limit int)) *MockNotificationStore_GetUnNotifiedMessages_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int))
	})
	return _c
}

func (_c *MockNotificationStore_GetUnNotifiedMessages_Call) Return(_a0 []database.NotificationUserMessageView, _a1 int, _a2 error) *MockNotificationStore_GetUnNotifiedMessages_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockNotificationStore_GetUnNotifiedMessages_Call) RunAndReturn(run func(context.Context, string, int) ([]database.NotificationUserMessageView, int, error)) *MockNotificationStore_GetUnNotifiedMessages_Call {
	_c.Call.Return(run)
	return _c
}

// GetUnreadCount provides a mock function with given fields: ctx, uid
func (_m *MockNotificationStore) GetUnreadCount(ctx context.Context, uid string) (int64, error) {
	ret := _m.Called(ctx, uid)

	if len(ret) == 0 {
		panic("no return value specified for GetUnreadCount")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int64, error)); ok {
		return rf(ctx, uid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, uid)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNotificationStore_GetUnreadCount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUnreadCount'
type MockNotificationStore_GetUnreadCount_Call struct {
	*mock.Call
}

// GetUnreadCount is a helper method to define mock.On call
//   - ctx context.Context
//   - uid string
func (_e *MockNotificationStore_Expecter) GetUnreadCount(ctx interface{}, uid interface{}) *MockNotificationStore_GetUnreadCount_Call {
	return &MockNotificationStore_GetUnreadCount_Call{Call: _e.mock.On("GetUnreadCount", ctx, uid)}
}

func (_c *MockNotificationStore_GetUnreadCount_Call) Run(run func(ctx context.Context, uid string)) *MockNotificationStore_GetUnreadCount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockNotificationStore_GetUnreadCount_Call) Return(_a0 int64, _a1 error) *MockNotificationStore_GetUnreadCount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNotificationStore_GetUnreadCount_Call) RunAndReturn(run func(context.Context, string) (int64, error)) *MockNotificationStore_GetUnreadCount_Call {
	_c.Call.Return(run)
	return _c
}

// IsNotificationMessageExists provides a mock function with given fields: ctx, msgUUID
func (_m *MockNotificationStore) IsNotificationMessageExists(ctx context.Context, msgUUID string) (bool, error) {
	ret := _m.Called(ctx, msgUUID)

	if len(ret) == 0 {
		panic("no return value specified for IsNotificationMessageExists")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, msgUUID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, msgUUID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, msgUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNotificationStore_IsNotificationMessageExists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsNotificationMessageExists'
type MockNotificationStore_IsNotificationMessageExists_Call struct {
	*mock.Call
}

// IsNotificationMessageExists is a helper method to define mock.On call
//   - ctx context.Context
//   - msgUUID string
func (_e *MockNotificationStore_Expecter) IsNotificationMessageExists(ctx interface{}, msgUUID interface{}) *MockNotificationStore_IsNotificationMessageExists_Call {
	return &MockNotificationStore_IsNotificationMessageExists_Call{Call: _e.mock.On("IsNotificationMessageExists", ctx, msgUUID)}
}

func (_c *MockNotificationStore_IsNotificationMessageExists_Call) Run(run func(ctx context.Context, msgUUID string)) *MockNotificationStore_IsNotificationMessageExists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockNotificationStore_IsNotificationMessageExists_Call) Return(_a0 bool, _a1 error) *MockNotificationStore_IsNotificationMessageExists_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNotificationStore_IsNotificationMessageExists_Call) RunAndReturn(run func(context.Context, string) (bool, error)) *MockNotificationStore_IsNotificationMessageExists_Call {
	_c.Call.Return(run)
	return _c
}

// ListNotificationMessages provides a mock function with given fields: ctx, params
func (_m *MockNotificationStore) ListNotificationMessages(ctx context.Context, params database.ListNotificationsParams) ([]database.NotificationUserMessageView, int, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for ListNotificationMessages")
	}

	var r0 []database.NotificationUserMessageView
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, database.ListNotificationsParams) ([]database.NotificationUserMessageView, int, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, database.ListNotificationsParams) []database.NotificationUserMessageView); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]database.NotificationUserMessageView)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, database.ListNotificationsParams) int); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, database.ListNotificationsParams) error); ok {
		r2 = rf(ctx, params)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockNotificationStore_ListNotificationMessages_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListNotificationMessages'
type MockNotificationStore_ListNotificationMessages_Call struct {
	*mock.Call
}

// ListNotificationMessages is a helper method to define mock.On call
//   - ctx context.Context
//   - params database.ListNotificationsParams
func (_e *MockNotificationStore_Expecter) ListNotificationMessages(ctx interface{}, params interface{}) *MockNotificationStore_ListNotificationMessages_Call {
	return &MockNotificationStore_ListNotificationMessages_Call{Call: _e.mock.On("ListNotificationMessages", ctx, params)}
}

func (_c *MockNotificationStore_ListNotificationMessages_Call) Run(run func(ctx context.Context, params database.ListNotificationsParams)) *MockNotificationStore_ListNotificationMessages_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.ListNotificationsParams))
	})
	return _c
}

func (_c *MockNotificationStore_ListNotificationMessages_Call) Return(_a0 []database.NotificationUserMessageView, _a1 int, _a2 error) *MockNotificationStore_ListNotificationMessages_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockNotificationStore_ListNotificationMessages_Call) RunAndReturn(run func(context.Context, database.ListNotificationsParams) ([]database.NotificationUserMessageView, int, error)) *MockNotificationStore_ListNotificationMessages_Call {
	_c.Call.Return(run)
	return _c
}

// MarkAllAsRead provides a mock function with given fields: ctx, uid
func (_m *MockNotificationStore) MarkAllAsRead(ctx context.Context, uid string) error {
	ret := _m.Called(ctx, uid)

	if len(ret) == 0 {
		panic("no return value specified for MarkAllAsRead")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNotificationStore_MarkAllAsRead_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarkAllAsRead'
type MockNotificationStore_MarkAllAsRead_Call struct {
	*mock.Call
}

// MarkAllAsRead is a helper method to define mock.On call
//   - ctx context.Context
//   - uid string
func (_e *MockNotificationStore_Expecter) MarkAllAsRead(ctx interface{}, uid interface{}) *MockNotificationStore_MarkAllAsRead_Call {
	return &MockNotificationStore_MarkAllAsRead_Call{Call: _e.mock.On("MarkAllAsRead", ctx, uid)}
}

func (_c *MockNotificationStore_MarkAllAsRead_Call) Run(run func(ctx context.Context, uid string)) *MockNotificationStore_MarkAllAsRead_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockNotificationStore_MarkAllAsRead_Call) Return(_a0 error) *MockNotificationStore_MarkAllAsRead_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNotificationStore_MarkAllAsRead_Call) RunAndReturn(run func(context.Context, string) error) *MockNotificationStore_MarkAllAsRead_Call {
	_c.Call.Return(run)
	return _c
}

// MarkAsNotified provides a mock function with given fields: ctx, userUUID, msgUUIDs
func (_m *MockNotificationStore) MarkAsNotified(ctx context.Context, userUUID string, msgUUIDs []string) error {
	ret := _m.Called(ctx, userUUID, msgUUIDs)

	if len(ret) == 0 {
		panic("no return value specified for MarkAsNotified")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) error); ok {
		r0 = rf(ctx, userUUID, msgUUIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNotificationStore_MarkAsNotified_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarkAsNotified'
type MockNotificationStore_MarkAsNotified_Call struct {
	*mock.Call
}

// MarkAsNotified is a helper method to define mock.On call
//   - ctx context.Context
//   - userUUID string
//   - msgUUIDs []string
func (_e *MockNotificationStore_Expecter) MarkAsNotified(ctx interface{}, userUUID interface{}, msgUUIDs interface{}) *MockNotificationStore_MarkAsNotified_Call {
	return &MockNotificationStore_MarkAsNotified_Call{Call: _e.mock.On("MarkAsNotified", ctx, userUUID, msgUUIDs)}
}

func (_c *MockNotificationStore_MarkAsNotified_Call) Run(run func(ctx context.Context, userUUID string, msgUUIDs []string)) *MockNotificationStore_MarkAsNotified_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]string))
	})
	return _c
}

func (_c *MockNotificationStore_MarkAsNotified_Call) Return(_a0 error) *MockNotificationStore_MarkAsNotified_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNotificationStore_MarkAsNotified_Call) RunAndReturn(run func(context.Context, string, []string) error) *MockNotificationStore_MarkAsNotified_Call {
	_c.Call.Return(run)
	return _c
}

// MarkAsRead provides a mock function with given fields: ctx, uid, ids
func (_m *MockNotificationStore) MarkAsRead(ctx context.Context, uid string, ids []int64) error {
	ret := _m.Called(ctx, uid, ids)

	if len(ret) == 0 {
		panic("no return value specified for MarkAsRead")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []int64) error); ok {
		r0 = rf(ctx, uid, ids)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNotificationStore_MarkAsRead_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarkAsRead'
type MockNotificationStore_MarkAsRead_Call struct {
	*mock.Call
}

// MarkAsRead is a helper method to define mock.On call
//   - ctx context.Context
//   - uid string
//   - ids []int64
func (_e *MockNotificationStore_Expecter) MarkAsRead(ctx interface{}, uid interface{}, ids interface{}) *MockNotificationStore_MarkAsRead_Call {
	return &MockNotificationStore_MarkAsRead_Call{Call: _e.mock.On("MarkAsRead", ctx, uid, ids)}
}

func (_c *MockNotificationStore_MarkAsRead_Call) Run(run func(ctx context.Context, uid string, ids []int64)) *MockNotificationStore_MarkAsRead_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]int64))
	})
	return _c
}

func (_c *MockNotificationStore_MarkAsRead_Call) Return(_a0 error) *MockNotificationStore_MarkAsRead_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNotificationStore_MarkAsRead_Call) RunAndReturn(run func(context.Context, string, []int64) error) *MockNotificationStore_MarkAsRead_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateSetting provides a mock function with given fields: ctx, setting
func (_m *MockNotificationStore) UpdateSetting(ctx context.Context, setting *database.NotificationSetting) error {
	ret := _m.Called(ctx, setting)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSetting")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *database.NotificationSetting) error); ok {
		r0 = rf(ctx, setting)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNotificationStore_UpdateSetting_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSetting'
type MockNotificationStore_UpdateSetting_Call struct {
	*mock.Call
}

// UpdateSetting is a helper method to define mock.On call
//   - ctx context.Context
//   - setting *database.NotificationSetting
func (_e *MockNotificationStore_Expecter) UpdateSetting(ctx interface{}, setting interface{}) *MockNotificationStore_UpdateSetting_Call {
	return &MockNotificationStore_UpdateSetting_Call{Call: _e.mock.On("UpdateSetting", ctx, setting)}
}

func (_c *MockNotificationStore_UpdateSetting_Call) Run(run func(ctx context.Context, setting *database.NotificationSetting)) *MockNotificationStore_UpdateSetting_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*database.NotificationSetting))
	})
	return _c
}

func (_c *MockNotificationStore_UpdateSetting_Call) Return(_a0 error) *MockNotificationStore_UpdateSetting_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNotificationStore_UpdateSetting_Call) RunAndReturn(run func(context.Context, *database.NotificationSetting) error) *MockNotificationStore_UpdateSetting_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockNotificationStore creates a new instance of MockNotificationStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockNotificationStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockNotificationStore {
	mock := &MockNotificationStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
