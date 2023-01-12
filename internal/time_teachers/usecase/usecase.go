package usecase

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/config"
	consts "github.com/rinatkh/artstudio_back/internal/constants"
	timeTeacher "github.com/rinatkh/artstudio_back/internal/time_teachers"
	"github.com/rinatkh/artstudio_back/internal/time_teachers/models/convert"
	"github.com/rinatkh/artstudio_back/internal/time_teachers/models/core"
	"github.com/rinatkh/artstudio_back/internal/time_teachers/models/dto"
	"github.com/rinatkh/artstudio_back/internal/users"
	dtoUser "github.com/rinatkh/artstudio_back/internal/users/models/dto"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/sirupsen/logrus"
	"time"
)

type TimeTeacherUseCase struct {
	cfg             *config.Config
	log             *logrus.Entry
	repoTimeTeacher timeTeacher.TimeTeacherRepository
	userUC          users.UseCase
}

func NewSubjectUC(cfg *config.Config, log *logrus.Entry, repoTimeTeacher timeTeacher.TimeTeacherRepository, userUC users.UseCase) timeTeacher.UseCase {
	return &TimeTeacherUseCase{
		cfg:             cfg,
		log:             log,
		repoTimeTeacher: repoTimeTeacher,
		userUC:          userUC,
	}
}

func (u TimeTeacherUseCase) getTeacher(userId, teacherId string) (*dtoUser.User, error) {
	user, err := u.userUC.GetUser(&dtoUser.GetUserRequest{Id: userId})
	if err != nil {
		return nil, err
	}
	if user.Role != consts.Teacher && user.Role != consts.Admin {
		return nil, constants.ErrNoPrivileges
	}
	if user.Role == consts.Admin && teacherId == "" {
		return nil, constants.ErrNoPrivileges
	}
	if teacherId != "" {
		teacher, err := u.userUC.GetUser(&dtoUser.GetUserRequest{Id: teacherId})
		if err != nil {
			return nil, constants.NewCodedError(fmt.Sprintf("User %s is not teacher", teacherId), fiber.StatusConflict)
		}
		if teacher.Role != consts.Teacher {
			return nil, constants.NewCodedError(fmt.Sprintf("User %s is not teacher", teacherId), fiber.StatusConflict)
		}
		return &teacher.User, nil
	}
	return &user.User, nil
}

func (u TimeTeacherUseCase) DeleteTimeTeacher(params *dto.DeleteTimeTeacherRequest) (*dto.DeleteTimeTeacherResponse, error) {
	teacher, err := u.getTeacher(params.UserId, params.TeacherId)
	if err != nil {
		return nil, err
	}
	timeT, err := u.repoTimeTeacher.GetTimeTeacherById(params.Id)
	if err != nil {
		return nil, err
	}
	if timeT == nil {
		return nil, constants.ErrTimeTeachersDBNotFound
	}
	if teacher.Id != params.TeacherId {
		return nil, constants.ErrNoPrivileges
	}
	err = u.repoTimeTeacher.DeleteTimeTeacher(params.Id)
	if err != nil {
		return nil, err
	}
	return &dto.DeleteTimeTeacherResponse{}, nil
}

func (u TimeTeacherUseCase) UpdateTimeTeacher(params *dto.UpdateTimeTeacherRequest) (*dto.UpdateTimeTeacherResponse, error) {
	teacher, err := u.getTeacher(params.UserId, params.TeacherId)
	if err != nil {
		return nil, err
	}
	timeT, err := u.repoTimeTeacher.GetTimeTeacherById(params.Id)
	if err != nil {
		return nil, err
	}
	if timeT == nil {
		return nil, constants.ErrTimeTeachersDBNotFound
	}
	if teacher.Id != params.TeacherId {
		return nil, constants.ErrNoPrivileges
	}
	startTime, err := time.Parse(time.RFC3339, params.Time.StartTime)
	if err != nil {
		return nil, constants.ErrConvertData
	}
	finishTime, err := time.Parse(time.RFC3339, params.Time.FinishTime)
	if err != nil {
		return nil, constants.ErrConvertData
	}

	updateT, err := u.repoTimeTeacher.UpdateTimeTeacher(&core.TimeTeacher{
		Id:         params.Id,
		TeacherId:  teacher.Id,
		StartTime:  startTime,
		FinishTime: finishTime,
	})
	if err != nil {
		return nil, err
	}
	return &dto.UpdateTimeTeacherResponse{OneTimeTeacher: convert.ConvertTimeTeacher2DTO(updateT, teacher)}, nil
}

func (u TimeTeacherUseCase) GetTimeTeacher(params *dto.GetTimeTeacherRequest) (*dto.GetTimeTeacherResponse, error) {
	timeT, err := u.repoTimeTeacher.GetTimeTeacherById(params.Id)
	if err != nil {
		return nil, err
	}
	teacher, err := u.userUC.GetUser(&dtoUser.GetUserRequest{Id: timeT.TeacherId})
	if err != nil {
		return nil, err
	}
	return &dto.GetTimeTeacherResponse{OneTimeTeacher: convert.ConvertTimeTeacher2DTO(timeT, &teacher.User)}, nil
}

func (u TimeTeacherUseCase) GetTimeTeachers(params *dto.GetTimeTeachersRequest) (*dto.GetTimeTeachersResponse, error) {
	startTime, err := time.Parse(time.RFC3339, params.StartTime)
	if err != nil {
		return nil, constants.ErrConvertData
	}
	finishTime, err := time.Parse(time.RFC3339, params.FinishTime)
	if err != nil {
		return nil, constants.ErrConvertData
	}
	result, err := u.repoTimeTeacher.GetTimeTeachers(params.TeacherId, startTime, finishTime)
	if err != nil {
		return nil, err
	}
	teacher, err := u.userUC.GetUser(&dtoUser.GetUserRequest{Id: params.TeacherId})
	if err != nil {
		return nil, err
	}
	return &dto.GetTimeTeachersResponse{AllTimeTeacher: convert.ConvertAllTimeTeacher2DTO(result, &teacher.User)}, nil
}

func (u TimeTeacherUseCase) CreateTimeTeacher(params *dto.CreateTimeTeacherRequest) (*dto.CreateTimeTeacherResponse, error) {
	teacher, err := u.getTeacher(params.UserId, params.TeacherId)
	if err != nil {
		return nil, err
	}
	var temp []core.TimeTeacher
	for _, i := range params.Time {
		startTime, err := time.Parse(time.RFC3339, i.StartTime)
		if err != nil {
			return nil, constants.ErrConvertData
		}
		finishTime, err := time.Parse(time.RFC3339, i.FinishTime)
		if err != nil {
			return nil, constants.ErrConvertData
		}
		temp = append(temp, core.TimeTeacher{
			TeacherId:  teacher.Id,
			StartTime:  startTime,
			FinishTime: finishTime,
		})
	}
	var result []core.TimeTeacher
	for _, i := range temp {
		timeT, err := u.repoTimeTeacher.CreateTimeTeacher(&i)
		if err != nil {
			return nil, err
		}
		result = append(result, *timeT)
	}
	return &dto.CreateTimeTeacherResponse{AllTimeTeacher: convert.ConvertAllTimeTeacher2DTO(&result, teacher)}, nil
}
