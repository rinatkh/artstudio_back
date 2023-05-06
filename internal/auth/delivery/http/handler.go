package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/config"
	"github.com/rinatkh/artstudio_back/internal/auth"
	"github.com/rinatkh/artstudio_back/internal/auth/models/dto"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/rinatkh/artstudio_back/pkg/utils"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	cfg    *config.Config
	authUC auth.UseCase
	log    *logrus.Entry
}

func NewAuthHandler(authUC auth.UseCase, log *logrus.Entry, cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		authUC: authUC,
		log:    log,
		cfg:    cfg,
	}
}

func (u AuthHandler) LogoutUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		ctx.Cookie(utils.CreateHttpOnlyCookie(constants.CookieKeyAuthToken, "", 0))
		return ctx.JSON(true)
	}
}

func (u AuthHandler) LoginUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params dto.LoginUserRequest
		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}

		data, err := u.authUC.LoginUser(&params)
		if err != nil {
			return err
		}
		ctx.Cookie(utils.CreateHttpOnlyCookie(constants.CookieKeyAuthToken, data.AuthToken, u.cfg.Service.JwtTtl))
		return ctx.JSON(data)
	}
}

func (u AuthHandler) SignupUserPre() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params dto.SignupUserRequest
		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}

		data, err := u.authUC.SignupUserPre(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u AuthHandler) SignupUserConfirmed() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		data, err := u.authUC.SignupUser(ctx.Params("user_id"))
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}
