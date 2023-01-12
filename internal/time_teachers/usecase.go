package timeTeacher

import (
	"github.com/rinatkh/artstudio_back/internal/time_teachers/models/dto"
)

type UseCase interface {
	DeleteTimeTeacher(params *dto.DeleteTimeTeacherRequest) (*dto.DeleteTimeTeacherResponse, error)
	UpdateTimeTeacher(params *dto.UpdateTimeTeacherRequest) (*dto.UpdateTimeTeacherResponse, error)
	GetTimeTeacher(params *dto.GetTimeTeacherRequest) (*dto.GetTimeTeacherResponse, error)
	GetTimeTeachers(params *dto.GetTimeTeachersRequest) (*dto.GetTimeTeachersResponse, error)
	CreateTimeTeacher(params *dto.CreateTimeTeacherRequest) (*dto.CreateTimeTeacherResponse, error)
}
