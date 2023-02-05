package statisticStudent

import (
	"github.com/rinatkh/artstudio_back/internal/statistic_students/models/dto"
)

type UseCase interface {
	GetStatisticStudent(params *dto.GetStatisticStudentRequest) (*dto.GetStatisticStudentResponse, error)
	DeleteStatisticStudent(params *dto.DeleteStatisticStudentRequest) (*dto.DeleteStatisticStudentResponse, error)
	UpdateStatisticStudent(params *dto.UpdateStatisticStudentRequest) (*dto.UpdateStatisticStudentResponse, error)
	CreateStatisticStudent(params *dto.CreateStatisticStudentRequest) (*dto.CreateStatisticStudentResponse, error)
}
