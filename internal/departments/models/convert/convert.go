package convert

import (
	"github.com/rinatkh/artstudio_back/internal/departments/models/core"
	dtoDepartment "github.com/rinatkh/artstudio_back/internal/departments/models/dto"
	dtoSubject "github.com/rinatkh/artstudio_back/internal/subjects/models/dto"
)

func Department2DTO(department *core.Department, subject *[]dtoSubject.Subject, length int64) dtoDepartment.Department {
	result := dtoDepartment.Department{
		Id:             department.Id,
		Name:           department.Name,
		Image:          department.Image,
		Subject:        *subject,
		LengthSubjects: length,
	}
	if department.Description != nil {
		result.Description = *department.Description
	}
	return result
}
