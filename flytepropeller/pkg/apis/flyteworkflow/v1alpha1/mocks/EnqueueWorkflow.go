// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// EnqueueWorkflow is an autogenerated mock type for the EnqueueWorkflow type
type EnqueueWorkflow struct {
	mock.Mock
}

type EnqueueWorkflow_Expecter struct {
	mock *mock.Mock
}

func (_m *EnqueueWorkflow) EXPECT() *EnqueueWorkflow_Expecter {
	return &EnqueueWorkflow_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: workflowID
func (_m *EnqueueWorkflow) Execute(workflowID string) {
	_m.Called(workflowID)
}

// EnqueueWorkflow_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type EnqueueWorkflow_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - workflowID string
func (_e *EnqueueWorkflow_Expecter) Execute(workflowID interface{}) *EnqueueWorkflow_Execute_Call {
	return &EnqueueWorkflow_Execute_Call{Call: _e.mock.On("Execute", workflowID)}
}

func (_c *EnqueueWorkflow_Execute_Call) Run(run func(workflowID string)) *EnqueueWorkflow_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *EnqueueWorkflow_Execute_Call) Return() *EnqueueWorkflow_Execute_Call {
	_c.Call.Return()
	return _c
}

func (_c *EnqueueWorkflow_Execute_Call) RunAndReturn(run func(string)) *EnqueueWorkflow_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewEnqueueWorkflow creates a new instance of EnqueueWorkflow. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEnqueueWorkflow(t interface {
	mock.TestingT
	Cleanup(func())
}) *EnqueueWorkflow {
	mock := &EnqueueWorkflow{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
