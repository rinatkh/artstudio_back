package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/internal/middleware"
)

func MapCabinetRoutes(router fiber.Router, h *CabinetHandler, mw *middleware.MDWManager) {
	router.Get("/cabinets/:cabinet_id", mw.VerifyTokenMiddleware(), h.GetCabinet())
	router.Put("/cabinets/:cabinet_id", mw.VerifyTokenMiddleware(), mw.VerifyAdminMiddleware(), h.UpdateCabinet())
	router.Delete("/cabinets/:cabinet_id", mw.VerifyTokenMiddleware(), mw.VerifyAdminMiddleware(), h.DeleteCabinet())
	router.Post("/cabinets", mw.VerifyTokenMiddleware(), mw.VerifyAdminMiddleware(), h.CreateCabinet())
}
