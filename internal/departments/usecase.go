package departments

import "github.com/rinatkh/artstudio_back/internal/departments/models/dto"

type UseCase interface {
	CreateDepartment(params *dto.CreateDepartmentRequest) (*dto.CreateDepartmentResponse, error)
	UpdateDepartment(params *dto.UpdateDepartmentRequest) (*dto.UpdateDepartmentResponse, error)
	DeleteDepartment(params *dto.DeleteDepartmentRequest) (*dto.DeleteDepartmentResponse, error)
	GetDepartment(params *dto.GetDepartmentRequest) (*dto.GetDepartmentResponse, error)
	GetDepartments(params *dto.GetDepartmentsRequest) (*dto.GetDepartmentsResponse, error)

	AddSubjDepartment(params *dto.UpdateDepartmentSubjectRequest) (*dto.UpdateDepartmentSubjectResponse, error)
	DelSubDepartment(params *dto.UpdateDepartmentSubjectRequest) (*dto.UpdateDepartmentSubjectResponse, error)
}
