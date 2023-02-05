package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/internal/statistic_teachers"
	"github.com/rinatkh/artstudio_back/internal/statistic_teachers/models/dto"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/sirupsen/logrus"
)

type StatisticTeacherHandler struct {
	statisticTeacherUC statisticTeacher.UseCase
	log                *logrus.Entry
}

func NewStatisticTeacherHandler(statisticTeacherUC statisticTeacher.UseCase, log *logrus.Entry) *StatisticTeacherHandler {
	return &StatisticTeacherHandler{
		statisticTeacherUC: statisticTeacherUC,
		log:                log,
	}
}

func (u StatisticTeacherHandler) GetStatisticTeacher() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.GetStatisticTeacherRequest
		params.TeacherId = ctx.Params("teacher_id")
		params.UserId = ctx.Get(constants.CtxKeyUserID)

		data, err := u.statisticTeacherUC.GetStatisticTeacher(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}
