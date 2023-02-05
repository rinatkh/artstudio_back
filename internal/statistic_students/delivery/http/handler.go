package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/internal/statistic_students"
	"github.com/rinatkh/artstudio_back/internal/statistic_students/models/dto"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/sirupsen/logrus"
)

type StatisticStudentHandler struct {
	timeStudentUC statisticStudent.UseCase
	log           *logrus.Entry
}

func NewStatisticStudentHandler(timeStudentUC statisticStudent.UseCase, log *logrus.Entry) *StatisticStudentHandler {
	return &StatisticStudentHandler{
		timeStudentUC: timeStudentUC,
		log:           log,
	}
}

func (u StatisticStudentHandler) GetStatisticStudent() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.GetStatisticStudentRequest
		params.StudentId = ctx.Params("student_id")
		params.UserId = ctx.Get(constants.CtxKeyUserID)

		data, err := u.timeStudentUC.GetStatisticStudent(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}
