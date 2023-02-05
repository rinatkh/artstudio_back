package http

import (
	"github.com/gofiber/fiber/v2"
	timeTeacher "github.com/rinatkh/artstudio_back/internal/time_teachers"
	"github.com/rinatkh/artstudio_back/internal/time_teachers/models/dto"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/rinatkh/artstudio_back/pkg/utils"
	"github.com/sirupsen/logrus"
	"strconv"
)

type TimeTeacherHandler struct {
	timeTeacherUC timeTeacher.UseCase
	log           *logrus.Entry
}

func NewTimeTeacherHandler(timeTeacherUC timeTeacher.UseCase, log *logrus.Entry) *TimeTeacherHandler {
	return &TimeTeacherHandler{
		timeTeacherUC: timeTeacherUC,
		log:           log,
	}
}

func (u TimeTeacherHandler) DeleteTimeTeacher() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.DeleteTimeTeacherRequest
		var err error
		params.Id, err = strconv.ParseInt(ctx.Params("time_teacher_id"), 10, 64)
		if err != nil {
			return err
		}
		params.UserId = ctx.Get(constants.CtxKeyUserID)
		params.TeacherId = ctx.Query("teacher_id")

		data, err := u.timeTeacherUC.DeleteTimeTeacher(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u TimeTeacherHandler) UpdateTimeTeacher() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.UpdateTimeTeacherRequest
		var err error
		params.Id, err = strconv.ParseInt(ctx.Params("time_teacher_id"), 10, 64)
		if err != nil {
			return err
		}
		params.UserId = ctx.Get(constants.CtxKeyUserID)
		params.TeacherId = ctx.Query("teacher_id")
		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}

		data, err := u.timeTeacherUC.UpdateTimeTeacher(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u TimeTeacherHandler) GetTimeTeacher() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.GetTimeTeacherRequest
		var err error
		params.Id, err = strconv.ParseInt(ctx.Params("time_teacher_id"), 10, 64)
		if err != nil {
			return err
		}
		params.UserId = ctx.Get(constants.CtxKeyUserID)

		data, err := u.timeTeacherUC.GetTimeTeacher(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u TimeTeacherHandler) GetTimeTeachers() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.GetTimeTeachersRequest

		params.UserId = ctx.Get(constants.CtxKeyUserID)
		params.TeacherId = ctx.Query("teacher_id")
		var err error
		params.StartTime, err = strconv.ParseInt(ctx.Query("start_time"), 10, 64)
		if err != nil {
			return err
		}
		params.StartTime, err = strconv.ParseInt(ctx.Query("finish_time"), 10, 64)
		if err != nil {
			return err
		}

		data, err := u.timeTeacherUC.GetTimeTeachers(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u TimeTeacherHandler) CreateTimeTeacher() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.CreateTimeTeacherRequest
		params.UserId = ctx.Get(constants.CtxKeyUserID)
		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}

		data, err := u.timeTeacherUC.CreateTimeTeacher(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}
