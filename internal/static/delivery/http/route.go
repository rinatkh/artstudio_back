package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/internal/middleware"
)

func MaStaticRoutes(router fiber.Router, h *StaticHandler, mw *middleware.MDWManager) {
	router.Post("/file/upload", mw.VerifyTokenMiddleware(), h.UploadFile())
	router.Get("/file/:file_id", mw.VerifyTokenMiddleware(), h.GetFile())
	router.Post("/photo/upload", mw.VerifyTokenMiddleware(), h.UploadPhoto())
}
