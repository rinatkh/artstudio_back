package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/internal/middleware"
)

func MapCabinetRoutes(router fiber.Router, h *CabinetHandler, mw *middleware.MDWManager) {
	router.Get("/cabinet/cabinet_id", mw.VerifyTokenMiddleware(), h.GetCabinet())
	router.Put("/cabinet/:cabinet_id", mw.VerifyTokenMiddleware(), mw.VerifyAdminMiddleware(), h.UpdateCabinet())
	router.Delete("/cabinet/:cabinet_id", mw.VerifyTokenMiddleware(), mw.VerifyAdminMiddleware(), h.DeleteCabinet())
	router.Post("/cabinet", mw.VerifyTokenMiddleware(), mw.VerifyAdminMiddleware(), h.CreateCabinet())
}
