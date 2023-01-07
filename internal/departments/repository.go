package departments

import "github.com/rinatkh/artstudio_back/internal/departments/models/core"

type DepartmentRepository interface {
	GetDepartmentById(id int64) (*core.Department, error)
	GetDepartments(limit, offset int64) (*[]core.Department, int64, error)
	DeleteDepartment(id int64) error
	CreateDepartment(Department *core.Department) (*core.Department, error)
	UpdateDepartment(Department *core.Department) (*core.Department, error)
}
