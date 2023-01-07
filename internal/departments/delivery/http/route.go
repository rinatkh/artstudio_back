package http

import (
	"github.com/gofiber/fiber/v2"
)

func MapDepartmentRoutes(router fiber.Router, h *DepartmentHandler) {
	router.Get("/departments/:department_id", h.GetDepartment())
	router.Put("/departments/:department_id", h.UpdateDepartment())
	router.Delete("/departments/:department_id", h.DeleteDepartment())
	router.Get("/departments", h.GetDepartments())
	router.Post("/departments", h.CreateDepartment())
}
