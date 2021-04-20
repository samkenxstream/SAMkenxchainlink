// Code generated by mockery v2.6.0. DO NOT EDIT.

package mocks

import (
	context "context"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	models "github.com/smartcontractkit/chainlink/core/store/models"

	pipeline "github.com/smartcontractkit/chainlink/core/services/pipeline"

	time "time"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

// CreateSpec provides a mock function with given fields: ctx, db, taskDAG, maxTaskTimeout
func (_m *ORM) CreateSpec(ctx context.Context, db *gorm.DB, taskDAG pipeline.TaskDAG, maxTaskTimeout models.Interval) (int32, error) {
	ret := _m.Called(ctx, db, taskDAG, maxTaskTimeout)

	var r0 int32
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, pipeline.TaskDAG, models.Interval) int32); ok {
		r0 = rf(ctx, db, taskDAG, maxTaskTimeout)
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *gorm.DB, pipeline.TaskDAG, models.Interval) error); ok {
		r1 = rf(ctx, db, taskDAG, maxTaskTimeout)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DB provides a mock function with given fields:
func (_m *ORM) DB() *gorm.DB {
	ret := _m.Called()

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// DeleteRunsOlderThan provides a mock function with given fields: threshold
func (_m *ORM) DeleteRunsOlderThan(threshold time.Duration) error {
	ret := _m.Called(threshold)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration) error); ok {
		r0 = rf(threshold)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindBridge provides a mock function with given fields: name
func (_m *ORM) FindBridge(name models.TaskType) (models.BridgeType, error) {
	ret := _m.Called(name)

	var r0 models.BridgeType
	if rf, ok := ret.Get(0).(func(models.TaskType) models.BridgeType); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(models.BridgeType)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.TaskType) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindRun provides a mock function with given fields: id
func (_m *ORM) FindRun(id int64) (pipeline.Run, error) {
	ret := _m.Called(id)

	var r0 pipeline.Run
	if rf, ok := ret.Get(0).(func(int64) pipeline.Run); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(pipeline.Run)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertFinishedRun provides a mock function with given fields: ctx, run, trrs, saveSuccessfulTaskRuns
func (_m *ORM) InsertFinishedRun(ctx context.Context, run pipeline.Run, trrs []pipeline.TaskRunResult, saveSuccessfulTaskRuns bool) (int64, error) {
	ret := _m.Called(ctx, run, trrs, saveSuccessfulTaskRuns)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, pipeline.Run, []pipeline.TaskRunResult, bool) int64); ok {
		r0 = rf(ctx, run, trrs, saveSuccessfulTaskRuns)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pipeline.Run, []pipeline.TaskRunResult, bool) error); ok {
		r1 = rf(ctx, run, trrs, saveSuccessfulTaskRuns)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
