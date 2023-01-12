package statisticTeacher

import (
	"github.com/rinatkh/artstudio_back/internal/statistic_teachers/models/dto"
)

type UseCase interface {
	GetStatisticTeacher(params *dto.GetStatisticTeacherRequest) (*dto.GetStatisticTeacherResponse, error)
	DeleteStatisticTeacher(params *dto.DeleteStatisticTeacherRequest) (*dto.DeleteStatisticTeacherResponse, error)
	UpdateStatisticTeacher(params *dto.UpdateStatisticTeacherRequest) (*dto.UpdateStatisticTeacherResponse, error)
	CreateStatisticTeacher(params *dto.CreateStatisticTeacherRequest) (*dto.CreateStatisticTeacherResponse, error)
}
