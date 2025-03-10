// Code generated by mockery v2.43.2. DO NOT EDIT.

package servicesmocks

import (
	context "context"

	models "github.com/in-rich/uservice-reminders/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// MockGetReminderByIDService is an autogenerated mock type for the GetReminderByIDService type
type MockGetReminderByIDService struct {
	mock.Mock
}

type MockGetReminderByIDService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockGetReminderByIDService) EXPECT() *MockGetReminderByIDService_Expecter {
	return &MockGetReminderByIDService_Expecter{mock: &_m.Mock}
}

// Exec provides a mock function with given fields: ctx, selector
func (_m *MockGetReminderByIDService) Exec(ctx context.Context, selector *models.GetReminderByID) (*models.Reminder, error) {
	ret := _m.Called(ctx, selector)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 *models.Reminder
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.GetReminderByID) (*models.Reminder, error)); ok {
		return rf(ctx, selector)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.GetReminderByID) *models.Reminder); ok {
		r0 = rf(ctx, selector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Reminder)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.GetReminderByID) error); ok {
		r1 = rf(ctx, selector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGetReminderByIDService_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockGetReminderByIDService_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - ctx context.Context
//   - selector *models.GetReminderByID
func (_e *MockGetReminderByIDService_Expecter) Exec(ctx interface{}, selector interface{}) *MockGetReminderByIDService_Exec_Call {
	return &MockGetReminderByIDService_Exec_Call{Call: _e.mock.On("Exec", ctx, selector)}
}

func (_c *MockGetReminderByIDService_Exec_Call) Run(run func(ctx context.Context, selector *models.GetReminderByID)) *MockGetReminderByIDService_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.GetReminderByID))
	})
	return _c
}

func (_c *MockGetReminderByIDService_Exec_Call) Return(_a0 *models.Reminder, _a1 error) *MockGetReminderByIDService_Exec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGetReminderByIDService_Exec_Call) RunAndReturn(run func(context.Context, *models.GetReminderByID) (*models.Reminder, error)) *MockGetReminderByIDService_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockGetReminderByIDService creates a new instance of MockGetReminderByIDService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGetReminderByIDService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGetReminderByIDService {
	mock := &MockGetReminderByIDService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
