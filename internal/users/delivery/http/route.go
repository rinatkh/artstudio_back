package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/internal/middleware"
)

func MapUserRoutes(router fiber.Router, h *UserHandler, mw *middleware.MDWManager) {
	router.Get("/users/:user_id", mw.VerifyTokenMiddleware(), h.GetUser())
	router.Put("/users/:user_id", mw.VerifyTokenMiddleware(), h.UpdateUser())
	router.Delete("/users/:user_id", mw.VerifyTokenMiddleware(), h.DeleteUser())
	router.Get("/users", mw.VerifyTokenMiddleware(), h.GetUsers())
}
