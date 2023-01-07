package http

import (
	"github.com/gofiber/fiber/v2"
)

func MapSubjectRoutes(router fiber.Router, h *SubjectHandler) {
	router.Get("/subjects/:subject_id", h.GetSubject())
	router.Put("/subjects/:subject_id", h.UpdateSubject())
	router.Delete("/subjects/:subject_id", h.DeleteSubject())
	router.Get("/subjects", h.GetSubjects())
	router.Post("/subjects", h.CreateSubject())
}
