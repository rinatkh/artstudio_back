package convert

import (
	"github.com/rinatkh/artstudio_back/internal/schedules/models/core"
	"github.com/rinatkh/artstudio_back/internal/schedules/models/dto"
	dtoSubject "github.com/rinatkh/artstudio_back/internal/subjects/models/dto"
	dtoUser "github.com/rinatkh/artstudio_back/internal/users/models/dto"
)

func Schedule2DTO(schedule *core.Schedule, subject dtoSubject.Subject, user dtoUser.User, cabinetName string) dto.Schedule {
	return dto.Schedule{
		Subject:     subject,
		WhenTime:    schedule.WhenTime,
		Student:     user,
		IsPaid:      schedule.IsPaid,
		Description: schedule.Description,
		IsConfirmed: schedule.IsConfirmed,
		Cabinet:     cabinetName,
	}
}
