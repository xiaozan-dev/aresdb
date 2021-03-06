// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// EnumMutator is an autogenerated mock type for the EnumMutator type
type EnumMutator struct {
	mock.Mock
}

// ExtendEnumCases provides a mock function with given fields: namespace, table, column, enumCases
func (_m *EnumMutator) ExtendEnumCases(namespace string, table string, column string, enumCases []string) ([]int, error) {
	ret := _m.Called(namespace, table, column, enumCases)

	var r0 []int
	if rf, ok := ret.Get(0).(func(string, string, string, []string) []int); ok {
		r0 = rf(namespace, table, column, enumCases)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, []string) error); ok {
		r1 = rf(namespace, table, column, enumCases)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEnumCases provides a mock function with given fields: namespace, table, column
func (_m *EnumMutator) GetEnumCases(namespace string, table string, column string) ([]string, error) {
	ret := _m.Called(namespace, table, column)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, string, string) []string); ok {
		r0 = rf(namespace, table, column)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(namespace, table, column)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
