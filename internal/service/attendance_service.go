package service

import (
	"context"

	"github.com/BrondoL/wedding-be/config"
	"github.com/BrondoL/wedding-be/internal/constant"
	"github.com/BrondoL/wedding-be/internal/model"
	r "github.com/BrondoL/wedding-be/internal/repository"
	"github.com/BrondoL/wedding-be/internal/util"
	"github.com/BrondoL/wedding-be/pkg/cache"
)

type AttendanceService interface {
	GetAttendances(ctx context.Context) ([]*model.Attendance, error)
	CreateAttendance(ctx context.Context, m *model.Attendance) (*model.Attendance, error)
}

type attendanceService struct {
	cfg                  config.Config
	cache                cache.Cache
	attendanceRepository r.AttendanceRepository
}

type ASConfig struct {
	Cfg                  config.Config
	Cache                cache.Cache
	AttendanceRepository r.AttendanceRepository
}

func NewAttendanceService(c *ASConfig) AttendanceService {
	return &attendanceService{
		cfg:                  c.Cfg,
		cache:                c.Cache,
		attendanceRepository: c.AttendanceRepository,
	}
}

func (s *attendanceService) GetAttendances(ctx context.Context) ([]*model.Attendance, error) {
	var attendances []*model.Attendance

	err := s.cache.Get(ctx, constant.CacheKeyAttendance, &attendances)
	if err != nil {
		return nil, util.NewErrorWrapper(util.CodeServerError, "failed to get attendances from cache", nil, err)
	}

	if len(attendances) == 0 {
		attendances, err = s.attendanceRepository.FindAll()
		if err != nil {
			return nil, util.NewErrorWrapper(util.CodeServerError, "failed to get attendances from DB", nil, err)
		}

		err = s.cache.Set(ctx, constant.CacheKeyAttendance, attendances, constant.CacheTTLOneDay)
		if err != nil {
			return nil, util.NewErrorWrapper(util.CodeServerError, "failed to set attendances to cache", nil, err)
		}
	}

	return attendances, nil
}

func (s *attendanceService) CreateAttendance(ctx context.Context, m *model.Attendance) (*model.Attendance, error) {
	attendance, err := s.attendanceRepository.Save(m)
	if err != nil {
		return nil, util.NewErrorWrapper(util.CodeServerError, "failed to create attendance", nil, err)
	}

	err = s.cache.Set(ctx, constant.CacheKeyAttendance, nil, constant.CacheTTLInvalidate)
	if err != nil {
		return nil, util.NewErrorWrapper(util.CodeServerError, "failed to delete attendances from cache", nil, err)
	}

	return attendance, nil
}
