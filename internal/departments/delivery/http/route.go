package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/internal/middleware"
)

func MapDepartmentRoutes(router fiber.Router, h *DepartmentHandler, mw *middleware.MDWManager) {
	router.Get("/departments/:department_id", h.GetDepartment())
	router.Put("/departments/:department_id", mw.VerifyTokenMiddleware(), mw.VerifyAdminMiddleware(), h.UpdateDepartment())
	router.Put("/departments/:department_id/add", mw.VerifyTokenMiddleware(), mw.VerifyAdminMiddleware(), h.AddSubjDepartment())
	router.Put("/departments/:department_id/delete", mw.VerifyTokenMiddleware(), mw.VerifyAdminMiddleware(), h.DelSubDepartment())
	router.Delete("/departments/:department_id", mw.VerifyTokenMiddleware(), mw.VerifyAdminMiddleware(), h.DeleteDepartment())
	router.Get("/departments", h.GetDepartments())
	router.Post("/departments", mw.VerifyTokenMiddleware(), mw.VerifyAdminMiddleware(), h.CreateDepartment())
}
