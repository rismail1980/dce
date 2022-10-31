// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import account "github.com/Optum/dce/pkg/account"
import mock "github.com/stretchr/testify/mock"

// SingleReader is an autogenerated mock type for the SingleReader type
type SingleReader struct {
	mock.Mock
}

// Get provides a mock function with given fields: ID
func (_m *SingleReader) Get(ID string) (*account.Account, error) {
	ret := _m.Called(ID)

	var r0 *account.Account
	if rf, ok := ret.Get(0).(func(string) *account.Account); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*account.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
