// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	types "github.com/smartcontractkit/mcms/types"
)

// Simulator is an autogenerated mock type for the Simulator type
type Simulator struct {
	mock.Mock
}

type Simulator_Expecter struct {
	mock *mock.Mock
}

func (_m *Simulator) EXPECT() *Simulator_Expecter {
	return &Simulator_Expecter{mock: &_m.Mock}
}

// SimulateOperation provides a mock function with given fields: ctx, metadata, operation
func (_m *Simulator) SimulateOperation(ctx context.Context, metadata types.ChainMetadata, operation types.Operation) error {
	ret := _m.Called(ctx, metadata, operation)

	if len(ret) == 0 {
		panic("no return value specified for SimulateOperation")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.ChainMetadata, types.Operation) error); ok {
		r0 = rf(ctx, metadata, operation)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Simulator_SimulateOperation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SimulateOperation'
type Simulator_SimulateOperation_Call struct {
	*mock.Call
}

// SimulateOperation is a helper method to define mock.On call
//   - ctx context.Context
//   - metadata types.ChainMetadata
//   - operation types.Operation
func (_e *Simulator_Expecter) SimulateOperation(ctx interface{}, metadata interface{}, operation interface{}) *Simulator_SimulateOperation_Call {
	return &Simulator_SimulateOperation_Call{Call: _e.mock.On("SimulateOperation", ctx, metadata, operation)}
}

func (_c *Simulator_SimulateOperation_Call) Run(run func(ctx context.Context, metadata types.ChainMetadata, operation types.Operation)) *Simulator_SimulateOperation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.ChainMetadata), args[2].(types.Operation))
	})
	return _c
}

func (_c *Simulator_SimulateOperation_Call) Return(_a0 error) *Simulator_SimulateOperation_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Simulator_SimulateOperation_Call) RunAndReturn(run func(context.Context, types.ChainMetadata, types.Operation) error) *Simulator_SimulateOperation_Call {
	_c.Call.Return(run)
	return _c
}

// SimulateSetRoot provides a mock function with given fields: ctx, originCaller, metadata, proof, root, validUntil, sortedSignatures
func (_m *Simulator) SimulateSetRoot(ctx context.Context, originCaller string, metadata types.ChainMetadata, proof []common.Hash, root [32]byte, validUntil uint32, sortedSignatures []types.Signature) error {
	ret := _m.Called(ctx, originCaller, metadata, proof, root, validUntil, sortedSignatures)

	if len(ret) == 0 {
		panic("no return value specified for SimulateSetRoot")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ChainMetadata, []common.Hash, [32]byte, uint32, []types.Signature) error); ok {
		r0 = rf(ctx, originCaller, metadata, proof, root, validUntil, sortedSignatures)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Simulator_SimulateSetRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SimulateSetRoot'
type Simulator_SimulateSetRoot_Call struct {
	*mock.Call
}

// SimulateSetRoot is a helper method to define mock.On call
//   - ctx context.Context
//   - originCaller string
//   - metadata types.ChainMetadata
//   - proof []common.Hash
//   - root [32]byte
//   - validUntil uint32
//   - sortedSignatures []types.Signature
func (_e *Simulator_Expecter) SimulateSetRoot(ctx interface{}, originCaller interface{}, metadata interface{}, proof interface{}, root interface{}, validUntil interface{}, sortedSignatures interface{}) *Simulator_SimulateSetRoot_Call {
	return &Simulator_SimulateSetRoot_Call{Call: _e.mock.On("SimulateSetRoot", ctx, originCaller, metadata, proof, root, validUntil, sortedSignatures)}
}

func (_c *Simulator_SimulateSetRoot_Call) Run(run func(ctx context.Context, originCaller string, metadata types.ChainMetadata, proof []common.Hash, root [32]byte, validUntil uint32, sortedSignatures []types.Signature)) *Simulator_SimulateSetRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(types.ChainMetadata), args[3].([]common.Hash), args[4].([32]byte), args[5].(uint32), args[6].([]types.Signature))
	})
	return _c
}

func (_c *Simulator_SimulateSetRoot_Call) Return(_a0 error) *Simulator_SimulateSetRoot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Simulator_SimulateSetRoot_Call) RunAndReturn(run func(context.Context, string, types.ChainMetadata, []common.Hash, [32]byte, uint32, []types.Signature) error) *Simulator_SimulateSetRoot_Call {
	_c.Call.Return(run)
	return _c
}

// NewSimulator creates a new instance of Simulator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSimulator(t interface {
	mock.TestingT
	Cleanup(func())
}) *Simulator {
	mock := &Simulator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
