package http

import (
	"github.com/gofiber/fiber/v2"
)

func MapAuthRoutes(router fiber.Router, h *AuthHandler) {
	router.Post("/signup", h.SignupUser())
	router.Post("/login", h.LoginUser())
	router.Delete("/logout", h.LogoutUser())
}
