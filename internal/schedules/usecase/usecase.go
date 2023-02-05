package usecase

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/config"
	"github.com/rinatkh/artstudio_back/internal/cabinet"
	coreCabinet "github.com/rinatkh/artstudio_back/internal/cabinet/models/core"
	consts "github.com/rinatkh/artstudio_back/internal/constants"
	"github.com/rinatkh/artstudio_back/internal/schedules"
	"github.com/rinatkh/artstudio_back/internal/schedules/models/convert"
	"github.com/rinatkh/artstudio_back/internal/schedules/models/core"
	"github.com/rinatkh/artstudio_back/internal/schedules/models/dto"
	"github.com/rinatkh/artstudio_back/internal/subjects"
	dtoSubject "github.com/rinatkh/artstudio_back/internal/subjects/models/dto"
	timeTeacher "github.com/rinatkh/artstudio_back/internal/time_teachers"
	coreTimeTeacher "github.com/rinatkh/artstudio_back/internal/time_teachers/models/core"
	"github.com/rinatkh/artstudio_back/internal/users"
	dtoUser "github.com/rinatkh/artstudio_back/internal/users/models/dto"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/rinatkh/artstudio_back/pkg/utils"
	"github.com/sirupsen/logrus"
	"time"
)

type ScheduleUseCase struct {
	cfg             *config.Config
	log             *logrus.Entry
	repoSchedule    schedules.ScheduleRepository
	UserUC          users.UseCase
	SubjectUC       subjects.UseCase
	repoTimeTeacher timeTeacher.TimeTeacherRepository
	repoCabinet     cabinet.CabinetRepository
}

func NewScheduleUC(cfg *config.Config, log *logrus.Entry, repoSchedule schedules.ScheduleRepository,
	UserUC users.UseCase, SubjectUC subjects.UseCase,
	repoTimeTeacher timeTeacher.TimeTeacherRepository, repoCabinet cabinet.CabinetRepository) schedules.UseCase {
	return &ScheduleUseCase{
		cfg:             cfg,
		log:             log,
		repoSchedule:    repoSchedule,
		UserUC:          UserUC,
		SubjectUC:       SubjectUC,
		repoTimeTeacher: repoTimeTeacher,
		repoCabinet:     repoCabinet,
	}
}

func (u ScheduleUseCase) checkUser(userId, teacherId string) (*dtoUser.User, error) {
	user, err := u.UserUC.GetUser(&dtoUser.GetUserRequest{Id: userId})
	if err != nil {
		return nil, err
	}
	if user.Role != consts.Teacher && user.Role != consts.Admin {
		return nil, constants.ErrNoPrivileges
	}
	if user.Role == consts.Teacher && user.Id != teacherId {
		return nil, constants.ErrNoPrivileges
	}
	teacher, err := u.UserUC.GetUser(&dtoUser.GetUserRequest{Id: teacherId})
	if err != nil {
		return nil, err
	}
	return &teacher.User, nil
}

