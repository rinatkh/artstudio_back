package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/config"
	"github.com/rinatkh/artstudio_back/internal/static"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/rinatkh/artstudio_back/pkg/utils"
	"github.com/sirupsen/logrus"
)

type StaticHandler struct {
	cfg      *config.Config
	staticUC static.UseCase
	log      *logrus.Entry
}

func NewStaticHandler(staticUC static.UseCase, log *logrus.Entry, cfg *config.Config) *StaticHandler {
	return &StaticHandler{
		staticUC: staticUC,
		log:      log,
		cfg:      cfg,
	}
}

func (u StaticHandler) UploadPhoto() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		image, err := ctx.FormFile("image")
		if err != nil {
			return err
		}
		data, err := u.staticUC.UploadImage(image)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u StaticHandler) UploadFile() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		file, err := ctx.FormFile("file")
		if err != nil {
			return err
		}
		data, err := u.staticUC.UploadFile(file)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u StaticHandler) GetFile() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params static.GetFileRequest
		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}
		return ctx.SendFile("/opt/files" + params.URL)
	}
}
