package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/internal/middleware"
)

func MapStatisticTeacherRoutes(router fiber.Router, h *StatisticTeacherHandler, mw *middleware.MDWManager) {
	router.Get("/statistic_teachers/:teacher_id", mw.VerifyTokenMiddleware(), h.GetStatisticTeacher())
}
