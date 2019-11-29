// Code generated by mockery v1.0.0. DO NOT EDIT.

package activedirectory

import mock "github.com/stretchr/testify/mock"

// MockAPIInterface is an autogenerated mock type for the APIInterface type
type MockAPIInterface struct {
	mock.Mock
}

// connect provides a mock function with given fields:
func (_m *MockAPIInterface) connect() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// createComputer provides a mock function with given fields: cn, ou, description
func (_m *MockAPIInterface) createComputer(cn string, ou string, description string) error {
	ret := _m.Called(cn, ou, description)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(cn, ou, description)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// createOU provides a mock function with given fields: name, baseOU, description
func (_m *MockAPIInterface) createOU(name string, baseOU string, description string) error {
	ret := _m.Called(name, baseOU, description)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(name, baseOU, description)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// createObject provides a mock function with given fields: dn, class, attributes
func (_m *MockAPIInterface) createObject(dn string, class []string, attributes map[string][]string) error {
	ret := _m.Called(dn, class, attributes)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string, map[string][]string) error); ok {
		r0 = rf(dn, class, attributes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// deleteComputer provides a mock function with given fields: cn, ou
func (_m *MockAPIInterface) deleteComputer(cn string, ou string) error {
	ret := _m.Called(cn, ou)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(cn, ou)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// deleteOU provides a mock function with given fields: dn
func (_m *MockAPIInterface) deleteOU(dn string) error {
	ret := _m.Called(dn)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(dn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// deleteObject provides a mock function with given fields: dn
func (_m *MockAPIInterface) deleteObject(dn string) error {
	ret := _m.Called(dn)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(dn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// getComputer provides a mock function with given fields: name
func (_m *MockAPIInterface) getComputer(name string) (*Computer, error) {
	ret := _m.Called(name)

	var r0 *Computer
	if rf, ok := ret.Get(0).(func(string) *Computer); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Computer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// getDomainDN provides a mock function with given fields:
func (_m *MockAPIInterface) getDomainDN() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// getOU provides a mock function with given fields: name, baseOU
func (_m *MockAPIInterface) getOU(name string, baseOU string) (*OU, error) {
	ret := _m.Called(name, baseOU)

	var r0 *OU
	if rf, ok := ret.Get(0).(func(string, string) *OU); ok {
		r0 = rf(name, baseOU)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*OU)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, baseOU)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// getObject provides a mock function with given fields: dn, attributes
func (_m *MockAPIInterface) getObject(dn string, attributes []string) (*Object, error) {
	ret := _m.Called(dn, attributes)

	var r0 *Object
	if rf, ok := ret.Get(0).(func(string, []string) *Object); ok {
		r0 = rf(dn, attributes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Object)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(dn, attributes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// moveOU provides a mock function with given fields: cn, baseOU, newOU
func (_m *MockAPIInterface) moveOU(cn string, baseOU string, newOU string) error {
	ret := _m.Called(cn, baseOU, newOU)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(cn, baseOU, newOU)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// searchObject provides a mock function with given fields: filter, baseDN, attributes
func (_m *MockAPIInterface) searchObject(filter string, baseDN string, attributes []string) ([]*Object, error) {
	ret := _m.Called(filter, baseDN, attributes)

	var r0 []*Object
	if rf, ok := ret.Get(0).(func(string, string, []string) []*Object); ok {
		r0 = rf(filter, baseDN, attributes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Object)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, []string) error); ok {
		r1 = rf(filter, baseDN, attributes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// updateComputerDescription provides a mock function with given fields: cn, ou, description
func (_m *MockAPIInterface) updateComputerDescription(cn string, ou string, description string) error {
	ret := _m.Called(cn, ou, description)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(cn, ou, description)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// updateComputerOU provides a mock function with given fields: cn, ou, newOU
func (_m *MockAPIInterface) updateComputerOU(cn string, ou string, newOU string) error {
	ret := _m.Called(cn, ou, newOU)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(cn, ou, newOU)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// updateOUDescription provides a mock function with given fields: cn, baseOU, description
func (_m *MockAPIInterface) updateOUDescription(cn string, baseOU string, description string) error {
	ret := _m.Called(cn, baseOU, description)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(cn, baseOU, description)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// updateOUName provides a mock function with given fields: name, baseOU, newName
func (_m *MockAPIInterface) updateOUName(name string, baseOU string, newName string) error {
	ret := _m.Called(name, baseOU, newName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(name, baseOU, newName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// updateObject provides a mock function with given fields: dn, classes, added, changed, removed
func (_m *MockAPIInterface) updateObject(dn string, classes []string, added map[string][]string, changed map[string][]string, removed map[string][]string) error {
	ret := _m.Called(dn, classes, added, changed, removed)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string, map[string][]string, map[string][]string, map[string][]string) error); ok {
		r0 = rf(dn, classes, added, changed, removed)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}