func (u ScheduleUseCase) CreateSchedule(params *dto.CreateScheduleRequest) (*dto.CreateScheduleResponse, error) {
	if err := utils.IsTime15MinDuration(params.WhenTime); err != nil {
		return nil, err
	}
	subject, err := u.SubjectUC.GetSubject(&dtoSubject.GetSubjectRequest{Id: params.SubjectId})
	if err != nil {
		return nil, err
	}
	user, err := u.UserUC.GetUser(&dtoUser.GetUserRequest{Id: params.UserId})
	if err != nil {
		return nil, err
	}
	student, err := u.UserUC.GetUser(&dtoUser.GetUserRequest{Id: params.StudentId})
	if err != nil {
		return nil, constants.NewCodedError(fmt.Sprintf("User %s is not student", params.StudentId), fiber.StatusConflict)
	}
	if student.Role != consts.Student {
		return nil, constants.NewCodedError(fmt.Sprintf("User %s is not student", params.StudentId), fiber.StatusConflict)
	}
	var cabinetTime *[]coreCabinet.CabinetTime
	if subject.CabinetId == 0 {
		cabinetTime, err = u.repoCabinet.GetAllCabinetTimes(params.WhenTime, time.Unix(params.WhenTime, 0).Add(consts.LessonTime+consts.Duration).Unix())
		if err != nil {
			return nil, err
		}
	} else {
		cabinetTime, err = u.repoCabinet.GetCabinetTimes(subject.CabinetId, params.WhenTime, time.Unix(params.WhenTime, 0).Add(consts.LessonTime+consts.Duration).Unix())
		if err != nil {
			return nil, err
		}
	}
	teacherTime, err := u.repoTimeTeacher.GetTimeTeachers(subject.Teacher.Id, params.WhenTime, time.Unix(params.WhenTime, 0).Add(consts.LessonTime+consts.Duration).Unix())
	if err != nil {
		return nil, err
	}
	// check
	_, err = u.repoSchedule.GetSubjectSchedules(subject.Id, params.WhenTime, time.Unix(params.WhenTime, 0).Add(consts.LessonTime+consts.Duration).Unix())
	if err != constants.ErrScheduleDBNotFound {
		return nil, constants.NewCodedError("Cabinet is not free at this time", fiber.StatusConflict)
	}
	cab, err := u.repoCabinet.GetCabinetById((*cabinetTime)[0].CabinetId)
	if err != nil {
		return nil, err
	}
	//TODO payment
	var schedule core.Schedule
	if user.Role == consts.Student {
		if user.Id != params.StudentId {
			return nil, constants.ErrNoPrivileges
		}
		schedule = core.Schedule{
			SubjectId:   subject.Id,
			StudentId:   user.Id,
			IsPaid:      false,
			IsConfirmed: false,
			IsFinished:  false,
			CabinetId:   cab.Id,
			WhenTime:    params.WhenTime,
		}
	} else if user.Role == consts.Teacher {
		if user.Id != subject.Teacher.Id {
			return nil, constants.ErrNoPrivileges
		}
		schedule = core.Schedule{
			SubjectId:   subject.Id,
			StudentId:   student.Id,
			IsPaid:      false,
			CabinetId:   cab.Id,
			IsConfirmed: true,
			IsFinished:  false,
			WhenTime:    params.WhenTime,
		}
	} else {
		schedule = core.Schedule{
			SubjectId:   subject.Id,
			CabinetId:   cab.Id,
			StudentId:   student.Id,
			IsPaid:      false,
			IsConfirmed: true,
			IsFinished:  false,
			WhenTime:    params.WhenTime,
		}
	}
	// Updating
	//cabinet
	err = u.repoCabinet.DeleteCabinetTime((*cabinetTime)[0].CabinetId)
	if err != nil {
		return nil, err
	}
	if time.Unix(params.WhenTime, 0).Sub(time.Unix((*cabinetTime)[0].StartTime, 0)) >= consts.LessonTime+consts.Duration {
		cabinetTimeNew := coreCabinet.CabinetTime{
			SubjectId:  (*cabinetTime)[0].SubjectId,
			CabinetId:  (*cabinetTime)[0].CabinetId,
			StartTime:  (*cabinetTime)[0].StartTime,
			FinishTime: params.WhenTime,
		}
		_, err = u.repoCabinet.CreateCabinetTime(&cabinetTimeNew)
		if err != nil {
			return nil, err
		}
	}
	if time.Unix((*cabinetTime)[0].FinishTime, 0).Sub(time.Unix(params.WhenTime, 0).Add(consts.LessonTime+consts.Duration)) >= consts.LessonTime+consts.Duration {
		cabinetTimeNew := coreCabinet.CabinetTime{
			SubjectId:  (*cabinetTime)[0].SubjectId,
			CabinetId:  (*cabinetTime)[0].CabinetId,
			StartTime:  time.Unix(params.WhenTime, 0).Add(consts.LessonTime + consts.Duration).Unix(),
			FinishTime: (*cabinetTime)[0].FinishTime,
		}
		_, err = u.repoCabinet.CreateCabinetTime(&cabinetTimeNew)
		if err != nil {
			return nil, err
		}
	}
	// Updating
	//timeTeacher
	err = u.repoTimeTeacher.DeleteTimeTeacher((*teacherTime)[0].Id)
	if err != nil {
		return nil, err
	}
	if time.Unix(params.WhenTime, 0).Sub(time.Unix((*teacherTime)[0].StartTime, 0)) >= consts.LessonTime+consts.Duration {
		teacherTimeNew := coreTimeTeacher.TimeTeacher{
			TeacherId:  (*teacherTime)[0].TeacherId,
			StartTime:  (*teacherTime)[0].StartTime,
			FinishTime: params.WhenTime,
		}
		_, err = u.repoTimeTeacher.CreateTimeTeacher(&teacherTimeNew)
		if err != nil {
			return nil, err
		}
	}
	if time.Unix((*teacherTime)[0].FinishTime, 0).Sub(time.Unix(params.WhenTime, 0).Add(consts.LessonTime+consts.Duration)) >= consts.LessonTime+consts.Duration {
		teacherTimeNew := coreTimeTeacher.TimeTeacher{
			TeacherId:  (*teacherTime)[0].TeacherId,
			StartTime:  time.Unix(params.WhenTime, 0).Add(consts.LessonTime + consts.Duration).Unix(),
			FinishTime: (*cabinetTime)[0].FinishTime,
		}
		_, err = u.repoTimeTeacher.CreateTimeTeacher(&teacherTimeNew)
		if err != nil {
			return nil, err
		}
	}

	res, err := u.repoSchedule.CreateSchedule(&schedule)
	if err != nil {
		return nil, err
	}
	return &dto.CreateScheduleResponse{Schedule: convert.Schedule2DTO(res, subject.Subject, student.User, cab.Name)}, nil
}

