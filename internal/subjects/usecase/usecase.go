package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/config"
	consts "github.com/rinatkh/artstudio_back/internal/constants"
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
	userUC      users.UseCase
}

func NewSubjectUC(cfg *config.Config, log *logrus.Entry, repoSubject subjects.SubjectRepository, userUC users.UseCase) subjects.UseCase {
	return &SubjectUseCase{
		cfg:         cfg,
		log:         log,
		repoSubject: repoSubject,
		userUC:      userUC,
	}
}

func (u SubjectUseCase) CreateSubject(params *dtoSubject.CreateSubjectRequest) (*dtoSubject.CreateSubjectResponse, error) {
	user, err := u.userUC.GetUser(&dtoUser.GetUserRequest{Id: params.UserId})
	if err != nil {
		return nil, err
	}
	subject := core.Subject{
		Name:        params.Name,
		Description: params.Description,
		Image:       params.Image,
		CabinetId:   params.CabinetId,
	}
	var tutor dtoUser.User
	if params.TeacherId != "" {
		teacher, err := u.userUC.GetUser(&dtoUser.GetUserRequest{Id: params.TeacherId})
		if err != nil {
			return nil, err
		}
		if user.Role != consts.Admin || teacher.Role != consts.Teacher {
			return nil, constants.ErrNoPrivileges
		}
		subject.TeacherId = params.TeacherId
		tutor = teacher.User
	} else {
		if user.Role == consts.Admin {
			return nil, constants.NewCodedError("teacher_id is empty", fiber.StatusConflict)
		}
		if user.Role != consts.Teacher {
			return nil, constants.ErrNoPrivileges
		}
		subject.TeacherId = params.UserId
		tutor = user.User
	}
	result, err := u.repoSubject.CreateSubject(&subject)
	if err != nil {
		return nil, err
	}

	return &dtoSubject.CreateSubjectResponse{Subject: convert.Subject2DTO(result, &tutor)}, nil
}

func (u SubjectUseCase) UpdateSubject(params *dtoSubject.UpdateSubjectRequest) (*dtoSubject.UpdateSubjectResponse, error) {
	check, err := u.repoSubject.GetSubjectById(params.Id)
	if err != nil {
		return nil, err
	}
	if check == nil {
		return nil, constants.ErrSubjectDBNotFound
	}
	user, err := u.userUC.GetUser(&dtoUser.GetUserRequest{Id: params.UserId})
	if err != nil {
		return nil, err
	}
	subject := core.Subject{
		Id:          params.Id,
		Name:        params.Name,
		Description: params.Description,
		Image:       params.Image,
		CabinetId:   params.CabinetId,
	}
	var tutor dtoUser.User
	if params.TeacherId != "" {
		teacher, err := u.userUC.GetUser(&dtoUser.GetUserRequest{Id: params.TeacherId})
		if err != nil {
			return nil, err
		}
		if user.Role != consts.Admin || teacher.Role != consts.Teacher {
			return nil, constants.ErrNoPrivileges
		}
		subject.TeacherId = params.TeacherId
		tutor = teacher.User
	} else {
		if user.Role == consts.Admin {
			return nil, constants.NewCodedError("teacher_id is empty", fiber.StatusConflict)
		}
		if user.Role != consts.Teacher || user.Id != check.TeacherId {
			return nil, constants.ErrNoPrivileges
		}
		subject.TeacherId = params.UserId
		tutor = user.User
	}

	result, err := u.repoSubject.UpdateSubject(&subject)
	if err != nil {
		return nil, err
	}

	return &dtoSubject.UpdateSubjectResponse{Subject: convert.Subject2DTO(result, &tutor)}, nil
}

func (u SubjectUseCase) DeleteSubject(params *dtoSubject.DeleteSubjectRequest) (*dtoSubject.DeleteSubjectResponse, error) {
	check, err := u.repoSubject.GetSubjectById(params.Id)
	if err != nil {
		return nil, err
	}
	if check == nil {
		return nil, constants.ErrSubjectDBNotFound
	}
	teacher, err := u.userUC.GetUser(&dtoUser.GetUserRequest{Id: params.UserId})
	if err != nil {
		return nil, err
	}
	if teacher.Role != consts.Admin && teacher.Role != consts.Teacher {
		return nil, constants.ErrNoPrivileges
	}
	if teacher.Role == consts.Teacher && teacher.Id != check.TeacherId {
		return nil, constants.ErrNoPrivileges
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
	teacher, err := u.userUC.GetUser(&dtoUser.GetUserRequest{Id: result.TeacherId})
	if err != nil {
		return nil, err
	}
	return &dtoSubject.GetSubjectResponse{Subject: convert.Subject2DTO(result, &teacher.User)}, nil
}

func (u SubjectUseCase) GetSubjects(params *dtoSubject.GetSubjectsRequest) (*dtoSubject.GetSubjectsResponse, error) {
	var teacherId string
	if params.TeacherId != "" {
		teacher, err := u.userUC.GetUser(&dtoUser.GetUserRequest{Id: params.TeacherId})
		if err != nil {
			return nil, err
		}
		if teacher.Role == consts.Teacher {
			teacherId = params.TeacherId
		}
	}

	list, length, err := u.repoSubject.GetSubjects(params.Limit, params.Offset, teacherId)
	if err != nil {
		return nil, err
	}
	if list == nil {
		return &dtoSubject.GetSubjectsResponse{}, nil
	}
	var result []dtoSubject.Subject
	for _, i := range *list {
		teacher, err := u.userUC.GetUser(&dtoUser.GetUserRequest{Id: i.TeacherId})
		if err != nil {
			return nil, err
		}
		result = append(result, convert.Subject2DTO(&i, &teacher.User))
	}
	return &dtoSubject.GetSubjectsResponse{Subjects: result, Length: length}, nil
}
