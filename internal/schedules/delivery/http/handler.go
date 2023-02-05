package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/internal/schedules"
	"github.com/rinatkh/artstudio_back/internal/schedules/models/dto"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/rinatkh/artstudio_back/pkg/utils"
	"github.com/sirupsen/logrus"
	"strconv"
)

type ScheduleHandler struct {
	scheduleUC schedules.UseCase
	log        *logrus.Entry
}

func NewScheduleHandler(scheduleUC schedules.UseCase, log *logrus.Entry) *ScheduleHandler {
	return &ScheduleHandler{
		scheduleUC: scheduleUC,
		log:        log,
	}
}

func (u ScheduleHandler) DeleteSchedule() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.DeleteScheduleRequest
		var err error
		params.ScheduleId, err = strconv.ParseInt(ctx.Params("schedule_id", "20"),
			10, 64)
		if err != nil {
			return err
		}
		params.UserId = ctx.Get(constants.CtxKeyUserID)
		data, err := u.scheduleUC.DeleteSchedule(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u ScheduleHandler) UpdateSchedule() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.UpdateScheduleRequest
		var err error
		params.ScheduleId, err = strconv.ParseInt(ctx.Params("schedule_id"), 10, 64)
		if err != nil {
			return err
		}
		params.UserId = ctx.Get(constants.CtxKeyUserID)
		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}

		data, err := u.scheduleUC.UpdateSchedule(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u ScheduleHandler) CreateSchedule() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.CreateScheduleRequest
		params.UserId = ctx.Get(constants.CtxKeyUserID)
		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}
		data, err := u.scheduleUC.CreateSchedule(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u ScheduleHandler) GetTeacherSchedules() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.GetTeacherSchedulesRequest
		params.UserId = ctx.Get(constants.CtxKeyUserID)
		params.TeacherId = ctx.Query("teacher_id")
		var err error
		params.StartTime, err = strconv.ParseInt(ctx.Query("start_time"),
			10, 64)
		if err != nil {
			return err
		}
		params.FinishTime, err = strconv.ParseInt(ctx.Query("finish_time"),
			10, 64)
		if err != nil {
			return err
		}
		data, err := u.scheduleUC.GetTeacherSchedules(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}
func (u ScheduleHandler) GetStudentSchedules() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.GetStudentSchedulesRequest
		params.UserId = ctx.Get(constants.CtxKeyUserID)
		params.StudentId = ctx.Query("student_id")
		var err error
		params.StartTime, err = strconv.ParseInt(ctx.Query("start_time"),
			10, 64)
		if err != nil {
			return err
		}
		params.FinishTime, err = strconv.ParseInt(ctx.Query("finish_time"),
			10, 64)
		if err != nil {
			return err
		}
		data, err := u.scheduleUC.GetStudentSchedules(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u ScheduleHandler) GetSchedule() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.GetScheduleRequest
		params.UserId = ctx.Get(constants.CtxKeyUserID)
		var err error
		params.ScheduleId, err = strconv.ParseInt(ctx.Params("schedule_id"),
			10, 64)
		if err != nil {
			return err
		}
		data, err := u.scheduleUC.GetSchedule(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u ScheduleHandler) GetSlots() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params dto.GetSlotsRequest
		params.UserId = ctx.Get(constants.CtxKeyUserID)
		var err error
		params.FinishTime, err = strconv.ParseInt(ctx.Query("finish_time"), 10, 64)
		if err != nil {
			return err
		}
		params.StartTime, err = strconv.ParseInt(ctx.Query("start_time"), 10, 64)
		if err != nil {
			return err
		}
		params.SubjectId, err = strconv.ParseInt(ctx.Query("subject_id"), 10, 64)
		if err != nil {
			return err
		}
		data, err := u.scheduleUC.GetSlotsSchedules(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}
