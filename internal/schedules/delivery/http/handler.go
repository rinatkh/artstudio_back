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
		params.Id = ctx.Params("schedule_id")
		params.ScheduleId = ctx.Get(constants.CtxKeyUserID)
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
		params.Id = ctx.Params("schedule_id")
		params.ScheduleId = ctx.Get(constants.CtxKeyUserID)
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

func (u ScheduleHandler) CreateSchedules() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.GetScheduleRequest
		params.Id = ctx.Params("schedule_id")
		if params.Id == "me" {
			params.Id = ctx.Get(constants.CtxKeyUserID)
		}
		data, err := u.scheduleUC.GetSchedule(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u ScheduleHandler) GetSchedules() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.GetchedulesRequest
		var err error
		params.Limit, err = strconv.ParseInt(ctx.Query("limit", "20"),
			10, 64)
		if err != nil {
			return err
		}
		params.Offset, err = strconv.ParseInt(ctx.Query("offset", "0"),
			10, 64)
		if err != nil {
			return err
		}
		data, err := u.scheduleUC.Getchedules(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}
