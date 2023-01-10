package convert

import (
	"github.com/rinatkh/artstudio_back/internal/subjects/models/core"
	dtoSubject "github.com/rinatkh/artstudio_back/internal/subjects/models/dto"
	dtoUser "github.com/rinatkh/artstudio_back/internal/users/models/dto"
)

func Subject2DTO(subject *core.Subject, teacher *dtoUser.User) dtoSubject.Subject {
	result := dtoSubject.Subject{
		Id:      subject.Id,
		Name:    subject.Name,
		Image:   subject.Image,
		Teacher: *teacher,
	}
	if subject.Description != "" {
		result.Description = subject.Description
	}
	return result
}
