package http

import (
	"github.com/gofiber/fiber/v2"
	timeTeacher "github.com/rinatkh/artstudio_back/internal/statistic_teachers"
	"github.com/rinatkh/artstudio_back/internal/statistic_teachers/models/dto"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/sirupsen/logrus"
)

type StatisticTeacherHandler struct {
	timeTeacherUC timeTeacher.UseCase
	log           *logrus.Entry
}

func NewStatisticTeacherHandler(timeTeacherUC timeTeacher.UseCase, log *logrus.Entry) *StatisticTeacherHandler {
	return &StatisticTeacherHandler{
		timeTeacherUC: timeTeacherUC,
		log:           log,
	}
}

func (u StatisticTeacherHandler) GetStatisticTeacher() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.GetStatisticTeacherRequest
		params.TeacherId = ctx.Params("teacher_id")
		params.UserId = ctx.Get(constants.CtxKeyUserID)

		data, err := u.timeTeacherUC.GetStatisticTeacher(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}
