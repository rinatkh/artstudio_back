package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/internal/departments"
	"github.com/rinatkh/artstudio_back/internal/departments/models/dto"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/rinatkh/artstudio_back/pkg/utils"
	"github.com/sirupsen/logrus"
	"strconv"
)

type DepartmentHandler struct {
	departmentUC departments.UseCase
	log          *logrus.Entry
}

func NewDepartmentHandler(departmentUC departments.UseCase, log *logrus.Entry) *DepartmentHandler {
	return &DepartmentHandler{
		departmentUC: departmentUC,
		log:          log,
	}
}

func (u DepartmentHandler) DeleteDepartment() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.DeleteDepartmentRequest
		var err error
		params.Id, err = strconv.ParseInt(ctx.Params("department_id"), 10, 64)
		if err != nil {
			return err
		}
		params.UserId = ctx.Get(constants.CtxKeyUserID)
		data, err := u.departmentUC.DeleteDepartment(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u DepartmentHandler) UpdateDepartment() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.UpdateDepartmentRequest
		var err error
		params.UserId = ctx.Get(constants.CtxKeyUserID)
		params.Id, err = strconv.ParseInt(ctx.Params("department_id"), 10, 64)
		if err != nil {
			return err
		}
		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}

		data, err := u.departmentUC.UpdateDepartment(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u DepartmentHandler) GetDepartment() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.GetDepartmentRequest
		var err error
		params.Id, err = strconv.ParseInt(ctx.Params("department_id"), 10, 64)
		if err != nil {
			return err
		}
		params.LimitSubjects, err = strconv.ParseInt(ctx.Query("limit_subjects", "20"),
			10, 64)
		if err != nil {
			return err
		}
		params.OffsetSubjects, err = strconv.ParseInt(ctx.Query("offset_subjects", "0"),
			10, 64)
		if err != nil {
			return err
		}
		data, err := u.departmentUC.GetDepartment(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u DepartmentHandler) GetDepartments() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.GetDepartmentsRequest
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
		params.LimitSubjects, err = strconv.ParseInt(ctx.Query("limit_subjects", "20"),
			10, 64)
		if err != nil {
			return err
		}
		params.OffsetSubjects, err = strconv.ParseInt(ctx.Query("offset_subjects", "0"),
			10, 64)
		if err != nil {
			return err
		}
		data, err := u.departmentUC.GetDepartments(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u DepartmentHandler) CreateDepartment() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.CreateDepartmentRequest
		params.UserId = ctx.Get(constants.CtxKeyUserID)
		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}

		data, err := u.departmentUC.CreateDepartment(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}
