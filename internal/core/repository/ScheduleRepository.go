package repository

import (
	"go-hospital-server/internal/core/entity/models"
)

type ScheduleRepository interface {
	Create(models.Schedule) error
	Update(models.Schedule) error
	Delete(int) error
}
