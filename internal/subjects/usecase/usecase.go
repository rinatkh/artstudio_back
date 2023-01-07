package usecase

import (
	"github.com/rinatkh/artstudio_back/config"
	"github.com/rinatkh/artstudio_back/internal/subjects"
	"github.com/rinatkh/artstudio_back/internal/subjects/models/convert"
	"github.com/rinatkh/artstudio_back/internal/subjects/models/core"
	dtoSubject "github.com/rinatkh/artstudio_back/internal/subjects/models/dto"
	"github.com/rinatkh/artstudio_back/internal/users"
	dtoUser "github.com/rinatkh/artstudio_back/internal/users/models/dto"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/sirupsen/logrus"
)

type SubjectUseCase struct {
	cfg         *config.Config
	log         *logrus.Entry
	repoSubject subjects.SubjectRepository
	userUS      users.UseCase
}

func NewSubjectUC(cfg *config.Config, log *logrus.Entry, repoSubject subjects.SubjectRepository, userUS users.UseCase) subjects.UseCase {
	return &SubjectUseCase{
		cfg:         cfg,
		log:         log,
		repoSubject: repoSubject,
		userUS:      userUS,
	}
}

func (u SubjectUseCase) CreateSubject(params *dtoSubject.CreateSubjectRequest) (*dtoSubject.CreateSubjectResponse, error) {
	teacher, err := u.userUS.GetUser(&dtoUser.GetUserRequest{Id: params.TeacherId})
	if err != nil {
		return nil, err
	}
	subject := core.Subject{
		Name:        params.Name,
		Description: &params.Description,
		Image:       params.Image,
		TeacherId:   teacher.Id,
	}
	result, err := u.repoSubject.CreateSubject(&subject)
	if err != nil {
		return nil, err
	}

	return &dtoSubject.CreateSubjectResponse{Subject: convert.Subject2DTO(result, &teacher.User)}, nil
}

func (u SubjectUseCase) UpdateSubject(params *dtoSubject.UpdateSubjectRequest) (*dtoSubject.UpdateSubjectResponse, error) {
	check, err := u.repoSubject.GetSubjectById(params.Id)
	if err != nil {
		return nil, err
	}
	if check == nil {
		return nil, constants.ErrSubjectDBNotFound
	}
	teacher, err := u.userUS.GetUser(&dtoUser.GetUserRequest{Id: params.TeacherId})
	if err != nil {
		return nil, err
	}
	subject := core.Subject{
		Name:        params.Name,
		Description: &params.Description,
		Image:       params.Image,
		TeacherId:   teacher.Id,
	}
	result, err := u.repoSubject.UpdateSubject(&subject)
	if err != nil {
		return nil, err
	}

	return &dtoSubject.UpdateSubjectResponse{Subject: convert.Subject2DTO(result, &teacher.User)}, nil
}

func (u SubjectUseCase) DeleteSubject(params *dtoSubject.DeleteSubjectRequest) (*dtoSubject.DeleteSubjectResponse, error) {
	check, err := u.repoSubject.GetSubjectById(params.Id)
	if err != nil {
		return nil, err
	}
	if check == nil {
		return nil, constants.ErrSubjectDBNotFound
	}
	err = u.repoSubject.DeleteSubject(params.Id)
	if err != nil {
		return nil, err
	}
	return &dtoSubject.DeleteSubjectResponse{}, nil
}

func (u SubjectUseCase) GetSubject(params *dtoSubject.GetSubjectRequest) (*dtoSubject.GetSubjectResponse, error) {
	result, err := u.repoSubject.GetSubjectById(params.Id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, constants.ErrSubjectDBNotFound
	}
	teacher, err := u.userUS.GetUser(&dtoUser.GetUserRequest{Id: result.TeacherId})
	if err != nil {
		return nil, err
	}
	return &dtoSubject.GetSubjectResponse{Subject: convert.Subject2DTO(result, &teacher.User)}, nil
}

func (u SubjectUseCase) GetSubjects(params *dtoSubject.GetSubjectsRequest) (*dtoSubject.GetSubjectsResponse, error) {
	list, length, err := u.repoSubject.GetSubjects(params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}
	if list == nil {
		return &dtoSubject.GetSubjectsResponse{}, nil
	}
	var result []dtoSubject.Subject
	for _, i := range *list {
		teacher, err := u.userUS.GetUser(&dtoUser.GetUserRequest{Id: i.TeacherId})
		if err != nil {
			return nil, err
		}
		result = append(result, convert.Subject2DTO(&i, &teacher.User))
	}
	return &dtoSubject.GetSubjectsResponse{Subjects: result, Length: length}, nil
}
