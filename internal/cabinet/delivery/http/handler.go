package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/internal/cabinet"
	"github.com/rinatkh/artstudio_back/internal/cabinet/models/dto"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/rinatkh/artstudio_back/pkg/utils"
	"github.com/sirupsen/logrus"
	"strconv"
)

type CabinetHandler struct {
	cabinetUC cabinet.UseCase
	log       *logrus.Entry
}

func NewCabinetHandler(cabinetUC cabinet.UseCase, log *logrus.Entry) *CabinetHandler {
	return &CabinetHandler{
		cabinetUC: cabinetUC,
		log:       log,
	}
}

func (u CabinetHandler) GetCabinet() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.GetCabinetRequest
		var err error
		params.CabinetId, err = strconv.ParseInt(ctx.Params("cabinet_id"), 10, 64)
		if err != nil {
			return err
		}
		data, err := u.cabinetUC.GetCabinet(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u CabinetHandler) DeleteCabinet() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.DeleteCabinetRequest

		var err error
		params.CabinetId, err = strconv.ParseInt(ctx.Params("cabinet_id"), 10, 64)
		if err != nil {
			return err
		}

		data, err := u.cabinetUC.DeleteCabinet(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u CabinetHandler) UpdateCabinet() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.UpdateCabinetRequest

		var err error
		params.CabinetId, err = strconv.ParseInt(ctx.Params("cabinet_id"), 10, 64)
		if err != nil {
			return err
		}

		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}

		data, err := u.cabinetUC.UpdateCabinet(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u CabinetHandler) CreateCabinet() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.CreateCabinetRequest

		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}

		data, err := u.cabinetUC.CreateCabinet(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}
