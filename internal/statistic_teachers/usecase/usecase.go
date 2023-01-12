package usecase

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/config"
	consts "github.com/rinatkh/artstudio_back/internal/constants"
	statisticTeacher "github.com/rinatkh/artstudio_back/internal/statistic_teachers"
	"github.com/rinatkh/artstudio_back/internal/statistic_teachers/models/convert"
	"github.com/rinatkh/artstudio_back/internal/statistic_teachers/models/core"
	"github.com/rinatkh/artstudio_back/internal/statistic_teachers/models/dto"
	"github.com/rinatkh/artstudio_back/internal/users"
	dtoUser "github.com/rinatkh/artstudio_back/internal/users/models/dto"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/sirupsen/logrus"
)

type StatisticTeacherUseCase struct {
	cfg                  *config.Config
	log                  *logrus.Entry
	repoStatisticTeacher statisticTeacher.StatisticTeacherRepository
	userUC               users.UseCase
}

func NewStatisticTeacherUC(cfg *config.Config, log *logrus.Entry, repoStatisticTeacher statisticTeacher.StatisticTeacherRepository, userUC users.UseCase) statisticTeacher.UseCase {
	return &StatisticTeacherUseCase{
		cfg:                  cfg,
		log:                  log,
		repoStatisticTeacher: repoStatisticTeacher,
		userUC:               userUC,
	}
}

func (u StatisticTeacherUseCase) getTeacher(userId, teacherId string) (*dtoUser.User, error) {
	user, err := u.userUC.GetUser(&dtoUser.GetUserRequest{Id: userId})
	if err != nil {
		return nil, err
	}
	if user.Role != consts.Teacher && user.Role != consts.Admin {
		return nil, constants.ErrNoPrivileges
	}
	if user.Role == consts.Teacher && teacherId != userId {
		return nil, constants.ErrNoPrivileges
	}
	teacher, err := u.userUC.GetUser(&dtoUser.GetUserRequest{Id: teacherId})
	if err != nil {
		return nil, constants.NewCodedError(fmt.Sprintf("User %s is not teacher", teacherId), fiber.StatusConflict)
	}
	if teacher.Role != consts.Teacher {
		return nil, constants.NewCodedError(fmt.Sprintf("User %s is not teacher", teacherId), fiber.StatusConflict)
	}
	return &teacher.User, nil
}

func (u StatisticTeacherUseCase) GetStatisticTeacher(params *dto.GetStatisticTeacherRequest) (*dto.GetStatisticTeacherResponse, error) {
	teacher, err := u.getTeacher(params.UserId, params.TeacherId)
	if err != nil {
		return nil, err
	}
	res, err := u.repoStatisticTeacher.GetStatisticTeacherById(teacher.Id)
	if err != nil {
		return nil, err
	}
	return &dto.GetStatisticTeacherResponse{StatisticTeacher: convert.ConvertStatisticTeacher2DTO(res), Teacher: *teacher}, nil
}

func (u StatisticTeacherUseCase) CreateStatisticTeacher(params *dto.CreateStatisticTeacherRequest) (*dto.CreateStatisticTeacherResponse, error) {
	res, err := u.repoStatisticTeacher.CreateStatisticTeacher(&core.StatisticTeacher{TeacherId: params.TeacherId})
	if err != nil {
		return nil, err
	}
	return &dto.CreateStatisticTeacherResponse{StatisticTeacher: convert.ConvertStatisticTeacher2DTO(res)}, nil
}

func (u StatisticTeacherUseCase) DeleteStatisticTeacher(params *dto.DeleteStatisticTeacherRequest) (*dto.DeleteStatisticTeacherResponse, error) {
	err := u.repoStatisticTeacher.DeleteStatisticTeacher(params.TeacherId)
	if err != nil {
		return nil, err
	}
	return &dto.DeleteStatisticTeacherResponse{}, nil
}

func (u StatisticTeacherUseCase) UpdateStatisticTeacher(params *dto.UpdateStatisticTeacherRequest) (*dto.UpdateStatisticTeacherResponse, error) {
	res, err := u.repoStatisticTeacher.UpdateStatisticTeacher(&core.StatisticTeacher{
		TeacherId:     params.TeacherId,
		FutureLessons: params.FutureLessons,
		NeedDzLessons: params.NeedDzLessons,
		PastLessons:   params.PastLessons,
	})
	if err != nil {
		return nil, err
	}
	return &dto.UpdateStatisticTeacherResponse{StatisticTeacher: convert.ConvertStatisticTeacher2DTO(res)}, nil
}
