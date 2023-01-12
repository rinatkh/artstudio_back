package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/internal/middleware"
)

func MapSchedulesRoutes(router fiber.Router, h *ScheduleHandler, mw *middleware.MDWManager) {
	router.Put("/schedules/:schedule_id", mw.VerifyTokenMiddleware(), h.UpdateSchedule())
	router.Delete("/schedules/:schedule_id", mw.VerifyTokenMiddleware(), h.DeleteSchedule())
	router.Get("/schedules", mw.VerifyTokenMiddleware(), h.GetSchedules())
	router.Post("/schedules", mw.VerifyTokenMiddleware(), h.CreateSchedules())
}

// от студента отправить заявку учителю
//
