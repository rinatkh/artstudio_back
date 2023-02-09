package convert

import (
	convertcabinet "github.com/rinatkh/artstudio_back/internal/cabinets/models/convert"
	dtoCabinet "github.com/rinatkh/artstudio_back/internal/cabinets/models/dto"
	"github.com/rinatkh/artstudio_back/internal/subjects/models/core"
	dtoSubject "github.com/rinatkh/artstudio_back/internal/subjects/models/dto"
	dtoUser "github.com/rinatkh/artstudio_back/internal/users/models/dto"
)

func Subject2DTO(subject *core.Subject, teacher *dtoUser.User, cabinet *dtoCabinet.Cabinet) dtoSubject.Subject {
	result := dtoSubject.Subject{
		Id:      subject.Id,
		Name:    subject.Name,
		Image:   subject.Image,
		Teacher: *teacher,
		Cabinet: convertcabinet.ConvertCabinet2DTOWithoutTime(cabinet),
	}
	if subject.Description != "" {
		result.Description = subject.Description
	}
	return result
}
