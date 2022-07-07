package repository

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/response"
)

type DoctorScheduleRepository interface {
	FindAll() ([]response.ScheduleList, error)
	Create(models.DoctorSchedule) error
	Update(models.DoctorSchedule) error
	Delete(int) error
}
