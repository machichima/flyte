// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	core "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	mock "github.com/stretchr/testify/mock"
)

// Task is an autogenerated mock type for the Task type
type Task struct {
	mock.Mock
}

type Task_Expecter struct {
	mock *mock.Mock
}

func (_m *Task) EXPECT() *Task_Expecter {
	return &Task_Expecter{mock: &_m.Mock}
}

// GetCoreTask provides a mock function with given fields:
func (_m *Task) GetCoreTask() *core.TaskTemplate {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetCoreTask")
	}

	var r0 *core.TaskTemplate
	if rf, ok := ret.Get(0).(func() *core.TaskTemplate); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.TaskTemplate)
		}
	}

	return r0
}

// Task_GetCoreTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCoreTask'
type Task_GetCoreTask_Call struct {
	*mock.Call
}

// GetCoreTask is a helper method to define mock.On call
func (_e *Task_Expecter) GetCoreTask() *Task_GetCoreTask_Call {
	return &Task_GetCoreTask_Call{Call: _e.mock.On("GetCoreTask")}
}

func (_c *Task_GetCoreTask_Call) Run(run func()) *Task_GetCoreTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Task_GetCoreTask_Call) Return(_a0 *core.TaskTemplate) *Task_GetCoreTask_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Task_GetCoreTask_Call) RunAndReturn(run func() *core.TaskTemplate) *Task_GetCoreTask_Call {
	_c.Call.Return(run)
	return _c
}

// GetID provides a mock function with given fields:
func (_m *Task) GetID() *core.Identifier {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetID")
	}

	var r0 *core.Identifier
	if rf, ok := ret.Get(0).(func() *core.Identifier); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Identifier)
		}
	}

	return r0
}

// Task_GetID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetID'
type Task_GetID_Call struct {
	*mock.Call
}

// GetID is a helper method to define mock.On call
func (_e *Task_Expecter) GetID() *Task_GetID_Call {
	return &Task_GetID_Call{Call: _e.mock.On("GetID")}
}

func (_c *Task_GetID_Call) Run(run func()) *Task_GetID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Task_GetID_Call) Return(_a0 *core.Identifier) *Task_GetID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Task_GetID_Call) RunAndReturn(run func() *core.Identifier) *Task_GetID_Call {
	_c.Call.Return(run)
	return _c
}

// GetInterface provides a mock function with given fields:
func (_m *Task) GetInterface() *core.TypedInterface {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetInterface")
	}

	var r0 *core.TypedInterface
	if rf, ok := ret.Get(0).(func() *core.TypedInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.TypedInterface)
		}
	}

	return r0
}

// Task_GetInterface_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInterface'
type Task_GetInterface_Call struct {
	*mock.Call
}

// GetInterface is a helper method to define mock.On call
func (_e *Task_Expecter) GetInterface() *Task_GetInterface_Call {
	return &Task_GetInterface_Call{Call: _e.mock.On("GetInterface")}
}

func (_c *Task_GetInterface_Call) Run(run func()) *Task_GetInterface_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Task_GetInterface_Call) Return(_a0 *core.TypedInterface) *Task_GetInterface_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Task_GetInterface_Call) RunAndReturn(run func() *core.TypedInterface) *Task_GetInterface_Call {
	_c.Call.Return(run)
	return _c
}

// NewTask creates a new instance of Task. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTask(t interface {
	mock.TestingT
	Cleanup(func())
}) *Task {
	mock := &Task{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
