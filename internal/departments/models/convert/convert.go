package convert

import (
	"github.com/rinatkh/artstudio_back/internal/departments/models/core"
	"github.com/rinatkh/artstudio_back/internal/departments/models/dto"
)

func Department2DTO(department *core.Department) dto.Department {
	result := dto.Department{
		Id:    department.Id,
		Name:  department.Name,
		Image: department.Image,
	}
	if department.Description != nil {
		result.Description = *department.Description
	}
	return result
}
