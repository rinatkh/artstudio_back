package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/internal/middleware"
)

func MapTimeTeacherRoutes(router fiber.Router, h *TimeTeacherHandler, mw *middleware.MDWManager) {
	router.Get("/time_teachers/:time_teacher_id", mw.VerifyTokenMiddleware(), h.GetTimeTeacher())
	router.Put("/time_teachers/:time_teacher_id", mw.VerifyTokenMiddleware(), h.UpdateTimeTeacher())
	router.Delete("/time_teachers/:time_teacher_id", mw.VerifyTokenMiddleware(), h.DeleteTimeTeacher())
	router.Get("/time_teachers", mw.VerifyTokenMiddleware(), h.GetTimeTeachers())
	router.Post("/time_teachers", mw.VerifyTokenMiddleware(), h.CreateTimeTeacher())
}
