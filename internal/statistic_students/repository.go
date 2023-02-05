package statisticStudent

import (
	"github.com/rinatkh/artstudio_back/internal/statistic_students/models/core"
)

type StatisticStudentRepository interface {
	GetStatisticStudentById(id string) (*core.StatisticStudent, error)
	DeleteStatisticStudent(id string) error
	CreateStatisticStudent(statisticStudent *core.StatisticStudent) (*core.StatisticStudent, error)
	UpdateStatisticStudent(statisticStudent *core.StatisticStudent) (*core.StatisticStudent, error)
}
