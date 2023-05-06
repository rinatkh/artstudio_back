package convert

import (
	"github.com/rinatkh/artstudio_back/internal/users/models/core"
	"github.com/rinatkh/artstudio_back/internal/users/models/dto"
)

func User2DTO(user *core.User) dto.User {
	result := dto.User{
		Id:        user.Id,
		Firstname: user.Firstname,
		Surname:   user.Surname,
		BirthDate: user.BirthDate,
		//Age:       utils.RoundTime(time.Now().Sub(time.Unix(user.BirthDate, 0)).Seconds() / 31207680),
		Sex:   user.Sex,
		Image: user.Image,
		Role:  user.Role,
	}
	if user.Middlename != "" {
		result.Middlename = user.Middlename
	}
	return result
}
