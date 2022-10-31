// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Notificationer is an autogenerated mock type for the Notificationer type
type Notificationer struct {
	mock.Mock
}

// PublishMessage provides a mock function with given fields: topicArn, message, isJSON
func (_m *Notificationer) PublishMessage(topicArn *string, message *string, isJSON bool) (*string, error) {
	ret := _m.Called(topicArn, message, isJSON)

	var r0 *string
	if rf, ok := ret.Get(0).(func(*string, *string, bool) *string); ok {
		r0 = rf(topicArn, message, isJSON)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*string, *string, bool) error); ok {
		r1 = rf(topicArn, message, isJSON)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}