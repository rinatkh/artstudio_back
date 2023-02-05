package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/internal/middleware"
)

func MapStatisticStudentRoutes(router fiber.Router, h *StatisticStudentHandler, mw *middleware.MDWManager) {
	router.Get("/statistic_students/:student_id", mw.VerifyTokenMiddleware(), h.GetStatisticStudent())
}
