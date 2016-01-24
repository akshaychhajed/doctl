package mocks

import "github.com/bryanl/doit/do"
import "github.com/stretchr/testify/mock"

import "github.com/digitalocean/godo"

type DropletsService struct {
	mock.Mock
}

func (_m *DropletsService) List() (do.Droplets, error) {
	ret := _m.Called()

	var r0 do.Droplets
	if rf, ok := ret.Get(0).(func() do.Droplets); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(do.Droplets)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *DropletsService) Get(_a0 int) (*do.Droplet, error) {
	ret := _m.Called(_a0)

	var r0 *do.Droplet
	if rf, ok := ret.Get(0).(func(int) *do.Droplet); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.Droplet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *DropletsService) Create(_a0 *godo.DropletCreateRequest, _a1 bool) (*do.Droplet, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *do.Droplet
	if rf, ok := ret.Get(0).(func(*godo.DropletCreateRequest, bool) *do.Droplet); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*do.Droplet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*godo.DropletCreateRequest, bool) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *DropletsService) CreateMultiple(_a0 *godo.DropletMultiCreateRequest) (do.Droplets, error) {
	ret := _m.Called(_a0)

	var r0 do.Droplets
	if rf, ok := ret.Get(0).(func(*godo.DropletMultiCreateRequest) do.Droplets); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(do.Droplets)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*godo.DropletMultiCreateRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *DropletsService) Delete(_a0 int) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *DropletsService) Kernels(_a0 int) (do.Kernels, error) {
	ret := _m.Called(_a0)

	var r0 do.Kernels
	if rf, ok := ret.Get(0).(func(int) do.Kernels); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(do.Kernels)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *DropletsService) Snapshots(_a0 int) (do.Images, error) {
	ret := _m.Called(_a0)

	var r0 do.Images
	if rf, ok := ret.Get(0).(func(int) do.Images); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(do.Images)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *DropletsService) Backups(_a0 int) (do.Images, error) {
	ret := _m.Called(_a0)

	var r0 do.Images
	if rf, ok := ret.Get(0).(func(int) do.Images); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(do.Images)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *DropletsService) Actions(_a0 int) (do.Actions, error) {
	ret := _m.Called(_a0)

	var r0 do.Actions
	if rf, ok := ret.Get(0).(func(int) do.Actions); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(do.Actions)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *DropletsService) Neighbors(_a0 int) (do.Droplets, error) {
	ret := _m.Called(_a0)

	var r0 do.Droplets
	if rf, ok := ret.Get(0).(func(int) do.Droplets); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(do.Droplets)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}