package convert

import (
	"github.com/rinatkh/artstudio_back/internal/cabinets/models/core"
	"github.com/rinatkh/artstudio_back/internal/cabinets/models/dto"
)

func ConvertCabinet2DTO(param *core.Cabinet, paramTime *[]core.CabinetTime) dto.Cabinet {
	res := dto.Cabinet{
		CabinetId: param.Id,
		Name:      param.Name,
		Time:      nil,
	}
	if paramTime != nil {
		for _, i := range *paramTime {
			res.Time = append(res.Time, dto.Time{
				Id:         i.Id,
				StartTime:  i.StartTime,
				FinishTime: i.FinishTime,
			})
		}
	}
	return res
}

func ConvertCabinet2DTOWithoutTime(param *dto.Cabinet) dto.CabinetWithoutTime {
	return dto.CabinetWithoutTime{
		CabinetId: param.CabinetId,
		Name:      param.Name,
	}
}
