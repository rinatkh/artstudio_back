package usecase

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/config"
	consts "github.com/rinatkh/artstudio_back/internal/constants"
	"github.com/rinatkh/artstudio_back/internal/statistic_students"
	"github.com/rinatkh/artstudio_back/internal/statistic_students/models/convert"
	"github.com/rinatkh/artstudio_back/internal/statistic_students/models/core"
	"github.com/rinatkh/artstudio_back/internal/statistic_students/models/dto"
	"github.com/rinatkh/artstudio_back/internal/users"
	dtoUser "github.com/rinatkh/artstudio_back/internal/users/models/dto"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/sirupsen/logrus"
)

type StatisticStudentUseCase struct {
	cfg                  *config.Config
	log                  *logrus.Entry
	repoStatisticStudent statisticStudent.StatisticStudentRepository
	userUC               users.UseCase
}

func NewStatisticStudentUC(cfg *config.Config, log *logrus.Entry, repoStatisticStudent statisticStudent.StatisticStudentRepository, userUC users.UseCase) statisticStudent.UseCase {
	return &StatisticStudentUseCase{
		cfg:                  cfg,
		log:                  log,
		repoStatisticStudent: repoStatisticStudent,
		userUC:               userUC,
	}
}

func (u StatisticStudentUseCase) getStudent(userId, studentId string) (*dtoUser.User, error) {
	user, err := u.userUC.GetUser(&dtoUser.GetUserRequest{Id: userId})
	if err != nil {
		return nil, err
	}
	if user.Role == consts.Student && studentId != userId {
		return nil, constants.ErrNoPrivileges
	}
	student, err := u.userUC.GetUser(&dtoUser.GetUserRequest{Id: studentId})
	if err != nil {
		return nil, constants.NewCodedError(fmt.Sprintf("User %s is not student", studentId), fiber.StatusConflict)
	}
	if student.Role != consts.Student {
		return nil, constants.NewCodedError(fmt.Sprintf("User %s is not student", studentId), fiber.StatusConflict)
	}
	return &student.User, nil
}

func (u StatisticStudentUseCase) GetStatisticStudent(params *dto.GetStatisticStudentRequest) (*dto.GetStatisticStudentResponse, error) {
	student, err := u.getStudent(params.UserId, params.StudentId)
	if err != nil {
		return nil, err
	}
	res, err := u.repoStatisticStudent.GetStatisticStudentById(student.Id)
	if err != nil {
		return nil, err
	}
	return &dto.GetStatisticStudentResponse{StatisticStudent: convert.ConvertStatisticStudent2DTO(res), Student: *student}, nil
}

func (u StatisticStudentUseCase) CreateStatisticStudent(params *dto.CreateStatisticStudentRequest) (*dto.CreateStatisticStudentResponse, error) {
	res, err := u.repoStatisticStudent.CreateStatisticStudent(&core.StatisticStudent{StudentId: params.StudentId})
	if err != nil {
		return nil, err
	}
	return &dto.CreateStatisticStudentResponse{StatisticStudent: convert.ConvertStatisticStudent2DTO(res)}, nil
}

func (u StatisticStudentUseCase) DeleteStatisticStudent(params *dto.DeleteStatisticStudentRequest) (*dto.DeleteStatisticStudentResponse, error) {
	err := u.repoStatisticStudent.DeleteStatisticStudent(params.StudentId)
	if err != nil {
		return nil, err
	}
	return &dto.DeleteStatisticStudentResponse{}, nil
}

func (u StatisticStudentUseCase) UpdateStatisticStudent(params *dto.UpdateStatisticStudentRequest) (*dto.UpdateStatisticStudentResponse, error) {
	res, err := u.repoStatisticStudent.UpdateStatisticStudent(&core.StatisticStudent{
		StudentId:          params.StudentId,
		BalanceLessons:     params.BalanceLessons,
		NeedPaymentLessons: params.NeedPaymentLessons,
		DoneLessons:        params.DoneLessons,
	})
	if err != nil {
		return nil, err
	}
	return &dto.UpdateStatisticStudentResponse{StatisticStudent: convert.ConvertStatisticStudent2DTO(res)}, nil
}
