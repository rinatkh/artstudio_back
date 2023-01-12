package statisticTeacher

import (
	"github.com/rinatkh/artstudio_back/internal/statistic_teachers/models/core"
)

type StatisticTeacherRepository interface {
	GetStatisticTeacherById(id string) (*core.StatisticTeacher, error)
	DeleteStatisticTeacher(id string) error
	CreateStatisticTeacher(statisticTeacher *core.StatisticTeacher) (*core.StatisticTeacher, error)
	UpdateStatisticTeacher(statisticTeacher *core.StatisticTeacher) (*core.StatisticTeacher, error)
}
