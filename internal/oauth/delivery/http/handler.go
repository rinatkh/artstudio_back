package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/config"
	"github.com/rinatkh/artstudio_back/internal/oauth"
	"github.com/rinatkh/artstudio_back/internal/oauth/models/dto"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/rinatkh/artstudio_back/pkg/utils"
	"github.com/sirupsen/logrus"
)

type OauthHandler struct {
	cfg     *config.Config
	oauthUC oauth.UseCase
	log     *logrus.Entry
}

func NewOauthHandler(oauthUC oauth.UseCase, log *logrus.Entry, cfg *config.Config) *OauthHandler {
	return &OauthHandler{
		oauthUC: oauthUC,
		log:     log,
		cfg:     cfg,
	}
}

func (u OauthHandler) AuthenticateThroughTelergam() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params dto.AuthenticateThroughTelergamRequest
		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}

		err := u.oauthUC.AuthenticateThroughTelergam(&params)
		if err != nil {
			return err
		}
		authToken, err := utils.GenerateAuthToken(&utils.AuthTokenWrapper{UserID: params.ID}, u.cfg)
		if err != nil {
			return err
		}
		ctx.Cookie(utils.CreateHttpOnlyCookie(constants.CookieKeyAuthToken, authToken, u.cfg.Service.JwtTtl))
		return ctx.Redirect("/", fiber.StatusPermanentRedirect)
	}
}
