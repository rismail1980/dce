// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import db "github.com/Optum/dce/pkg/db"
import mock "github.com/stretchr/testify/mock"

// DBer is an autogenerated mock type for the DBer type
type DBer struct {
	mock.Mock
}

// FindAccountsByStatus provides a mock function with given fields: status
func (_m *DBer) FindAccountsByStatus(status db.AccountStatus) ([]*db.Account, error) {
	ret := _m.Called(status)

	var r0 []*db.Account
	if rf, ok := ret.Get(0).(func(db.AccountStatus) []*db.Account); ok {
		r0 = rf(status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*db.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(db.AccountStatus) error); ok {
		r1 = rf(status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindLeasesByAccount provides a mock function with given fields: accountID
func (_m *DBer) FindLeasesByAccount(accountID string) ([]*db.Lease, error) {
	ret := _m.Called(accountID)

	var r0 []*db.Lease
	if rf, ok := ret.Get(0).(func(string) []*db.Lease); ok {
		r0 = rf(accountID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*db.Lease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(accountID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindLeasesByPrincipal provides a mock function with given fields: principalID
func (_m *DBer) FindLeasesByPrincipal(principalID string) ([]*db.Lease, error) {
	ret := _m.Called(principalID)

	var r0 []*db.Lease
	if rf, ok := ret.Get(0).(func(string) []*db.Lease); ok {
		r0 = rf(principalID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*db.Lease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(principalID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindLeasesByStatus provides a mock function with given fields: status
func (_m *DBer) FindLeasesByStatus(status db.LeaseStatus) ([]*db.Lease, error) {
	ret := _m.Called(status)

	var r0 []*db.Lease
	if rf, ok := ret.Get(0).(func(db.LeaseStatus) []*db.Lease); ok {
		r0 = rf(status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*db.Lease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(db.LeaseStatus) error); ok {
		r1 = rf(status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccount provides a mock function with given fields: accountID
func (_m *DBer) GetAccount(accountID string) (*db.Account, error) {
	ret := _m.Called(accountID)

	var r0 *db.Account
	if rf, ok := ret.Get(0).(func(string) *db.Account); ok {
		r0 = rf(accountID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(accountID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLease provides a mock function with given fields: accountID, principalID
func (_m *DBer) GetLease(accountID string, principalID string) (*db.Lease, error) {
	ret := _m.Called(accountID, principalID)

	var r0 *db.Lease
	if rf, ok := ret.Get(0).(func(string, string) *db.Lease); ok {
		r0 = rf(accountID, principalID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.Lease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(accountID, principalID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLeaseByID provides a mock function with given fields: leaseID
func (_m *DBer) GetLeaseByID(leaseID string) (*db.Lease, error) {
	ret := _m.Called(leaseID)

	var r0 *db.Lease
	if rf, ok := ret.Get(0).(func(string) *db.Lease); ok {
		r0 = rf(leaseID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.Lease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(leaseID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLeases provides a mock function with given fields: input
func (_m *DBer) GetLeases(input db.GetLeasesInput) (db.GetLeasesOutput, error) {
	ret := _m.Called(input)

	var r0 db.GetLeasesOutput
	if rf, ok := ret.Get(0).(func(db.GetLeasesInput) db.GetLeasesOutput); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(db.GetLeasesOutput)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(db.GetLeasesInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReadyAccount provides a mock function with given fields:
func (_m *DBer) GetReadyAccount() (*db.Account, error) {
	ret := _m.Called()

	var r0 *db.Account
	if rf, ok := ret.Get(0).(func() *db.Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrphanAccount provides a mock function with given fields: accountID
func (_m *DBer) OrphanAccount(accountID string) (*db.Account, error) {
	ret := _m.Called(accountID)

	var r0 *db.Account
	if rf, ok := ret.Get(0).(func(string) *db.Account); ok {
		r0 = rf(accountID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(accountID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutAccount provides a mock function with given fields: account
func (_m *DBer) PutAccount(account db.Account) error {
	ret := _m.Called(account)

	var r0 error
	if rf, ok := ret.Get(0).(func(db.Account) error); ok {
		r0 = rf(account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutLease provides a mock function with given fields: lease
func (_m *DBer) PutLease(lease db.Lease) (*db.Lease, error) {
	ret := _m.Called(lease)

	var r0 *db.Lease
	if rf, ok := ret.Get(0).(func(db.Lease) *db.Lease); ok {
		r0 = rf(lease)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.Lease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(db.Lease) error); ok {
		r1 = rf(lease)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransitionAccountStatus provides a mock function with given fields: accountID, prevStatus, nextStatus
func (_m *DBer) TransitionAccountStatus(accountID string, prevStatus db.AccountStatus, nextStatus db.AccountStatus) (*db.Account, error) {
	ret := _m.Called(accountID, prevStatus, nextStatus)

	var r0 *db.Account
	if rf, ok := ret.Get(0).(func(string, db.AccountStatus, db.AccountStatus) *db.Account); ok {
		r0 = rf(accountID, prevStatus, nextStatus)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, db.AccountStatus, db.AccountStatus) error); ok {
		r1 = rf(accountID, prevStatus, nextStatus)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransitionLeaseStatus provides a mock function with given fields: accountID, principalID, prevStatus, nextStatus, leaseStatusReason
func (_m *DBer) TransitionLeaseStatus(accountID string, principalID string, prevStatus db.LeaseStatus, nextStatus db.LeaseStatus, leaseStatusReason db.LeaseStatusReason) (*db.Lease, error) {
	ret := _m.Called(accountID, principalID, prevStatus, nextStatus, leaseStatusReason)

	var r0 *db.Lease
	if rf, ok := ret.Get(0).(func(string, string, db.LeaseStatus, db.LeaseStatus, db.LeaseStatusReason) *db.Lease); ok {
		r0 = rf(accountID, principalID, prevStatus, nextStatus, leaseStatusReason)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.Lease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, db.LeaseStatus, db.LeaseStatus, db.LeaseStatusReason) error); ok {
		r1 = rf(accountID, principalID, prevStatus, nextStatus, leaseStatusReason)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAccountPrincipalPolicyHash provides a mock function with given fields: accountID, prevHash, nextHash
func (_m *DBer) UpdateAccountPrincipalPolicyHash(accountID string, prevHash string, nextHash string) (*db.Account, error) {
	ret := _m.Called(accountID, prevHash, nextHash)

	var r0 *db.Account
	if rf, ok := ret.Get(0).(func(string, string, string) *db.Account); ok {
		r0 = rf(accountID, prevHash, nextHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(accountID, prevHash, nextHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertLease provides a mock function with given fields: lease
func (_m *DBer) UpsertLease(lease db.Lease) (*db.Lease, error) {
	ret := _m.Called(lease)

	var r0 *db.Lease
	if rf, ok := ret.Get(0).(func(db.Lease) *db.Lease); ok {
		r0 = rf(lease)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db.Lease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(db.Lease) error); ok {
		r1 = rf(lease)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}