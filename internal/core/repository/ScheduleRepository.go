package repository

import (
	"go-hospital-server/internal/core/entity/models"
	"go-hospital-server/internal/core/entity/response"
)

type ScheduleRepository interface {
	FindByDate(string) (response.Session, error)
	Create(models.Schedule) error
	Update(models.Schedule) error
	Delete(int) error
}
