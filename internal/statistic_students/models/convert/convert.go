package convert

import (
	"github.com/rinatkh/artstudio_back/internal/statistic_students/models/core"
	"github.com/rinatkh/artstudio_back/internal/statistic_students/models/dto"
)

func ConvertStatisticStudent2DTO(param *core.StatisticStudent) dto.StatisticStudent {
	return dto.StatisticStudent{
		BalanceLessons:     param.BalanceLessons,
		DoneLessons:        param.DoneLessons,
		NeedPaymentLessons: param.NeedPaymentLessons,
	}
}
