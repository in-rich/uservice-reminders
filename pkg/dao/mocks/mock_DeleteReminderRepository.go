// Code generated by mockery v2.43.2. DO NOT EDIT.

package daomocks

import (
	context "context"

	entities "github.com/in-rich/uservice-reminders/pkg/entities"

	mock "github.com/stretchr/testify/mock"
)

// MockDeleteReminderRepository is an autogenerated mock type for the DeleteReminderRepository type
type MockDeleteReminderRepository struct {
	mock.Mock
}

type MockDeleteReminderRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDeleteReminderRepository) EXPECT() *MockDeleteReminderRepository_Expecter {
	return &MockDeleteReminderRepository_Expecter{mock: &_m.Mock}
}

// DeleteReminder provides a mock function with given fields: ctx, author, target, publicIdentifier
func (_m *MockDeleteReminderRepository) DeleteReminder(ctx context.Context, author string, target entities.Target, publicIdentifier string) (*entities.Reminder, error) {
	ret := _m.Called(ctx, author, target, publicIdentifier)

	if len(ret) == 0 {
		panic("no return value specified for DeleteReminder")
	}

	var r0 *entities.Reminder
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, entities.Target, string) (*entities.Reminder, error)); ok {
		return rf(ctx, author, target, publicIdentifier)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, entities.Target, string) *entities.Reminder); ok {
		r0 = rf(ctx, author, target, publicIdentifier)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Reminder)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, entities.Target, string) error); ok {
		r1 = rf(ctx, author, target, publicIdentifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDeleteReminderRepository_DeleteReminder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteReminder'
type MockDeleteReminderRepository_DeleteReminder_Call struct {
	*mock.Call
}

// DeleteReminder is a helper method to define mock.On call
//   - ctx context.Context
//   - author string
//   - target entities.Target
//   - publicIdentifier string
func (_e *MockDeleteReminderRepository_Expecter) DeleteReminder(ctx interface{}, author interface{}, target interface{}, publicIdentifier interface{}) *MockDeleteReminderRepository_DeleteReminder_Call {
	return &MockDeleteReminderRepository_DeleteReminder_Call{Call: _e.mock.On("DeleteReminder", ctx, author, target, publicIdentifier)}
}

func (_c *MockDeleteReminderRepository_DeleteReminder_Call) Run(run func(ctx context.Context, author string, target entities.Target, publicIdentifier string)) *MockDeleteReminderRepository_DeleteReminder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(entities.Target), args[3].(string))
	})
	return _c
}

func (_c *MockDeleteReminderRepository_DeleteReminder_Call) Return(_a0 *entities.Reminder, _a1 error) *MockDeleteReminderRepository_DeleteReminder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDeleteReminderRepository_DeleteReminder_Call) RunAndReturn(run func(context.Context, string, entities.Target, string) (*entities.Reminder, error)) *MockDeleteReminderRepository_DeleteReminder_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDeleteReminderRepository creates a new instance of MockDeleteReminderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDeleteReminderRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDeleteReminderRepository {
	mock := &MockDeleteReminderRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
