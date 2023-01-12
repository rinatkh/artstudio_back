package convert

import (
	"github.com/rinatkh/artstudio_back/internal/statistic_teachers/models/core"
	"github.com/rinatkh/artstudio_back/internal/statistic_teachers/models/dto"
)

func ConvertStatisticTeacher2DTO(param *core.StatisticTeacher) dto.StatisticTeacher {
	return dto.StatisticTeacher{
		FutureLessons: param.FutureLessons,
		NeedDzLessons: param.NeedDzLessons,
		PastLessons:   param.PastLessons,
	}
}