func (u ScheduleUseCase) UpdateSchedule(params *dto.UpdateScheduleRequest) (*dto.UpdateScheduleResponse, error) {
	if err := utils.IsTime15MinDuration(params.WhenTime); err != nil {
		return nil, err
	}
	schedule, err := u.repoSchedule.GetScheduleById(params.ScheduleId)
	if err != nil {
		return nil, err
	}
	subject, err := u.SubjectUC.GetSubject(&dtoSubject.GetSubjectRequest{Id: schedule.SubjectId})
	if err != nil {
		return nil, err
	}
	_, err = u.checkUser(params.UserId, subject.Teacher.Id)
	if err != nil {
		return nil, err
	}
	student, err := u.UserUC.GetUser(&dtoUser.GetUserRequest{Id: params.StudentId})
	if err != nil {
		return nil, constants.NewCodedError(fmt.Sprintf("User %s is not student", params.StudentId), fiber.StatusConflict)
	}
	if student.Role != consts.Student {
		return nil, constants.NewCodedError(fmt.Sprintf("User %s is not student", params.StudentId), fiber.StatusConflict)
	}
	cabinetId := params.CabinetId
	if schedule.WhenTime != params.WhenTime {
		var cabinetTime *[]coreCabinet.CabinetTime
		if subject.CabinetId == 0 {
			if params.CabinetId == 0 {
				cabinetTime, err = u.repoCabinet.GetAllCabinetTimes(params.WhenTime, time.Unix(params.WhenTime, 0).Add(consts.LessonTime+consts.Duration).Unix())
				if err != nil {
					return nil, err
				}
			} else {
				cabinetTime, err = u.repoCabinet.GetCabinetTimes(params.CabinetId, params.WhenTime, time.Unix(params.WhenTime, 0).Add(consts.LessonTime+consts.Duration).Unix())
				if err != nil {
					return nil, err
				}
			}
		} else {
			cabinetTime, err = u.repoCabinet.GetCabinetTimes(subject.CabinetId, params.WhenTime, time.Unix(params.WhenTime, 0).Add(consts.LessonTime+consts.Duration).Unix())
			if err != nil {
				return nil, err
			}
		}
		teacherTime, err := u.repoTimeTeacher.GetTimeTeachers(subject.Teacher.Id, params.WhenTime, time.Unix(params.WhenTime, 0).Add(consts.LessonTime+consts.Duration).Unix())
		if err != nil {
			return nil, err
		}
		// check
		_, err = u.repoSchedule.GetSubjectSchedules(subject.Id, params.WhenTime, time.Unix(params.WhenTime, 0).Add(consts.LessonTime+consts.Duration).Unix())
		if err != constants.ErrScheduleDBNotFound {
			return nil, err
		}
		// Updating
		//cabinet
		err = u.repoCabinet.DeleteCabinetTime((*cabinetTime)[0].CabinetId)
		if err != nil {
			return nil, err
		}
		if time.Unix(params.WhenTime, 0).Sub(time.Unix((*cabinetTime)[0].StartTime, 0)) >= consts.LessonTime+consts.Duration {
			cabinetTimeNew := coreCabinet.CabinetTime{
				SubjectId:  (*cabinetTime)[0].SubjectId,
				CabinetId:  (*cabinetTime)[0].CabinetId,
				StartTime:  (*cabinetTime)[0].StartTime,
				FinishTime: params.WhenTime,
			}
			_, err = u.repoCabinet.CreateCabinetTime(&cabinetTimeNew)
			if err != nil {
				return nil, err
			}
		}
		if time.Unix((*cabinetTime)[0].FinishTime, 0).Sub(time.Unix(params.WhenTime, 0).Add(consts.LessonTime+consts.Duration)) >= consts.LessonTime+consts.Duration {
			cabinetTimeNew := coreCabinet.CabinetTime{
				SubjectId:  (*cabinetTime)[0].SubjectId,
				CabinetId:  (*cabinetTime)[0].CabinetId,
				StartTime:  time.Unix(params.WhenTime, 0).Add(consts.LessonTime + consts.Duration).Unix(),
				FinishTime: (*cabinetTime)[0].FinishTime,
			}
			_, err = u.repoCabinet.CreateCabinetTime(&cabinetTimeNew)
			if err != nil {
				return nil, err
			}
		}
		cabinetTimeNew := coreCabinet.CabinetTime{
			SubjectId:  schedule.SubjectId,
			CabinetId:  schedule.CabinetId,
			StartTime:  schedule.WhenTime,
			FinishTime: time.Unix(schedule.WhenTime, 0).Add(consts.LessonTime + consts.Duration).Unix(),
		}
		_, err = u.repoCabinet.CreateCabinetTime(&cabinetTimeNew)
		if err != nil {
			return nil, err
		}
		// Updating
		//timeTeacher
		err = u.repoTimeTeacher.DeleteTimeTeacher((*teacherTime)[0].Id)
		if err != nil {
			return nil, err
		}
		if time.Unix(params.WhenTime, 0).Sub(time.Unix((*teacherTime)[0].StartTime, 0)) >= consts.LessonTime+consts.Duration {
			teacherTimeNew := coreTimeTeacher.TimeTeacher{
				TeacherId:  (*teacherTime)[0].TeacherId,
				StartTime:  (*teacherTime)[0].StartTime,
				FinishTime: params.WhenTime,
			}
			_, err = u.repoTimeTeacher.CreateTimeTeacher(&teacherTimeNew)
			if err != nil {
				return nil, err
			}
		}
		if time.Unix((*teacherTime)[0].FinishTime, 0).Sub(time.Unix(params.WhenTime, 0).Add(consts.LessonTime+consts.Duration)) >= consts.LessonTime+consts.Duration {
			teacherTimeNew := coreTimeTeacher.TimeTeacher{
				TeacherId:  (*teacherTime)[0].TeacherId,
				StartTime:  time.Unix(params.WhenTime, 0).Add(consts.LessonTime + consts.Duration).Unix(),
				FinishTime: (*cabinetTime)[0].FinishTime,
			}
			_, err = u.repoTimeTeacher.CreateTimeTeacher(&teacherTimeNew)
			if err != nil {
				return nil, err
			}
		}
		teacherTimeNew := coreTimeTeacher.TimeTeacher{
			TeacherId:  subject.Teacher.Id,
			StartTime:  schedule.WhenTime,
			FinishTime: time.Unix(schedule.WhenTime, 0).Add(consts.LessonTime + consts.Duration).Unix(),
		}
		_, err = u.repoTimeTeacher.CreateTimeTeacher(&teacherTimeNew)
		if err != nil {
			return nil, err
		}
		cabinetId = (*cabinetTime)[0].CabinetId
	}
	cab, err := u.repoCabinet.GetCabinetById(cabinetId)
	if err != nil {
		return nil, err
	}
	scheduleNew := core.Schedule{
		SubjectId:   subject.Id,
		StudentId:   student.Id,
		IsPaid:      params.IsPaid,
		CabinetId:   cabinetId,
		Description: params.Description,
		IsConfirmed: params.IsConfirmed,
		IsFinished:  params.IsFinished,
		WhenTime:    params.WhenTime,
	}

	res, err := u.repoSchedule.UpdateSchedule(&scheduleNew)
	if err != nil {
		return nil, err
	}
	return &dto.UpdateScheduleResponse{Schedule: convert.Schedule2DTO(res, subject.Subject, student.User, cab.Name)}, nil
}

