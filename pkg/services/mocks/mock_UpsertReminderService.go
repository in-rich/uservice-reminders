// Code generated by mockery v2.43.2. DO NOT EDIT.

package servicesmocks

import (
	context "context"

	models "github.com/in-rich/uservice-reminders/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// MockUpsertReminderService is an autogenerated mock type for the UpsertReminderService type
type MockUpsertReminderService struct {
	mock.Mock
}

type MockUpsertReminderService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUpsertReminderService) EXPECT() *MockUpsertReminderService_Expecter {
	return &MockUpsertReminderService_Expecter{mock: &_m.Mock}
}

// Exec provides a mock function with given fields: ctx, reminder
func (_m *MockUpsertReminderService) Exec(ctx context.Context, reminder *models.UpsertReminder) (*models.Reminder, string, error) {
	ret := _m.Called(ctx, reminder)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 *models.Reminder
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.UpsertReminder) (*models.Reminder, string, error)); ok {
		return rf(ctx, reminder)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.UpsertReminder) *models.Reminder); ok {
		r0 = rf(ctx, reminder)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Reminder)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.UpsertReminder) string); ok {
		r1 = rf(ctx, reminder)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *models.UpsertReminder) error); ok {
		r2 = rf(ctx, reminder)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockUpsertReminderService_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockUpsertReminderService_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - ctx context.Context
//   - reminder *models.UpsertReminder
func (_e *MockUpsertReminderService_Expecter) Exec(ctx interface{}, reminder interface{}) *MockUpsertReminderService_Exec_Call {
	return &MockUpsertReminderService_Exec_Call{Call: _e.mock.On("Exec", ctx, reminder)}
}

func (_c *MockUpsertReminderService_Exec_Call) Run(run func(ctx context.Context, reminder *models.UpsertReminder)) *MockUpsertReminderService_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.UpsertReminder))
	})
	return _c
}

func (_c *MockUpsertReminderService_Exec_Call) Return(_a0 *models.Reminder, _a1 string, _a2 error) *MockUpsertReminderService_Exec_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockUpsertReminderService_Exec_Call) RunAndReturn(run func(context.Context, *models.UpsertReminder) (*models.Reminder, string, error)) *MockUpsertReminderService_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUpsertReminderService creates a new instance of MockUpsertReminderService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUpsertReminderService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUpsertReminderService {
	mock := &MockUpsertReminderService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
