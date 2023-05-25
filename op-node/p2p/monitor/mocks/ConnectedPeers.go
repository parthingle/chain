// Code generated by mockery v2.28.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	peer "github.com/libp2p/go-libp2p/core/peer"
)

// ConnectedPeers is an autogenerated mock type for the ConnectedPeers type
type ConnectedPeers struct {
	mock.Mock
}

type ConnectedPeers_Expecter struct {
	mock *mock.Mock
}

func (_m *ConnectedPeers) EXPECT() *ConnectedPeers_Expecter {
	return &ConnectedPeers_Expecter{mock: &_m.Mock}
}

// ClosePeer provides a mock function with given fields: _a0
func (_m *ConnectedPeers) ClosePeer(_a0 peer.ID) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(peer.ID) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConnectedPeers_ClosePeer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClosePeer'
type ConnectedPeers_ClosePeer_Call struct {
	*mock.Call
}

// ClosePeer is a helper method to define mock.On call
//   - _a0 peer.ID
func (_e *ConnectedPeers_Expecter) ClosePeer(_a0 interface{}) *ConnectedPeers_ClosePeer_Call {
	return &ConnectedPeers_ClosePeer_Call{Call: _e.mock.On("ClosePeer", _a0)}
}

func (_c *ConnectedPeers_ClosePeer_Call) Run(run func(_a0 peer.ID)) *ConnectedPeers_ClosePeer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(peer.ID))
	})
	return _c
}

func (_c *ConnectedPeers_ClosePeer_Call) Return(_a0 error) *ConnectedPeers_ClosePeer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConnectedPeers_ClosePeer_Call) RunAndReturn(run func(peer.ID) error) *ConnectedPeers_ClosePeer_Call {
	_c.Call.Return(run)
	return _c
}

// Peers provides a mock function with given fields:
func (_m *ConnectedPeers) Peers() []peer.ID {
	ret := _m.Called()

	var r0 []peer.ID
	if rf, ok := ret.Get(0).(func() []peer.ID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]peer.ID)
		}
	}

	return r0
}

// ConnectedPeers_Peers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Peers'
type ConnectedPeers_Peers_Call struct {
	*mock.Call
}

// Peers is a helper method to define mock.On call
func (_e *ConnectedPeers_Expecter) Peers() *ConnectedPeers_Peers_Call {
	return &ConnectedPeers_Peers_Call{Call: _e.mock.On("Peers")}
}

func (_c *ConnectedPeers_Peers_Call) Run(run func()) *ConnectedPeers_Peers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ConnectedPeers_Peers_Call) Return(_a0 []peer.ID) *ConnectedPeers_Peers_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConnectedPeers_Peers_Call) RunAndReturn(run func() []peer.ID) *ConnectedPeers_Peers_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewConnectedPeers interface {
	mock.TestingT
	Cleanup(func())
}

// NewConnectedPeers creates a new instance of ConnectedPeers. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConnectedPeers(t mockConstructorTestingTNewConnectedPeers) *ConnectedPeers {
	mock := &ConnectedPeers{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