func (u ScheduleUseCase) DeleteSchedule(params *dto.DeleteScheduleRequest) (*dto.DeleteScheduleResponse, error) {
	schedule, err := u.repoSchedule.GetScheduleById(params.ScheduleId)
	if err != nil {
		return nil, err
	}
	subject, err := u.SubjectUC.GetSubject(&dtoSubject.GetSubjectRequest{Id: schedule.SubjectId})
	if err != nil {
		return nil, err
	}
	_, err = u.checkUser(params.UserId, subject.Teacher.Id)
	if err != nil {
		return nil, err
	}
	err = u.repoSchedule.DeleteSchedule(subject.Id)
	if err != nil {
		return nil, err
	}
	return &dto.DeleteScheduleResponse{}, nil
}

func (u ScheduleUseCase) GetSchedule(params *dto.GetScheduleRequest) (*dto.GetScheduleResponse, error) {
	res, err := u.repoSchedule.GetScheduleById(params.ScheduleId)
	if err != nil {
		return nil, err
	}
	subject, err := u.SubjectUC.GetSubject(&dtoSubject.GetSubjectRequest{Id: res.SubjectId})
	if err != nil {
		return nil, err
	}
	user, err := u.UserUC.GetUser(&dtoUser.GetUserRequest{Id: params.UserId})
	if err != nil {
		return nil, err
	}
	if user.Role == consts.Student && user.Id != res.StudentId {
		return nil, constants.ErrNoPrivileges
	}
	student, err := u.UserUC.GetUser(&dtoUser.GetUserRequest{Id: res.StudentId})
	if err != nil {
		return nil, err
	}
	cab, err := u.repoCabinet.GetCabinetById(subject.CabinetId)
	if err != nil {
		return nil, err
	}
	return &dto.GetScheduleResponse{Schedule: convert.Schedule2DTO(res, subject.Subject, student.User, cab.Name)}, nil
}

