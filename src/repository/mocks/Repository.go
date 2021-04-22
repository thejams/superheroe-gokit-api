// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	entity "superheroe-gokit-api/src/entity"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// AddSuperheroe provides a mock function with given fields: c
func (_m *Repository) AddSuperheroe(c *entity.Superheroe) *entity.Superheroe {
	ret := _m.Called(c)

	var r0 *entity.Superheroe
	if rf, ok := ret.Get(0).(func(*entity.Superheroe) *entity.Superheroe); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Superheroe)
		}
	}

	return r0
}

// ClearRepository provides a mock function with given fields:
func (_m *Repository) ClearRepository() {
	_m.Called()
}

// DeleteSuperheroe provides a mock function with given fields: id
func (_m *Repository) DeleteSuperheroe(id string) string {
	ret := _m.Called(id)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// EditSuperheroe provides a mock function with given fields: c
func (_m *Repository) EditSuperheroe(c *entity.Superheroe) *entity.Superheroe {
	ret := _m.Called(c)

	var r0 *entity.Superheroe
	if rf, ok := ret.Get(0).(func(*entity.Superheroe) *entity.Superheroe); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Superheroe)
		}
	}

	return r0
}

// GetSuperheroeById provides a mock function with given fields: id
func (_m *Repository) GetSuperheroeById(id string) *entity.Superheroe {
	ret := _m.Called(id)

	var r0 *entity.Superheroe
	if rf, ok := ret.Get(0).(func(string) *entity.Superheroe); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Superheroe)
		}
	}

	return r0
}

// GetSuperheroes provides a mock function with given fields:
func (_m *Repository) GetSuperheroes() []*entity.Superheroe {
	ret := _m.Called()

	var r0 []*entity.Superheroe
	if rf, ok := ret.Get(0).(func() []*entity.Superheroe); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Superheroe)
		}
	}

	return r0
}
