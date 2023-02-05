package convert

import (
	"github.com/rinatkh/artstudio_back/internal/cabinet/models/core"
	"github.com/rinatkh/artstudio_back/internal/cabinet/models/dto"
)

func ConvertCabinet2DTO(param *core.Cabinet, paramTime *[]core.CabinetTime) dto.Cabinet {
	res := dto.Cabinet{
		CabinetId: param.Id,
		Name:      param.Name,
		Time:      nil,
	}
	for _, i := range *paramTime {
		res.Time = append(res.Time, dto.Time{
			Id:         i.Id,
			StartTime:  i.StartTime,
			FinishTime: i.FinishTime,
		})
	}
	return res
}