func (u ScheduleUseCase) GetStudentSchedules(params *dto.GetStudentSchedulesRequest) (*dto.GetStudentSchedulesResponse, error) {
	if err := utils.IsTime15MinDuration(params.StartTime, params.FinishTime); err != nil {
		return nil, err
	}
	user, err := u.UserUC.GetUser(&dtoUser.GetUserRequest{Id: params.UserId})
	if err != nil {
		return nil, err
	}
	if user.Role == consts.Student && user.Id != params.StudentId {
		return nil, constants.ErrNoPrivileges
	}
	res, err := u.repoSchedule.GetStudentSchedules(params.StudentId, params.StartTime, params.FinishTime)
	if err != nil {
		return nil, err
	}
	var result []dto.Schedule
	for _, i := range *res {
		subject, err := u.SubjectUC.GetSubject(&dtoSubject.GetSubjectRequest{Id: i.SubjectId})
		if err != nil {
			return nil, err
		}
		student, err := u.UserUC.GetUser(&dtoUser.GetUserRequest{Id: params.UserId})
		if err != nil {
			return nil, err
		}
		cab, err := u.repoCabinet.GetCabinetById(subject.CabinetId)
		if err != nil {
			return nil, err
		}
		result = append(result, convert.Schedule2DTO(&i, subject.Subject, student.User, cab.Name))
	}
	return &dto.GetStudentSchedulesResponse{Schedules: result}, nil
}

