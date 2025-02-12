// Code generated by MockGen. DO NOT EDIT.
// Source: x/crisis/types/expected_keepers.go
//
// Generated by this command:
//
//	mockgen -source=x/crisis/types/expected_keepers.go -package testutil -destination x/crisis/testutil/expected_keepers_mocks.go
//

// Package testutil is a generated GoMock package.
package testutil

import (
	context "context"
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	gomock "go.uber.org/mock/gomock"
)

// MockSupplyKeeper is a mock of SupplyKeeper interface.
type MockSupplyKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockSupplyKeeperMockRecorder
	isgomock struct{}
}

// MockSupplyKeeperMockRecorder is the mock recorder for MockSupplyKeeper.
type MockSupplyKeeperMockRecorder struct {
	mock *MockSupplyKeeper
}

// NewMockSupplyKeeper creates a new mock instance.
func NewMockSupplyKeeper(ctrl *gomock.Controller) *MockSupplyKeeper {
	mock := &MockSupplyKeeper{ctrl: ctrl}
	mock.recorder = &MockSupplyKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSupplyKeeper) EXPECT() *MockSupplyKeeperMockRecorder {
	return m.recorder
}

// SendCoinsFromAccountToModule mocks base method.
func (m *MockSupplyKeeper) SendCoinsFromAccountToModule(ctx context.Context, senderAddr types.AccAddress, recipientModule string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromAccountToModule indicates an expected call of SendCoinsFromAccountToModule.
func (mr *MockSupplyKeeperMockRecorder) SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromAccountToModule", reflect.TypeOf((*MockSupplyKeeper)(nil).SendCoinsFromAccountToModule), ctx, senderAddr, recipientModule, amt)
}
