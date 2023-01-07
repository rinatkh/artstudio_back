package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/internal/subjects"
	"github.com/rinatkh/artstudio_back/internal/subjects/models/dto"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/rinatkh/artstudio_back/pkg/utils"
	"github.com/sirupsen/logrus"
	"strconv"
)

type SubjectHandler struct {
	subjectUC subjects.UseCase
	log       *logrus.Entry
}

func NewSubjectHandler(subjectUC subjects.UseCase, log *logrus.Entry) *SubjectHandler {
	return &SubjectHandler{
		subjectUC: subjectUC,
		log:       log,
	}
}

func (u SubjectHandler) DeleteSubject() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.DeleteSubjectRequest
		var err error
		params.Id, err = strconv.ParseInt(ctx.Params("subject_id"), 10, 64)
		if err != nil {
			return err
		}
		data, err := u.subjectUC.DeleteSubject(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u SubjectHandler) UpdateSubject() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.UpdateSubjectRequest
		var err error
		params.Id, err = strconv.ParseInt(ctx.Params("subject_id"), 10, 64)
		if err != nil {
			return err
		}
		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}

		data, err := u.subjectUC.UpdateSubject(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u SubjectHandler) GetSubject() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.GetSubjectRequest
		var err error
		params.Id, err = strconv.ParseInt(ctx.Params("subject_id"), 10, 64)
		if err != nil {
			return err
		}
		data, err := u.subjectUC.GetSubject(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u SubjectHandler) GetSubjects() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.GetSubjectsRequest
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
		data, err := u.subjectUC.GetSubjects(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u SubjectHandler) CreateSubject() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.CreateSubjectRequest
		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}

		data, err := u.subjectUC.CreateSubject(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}
