package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/internal/middleware"
)

func MapDepartmentRoutes(router fiber.Router, h *DepartmentHandler, mw *middleware.MDWManager) {
	router.Get("/departments/:department_id", h.GetDepartment())
	router.Put("/departments/:department_id", mw.VerifyTokenMiddleware(), h.UpdateDepartment())
	router.Delete("/departments/:department_id", mw.VerifyTokenMiddleware(), h.DeleteDepartment())
	router.Get("/departments", h.GetDepartments())
	router.Post("/departments", mw.VerifyTokenMiddleware(), h.CreateDepartment())
}
