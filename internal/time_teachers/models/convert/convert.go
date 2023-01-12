package convert

import (
	"github.com/rinatkh/artstudio_back/internal/time_teachers/models/core"
	"github.com/rinatkh/artstudio_back/internal/time_teachers/models/dto"
	dtoUser "github.com/rinatkh/artstudio_back/internal/users/models/dto"
)

func ConvertTimeTeacher2DTO(timeTeacher *core.TimeTeacher, user *dtoUser.User) dto.OneTimeTeacher {
	return dto.OneTimeTeacher{
		TeacherId: *user,
		Time: dto.TimeTeacher{
			Id:         timeTeacher.Id,
			StartTime:  timeTeacher.StartTime,
			FinishTime: timeTeacher.FinishTime,
		},
	}
}

func ConvertAllTimeTeacher2DTO(timeTeacher *[]core.TimeTeacher, user *dtoUser.User) dto.AllTimeTeacher {
	result := dto.AllTimeTeacher{
		Teacher: *user,
	}
	for _, i := range *timeTeacher {
		result.Time = append(result.Time, dto.TimeTeacher{
			Id:         i.Id,
			StartTime:  i.StartTime,
			FinishTime: i.FinishTime,
		})
	}
	return result
}
