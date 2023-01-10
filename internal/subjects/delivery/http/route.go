package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/internal/middleware"
)

func MapSubjectRoutes(router fiber.Router, h *SubjectHandler, mw *middleware.MDWManager) {
	router.Get("/subjects/:subject_id", h.GetSubject())
	router.Put("/subjects/:subject_id", mw.VerifyTokenMiddleware(), h.UpdateSubject())
	router.Delete("/subjects/:subject_id", mw.VerifyTokenMiddleware(), h.DeleteSubject())
	router.Get("/subjects", h.GetSubjects())
	router.Post("/subjects", mw.VerifyTokenMiddleware(), h.CreateSubject())
}
