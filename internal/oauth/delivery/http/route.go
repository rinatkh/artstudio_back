package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/internal/middleware"
)

func MapOauthRoutes(router fiber.Router, h *OauthHandler, mw *middleware.MDWManager) {
	router.Get("/oauth/telegram", h.AuthenticateThroughTelergam(), mw.OAuthTelegramMiddleware())
}
