package usecase

import (
	"github.com/rinatkh/artstudio_back/config"
	consts "github.com/rinatkh/artstudio_back/internal/constants"
	statisticStudent "github.com/rinatkh/artstudio_back/internal/statistic_students"
	coreStatisticStudent "github.com/rinatkh/artstudio_back/internal/statistic_students/models/core"
	statisticTeacher "github.com/rinatkh/artstudio_back/internal/statistic_teachers"
	coreStatisticTeacher "github.com/rinatkh/artstudio_back/internal/statistic_teachers/models/core"
	"github.com/rinatkh/artstudio_back/internal/users"
	"github.com/rinatkh/artstudio_back/internal/users/models/convert"
	"github.com/rinatkh/artstudio_back/internal/users/models/core"
	"github.com/rinatkh/artstudio_back/internal/users/models/dto"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/sirupsen/logrus"
	"time"
)

type UserUseCase struct {
	cfg                  *config.Config
	log                  *logrus.Entry
	repoUser             users.UserRepository
	repoStatisticStudent statisticStudent.StatisticStudentRepository
	repoStatisticTeacher statisticTeacher.StatisticTeacherRepository
}

func NewUserUC(cfg *config.Config, log *logrus.Entry, repoUser users.UserRepository, repoStatisticStudent statisticStudent.StatisticStudentRepository, repoStatisticTeacher statisticTeacher.StatisticTeacherRepository) users.UseCase {
	return &UserUseCase{
		cfg:                  cfg,
		log:                  log,
		repoUser:             repoUser,
		repoStatisticTeacher: repoStatisticTeacher,
		repoStatisticStudent: repoStatisticStudent,
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

	user := core.User{
		Firstname:  params.Firstname,
		Surname:    params.Surname,
		Middlename: params.Middlename,
		Sex:        u.getSex(params.Sex),
		BirthDate:  params.BirthDate,
		Role:       u.getRole(params.Role),
		Image:      params.Image,
		CreateAt:   time.Now().Unix(),
	}
	result, err := u.repoUser.CreateUser(&user)
	if err != nil {
		return nil, err
	}
	if result.Role == consts.Student {
		_, err = u.repoStatisticStudent.CreateStatisticStudent(&coreStatisticStudent.StatisticStudent{StudentId: result.Id})
		if err != nil {
			return nil, err
		}
	} else if result.Role == consts.Teacher {
		_, err = u.repoStatisticTeacher.CreateStatisticTeacher(&coreStatisticTeacher.StatisticTeacher{TeacherId: result.Id})
		if err != nil {
			return nil, err
		}
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
	user := core.User{
		Id:         params.Id,
		Firstname:  params.Firstname,
		Surname:    params.Surname,
		Middlename: params.Middlename,
		Sex:        u.getSex(params.Sex),
		BirthDate:  params.BirthDate,
		Role:       u.getRole(params.Role),
		Image:      params.Image,
	}
	result, err := u.repoUser.UpdateUser(&user)
	if err != nil {
		return nil, err
	}
	if check.Role != result.Role {
		if check.Role == consts.Student {
			err = u.repoStatisticStudent.DeleteStatisticStudent(check.Id)
			if err != nil {
				return nil, err
			}
		} else if check.Role == consts.Teacher {
			err = u.repoStatisticTeacher.DeleteStatisticTeacher(check.Id)
			if err != nil {
				return nil, err
			}
		}
		if result.Role == consts.Student {
			_, err = u.repoStatisticStudent.CreateStatisticStudent(&coreStatisticStudent.StatisticStudent{StudentId: result.Id})
			if err != nil {
				return nil, err
			}
		} else if result.Role == consts.Teacher {
			_, err = u.repoStatisticTeacher.CreateStatisticTeacher(&coreStatisticTeacher.StatisticTeacher{TeacherId: result.Id})
			if err != nil {
				return nil, err
			}
		}
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
