// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package crypt

import mock "github.com/stretchr/testify/mock"

// MockEncrypter is an autogenerated mock type for the Encrypter type
type MockEncrypter struct {
	mock.Mock
}

// CheckPassword provides a mock function with given fields: existing, new
func (_m *MockEncrypter) CheckPassword(existing []byte, new string) error {
	ret := _m.Called(existing, new)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, string) error); ok {
		r0 = rf(existing, new)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EncryptPassword provides a mock function with given fields: s
func (_m *MockEncrypter) EncryptPassword(s string) ([]byte, error) {
	ret := _m.Called(s)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(s)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
