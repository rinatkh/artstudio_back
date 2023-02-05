package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/internal/middleware"
)

func MapSchedulesRoutes(router fiber.Router, h *ScheduleHandler, mw *middleware.MDWManager) {
	router.Put("/schedules/:schedule_id", mw.VerifyTokenMiddleware(), h.UpdateSchedule())
	router.Delete("/schedules/:schedule_id", mw.VerifyTokenMiddleware(), h.DeleteSchedule())
	router.Get("/schedules/:schedule_id", mw.VerifyTokenMiddleware(), h.GetSchedule())
	router.Get("/schedules/teacher", mw.VerifyTokenMiddleware(), h.GetTeacherSchedules())
	router.Get("/schedules/student", mw.VerifyTokenMiddleware(), h.GetStudentSchedules())
	router.Post("/schedules", mw.VerifyTokenMiddleware(), h.CreateSchedule())

	router.Get("/slots", mw.VerifyTokenMiddleware(), h.c())
}
