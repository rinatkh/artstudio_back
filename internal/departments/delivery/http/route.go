package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/internal/middleware"
)

func MapDepartmentRoutes(router fiber.Router, h *DepartmentHandler, mw *middleware.MDWManager) {
	router.Get("/departments/:department_id", h.GetDepartment())
	router.Put("/departments/:department_id", h.UpdateDepartment())
	router.Delete("/departments/:department_id", h.DeleteDepartment())
	router.Get("/departments", h.GetDepartments())
	router.Post("/departments", h.CreateDepartment())
}
