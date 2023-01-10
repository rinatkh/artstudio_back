package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/config"
	"github.com/rinatkh/artstudio_back/internal/auth"
	consts "github.com/rinatkh/artstudio_back/internal/constants"
	"github.com/rinatkh/artstudio_back/internal/subjects"
	coreSubject "github.com/rinatkh/artstudio_back/internal/subjects/models/core"
	"github.com/rinatkh/artstudio_back/internal/users"
	"github.com/rinatkh/artstudio_back/internal/users/models/convert"
	"github.com/rinatkh/artstudio_back/internal/users/models/core"
	"github.com/rinatkh/artstudio_back/internal/users/models/dto"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/sirupsen/logrus"
	"time"
)

type UserUseCase struct {
	cfg          *config.Config
	log          *logrus.Entry
	repoUser     users.UserRepository
	repoSubjects subjects.SubjectRepository
	repoAuth     auth.AuthRepository
}

func NewUserUC(cfg *config.Config, log *logrus.Entry, repoUser users.UserRepository, repoSubjects subjects.SubjectRepository, repoAuth auth.AuthRepository) users.UseCase {
	return &UserUseCase{
		cfg:          cfg,
		log:          log,
		repoUser:     repoUser,
		repoSubjects: repoSubjects,
		repoAuth:     repoAuth,
	}
}
func (u UserUseCase) getRole(role string) string {
	switch role {
	case "t":
		return consts.Teacher
	case "a":
		return consts.Admin
	}
	return consts.Student
}

func (u UserUseCase) getSex(sex string) string {
	switch sex {
	case "m":
		return consts.MAN
	case "w":
		return consts.WOMAN
	}
	return consts.NOTHING
}

func (u UserUseCase) CreateUser(params *dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	date, err := time.Parse("2006-01-02", params.BirthDate)
	if err != nil {
		return nil, constants.ErrConvertData
	}
	user := core.User{
		Firstname:  params.Firstname,
		Surname:    params.Surname,
		Middlename: params.Middlename,
		Sex:        u.getSex(params.Sex),
		BirthDate:  date,
		Role:       u.getRole(params.Role),
		Image:      params.Image,
	}
	result, err := u.repoUser.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return &dto.CreateUserResponse{User: convert.User2DTO(result)}, nil
}

func (u UserUseCase) UpdateUser(params *dto.UpdateUserRequest) (*dto.UpdateUserResponse, error) {
	author, err := u.repoUser.GetUserById(params.UserId)
	if err != nil {
		return nil, err
	}
	if author == nil {
		return nil, constants.ErrNoPrivileges
	}
	if author.Role != consts.Admin && params.UserId != params.Id {
		return nil, constants.ErrNoPrivileges
	}
	check, err := u.repoUser.GetUserById(params.Id)
	if err != nil {
		return nil, err
	}
	if check == nil {
		return nil, constants.ErrUserDBNotFound
	}
	date, err := time.Parse("2006-01-02", params.BirthDate)
	if err != nil {
		return nil, constants.ErrConvertData
	}
	user := core.User{
		Id:         params.Id,
		Firstname:  params.Firstname,
		Surname:    params.Surname,
		Middlename: params.Middlename,
		Sex:        u.getSex(params.Sex),
		BirthDate:  date,
		Role:       u.getRole(params.Role),
		Image:      params.Image,
	}
	result, err := u.repoUser.UpdateUser(&user)
	if err != nil {
		return nil, err
	}

	return &dto.UpdateUserResponse{User: convert.User2DTO(result)}, nil
}

func (u UserUseCase) DeleteUser(params *dto.DeleteUserRequest) (*dto.DeleteUserResponse, error) {
	author, err := u.repoUser.GetUserById(params.UserId)
	if err != nil {
		return nil, err
	}
	if author == nil {
		return nil, constants.ErrNoPrivileges
	}
	if author.Role != consts.Admin && params.UserId != params.Id {
		return nil, constants.ErrNoPrivileges
	}
	check, err := u.repoUser.GetUserById(params.Id)
	if err != nil {
		return nil, err
	}
	if check == nil {
		return nil, constants.ErrUserDBNotFound
	}
	var list *[]coreSubject.Subject
	if check.Role == consts.Teacher {
		_, length, err := u.repoSubjects.GetSubjects(1, 0, check.Id)
		if err != nil {
			return nil, constants.NewCodedError("conflict to delete subjects of teacher", fiber.StatusConflict)
		}
		list, _, err = u.repoSubjects.GetSubjects(length, 0, check.Id)
		if err != nil {
			return nil, constants.NewCodedError("conflict to delete subjects of teacher", fiber.StatusConflict)
		}
	}
	if list != nil {
		for _, i := range *list {
			err := u.repoSubjects.DeleteSubject(i.Id)
			if err != nil {
				return nil, err
			}
		}
	}
	err = u.repoAuth.DeleteUser(params.Id)
	if err != nil {
		return nil, err
	}
	err = u.repoUser.DeleteUser(params.Id)
	if err != nil {
		return nil, err
	}
	return &dto.DeleteUserResponse{}, nil
}

func (u UserUseCase) GetUser(params *dto.GetUserRequest) (*dto.GetUserResponse, error) {
	result, err := u.repoUser.GetUserById(params.Id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, constants.ErrUserDBNotFound
	}
	return &dto.GetUserResponse{User: convert.User2DTO(result)}, nil
}

func (u UserUseCase) GetUsers(params *dto.GetUsersRequest) (*dto.GetUsersResponse, error) {
	list, length, err := u.repoUser.GetUsers(params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}
	if list == nil {
		return &dto.GetUsersResponse{}, nil
	}
	var result []dto.User
	for _, i := range *list {
		result = append(result, convert.User2DTO(&i))
	}
	return &dto.GetUsersResponse{Users: result, Length: length}, nil
}