func (u ScheduleUseCase) GetTeacherSchedules(params *dto.GetTeacherSchedulesRequest) (*dto.GetTeacherSchedulesResponse, error) {
	if err := utils.IsTime15MinDuration(params.StartTime, params.FinishTime); err != nil {
		return nil, err
	}
	user, err := u.UserUC.GetUser(&dtoUser.GetUserRequest{Id: params.UserId})
	if err != nil {
		return nil, err
	}
	if user.Role == consts.Student {
		return nil, constants.ErrNoPrivileges
	}
	if user.Role == consts.Teacher && user.Id != params.TeacherId {
		return nil, constants.ErrNoPrivileges
	}
	res, err := u.repoSchedule.GetTeacherSchedules(params.TeacherId, params.StartTime, params.FinishTime)
	if err != nil {
		return nil, err
	}
	var result []dto.Schedule
	for _, i := range *res {
		subject, err := u.SubjectUC.GetSubject(&dtoSubject.GetSubjectRequest{Id: i.SubjectId})
		if err != nil {
			return nil, err
		}
		student, err := u.UserUC.GetUser(&dtoUser.GetUserRequest{Id: params.UserId})
		if err != nil {
			return nil, err
		}
		cab, err := u.repoCabinet.GetCabinetById(subject.CabinetId)
		if err != nil {
			return nil, err
		}
		result = append(result, convert.Schedule2DTO(&i, subject.Subject, student.User, cab.Name))
	}
	return &dto.GetTeacherSchedulesResponse{Schedules: result}, nil
}

func (u ScheduleUseCase) GetSlotsSchedules(params *dto.GetSlotsRequest) (*dto.GetSlotsResponse, error) {
	if err := utils.IsTime15MinDuration(params.StartTime, params.FinishTime); err != nil {
		return nil, err
	}
	var result []int64
	subject, err := u.SubjectUC.GetSubject(&dtoSubject.GetSubjectRequest{Id: params.SubjectId})
	if err != nil {
		return nil, err
	}
	teacherTime, err := u.repoTimeTeacher.GetTimeTeachers(subject.Teacher.Id, params.StartTime, params.FinishTime)
	if err != nil {
		return nil, err
	}
	var cabinetTime *[]coreCabinet.CabinetTime
	if subject.CabinetId == 0 {
		cabinetTime, err = u.repoCabinet.GetAllCabinetTimes(params.StartTime, params.FinishTime)
		if err != nil {
			return nil, err
		}
	} else {
		cabinetTime, err = u.repoCabinet.GetCabinetTimes(subject.CabinetId, params.StartTime, params.FinishTime)
		if err != nil {
			return nil, err
		}
	}
	for _, i := range *teacherTime {
		for _, j := range *cabinetTime {
			isHaveCommonTime, firstTime, secondTime := utils.FindCommonTime(i.StartTime, i.FinishTime, j.StartTime, j.FinishTime)
			if !isHaveCommonTime {
				continue
			}
			if time.Unix(secondTime, 0).Sub(time.Unix(firstTime, 0)) < consts.LessonTime+consts.Duration {
				continue
			}
			for k := 0 * time.Minute; time.Unix(secondTime, 0).Sub(time.Unix(firstTime, 0)) > k; k += consts.LessonTime + consts.Duration {
				result = append(result, time.Unix(firstTime, 0).Add(k).Unix())
			}
		}
	}
	return &dto.GetSlotsResponse{WhenTime: utils.RemoveDuplicate(result)}, nil
}
