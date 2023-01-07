package usecase

import (
	"github.com/rinatkh/artstudio_back/config"
	"github.com/rinatkh/artstudio_back/internal/departments"
	"github.com/rinatkh/artstudio_back/internal/departments/models/convert"
	"github.com/rinatkh/artstudio_back/internal/departments/models/core"
	"github.com/rinatkh/artstudio_back/internal/departments/models/dto"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/sirupsen/logrus"
)

type DepartmentUseCase struct {
	cfg            *config.Config
	log            *logrus.Entry
	repoDepartment departments.DepartmentRepository
}

func NewDepartmentUC(cfg *config.Config, log *logrus.Entry, repoDepartment departments.DepartmentRepository) departments.UseCase {
	return &DepartmentUseCase{
		cfg:            cfg,
		log:            log,
		repoDepartment: repoDepartment,
	}
}

func (u DepartmentUseCase) CreateDepartment(params *dto.CreateDepartmentRequest) (*dto.CreateDepartmentResponse, error) {
	department := core.Department{
		Name:        params.Name,
		Description: &params.Description,
		Image:       params.Image,
	}
	result, err := u.repoDepartment.CreateDepartment(&department)
	if err != nil {
		return nil, err
	}

	return &dto.CreateDepartmentResponse{Department: convert.Department2DTO(result)}, nil
}

func (u DepartmentUseCase) UpdateDepartment(params *dto.UpdateDepartmentRequest) (*dto.UpdateDepartmentResponse, error) {
	check, err := u.repoDepartment.GetDepartmentById(params.Id)
	if err != nil {
		return nil, err
	}
	if check == nil {
		return nil, constants.ErrDepartmentDBNotFound
	}
	department := core.Department{
		Name:        params.Name,
		Description: &params.Description,
		Image:       params.Image,
	}
	result, err := u.repoDepartment.UpdateDepartment(&department)
	if err != nil {
		return nil, err
	}

	return &dto.UpdateDepartmentResponse{Department: convert.Department2DTO(result)}, nil
}

func (u DepartmentUseCase) DeleteDepartment(params *dto.DeleteDepartmentRequest) (*dto.DeleteDepartmentResponse, error) {
	check, err := u.repoDepartment.GetDepartmentById(params.Id)
	if err != nil {
		return nil, err
	}
	if check == nil {
		return nil, constants.ErrDepartmentDBNotFound
	}
	err = u.repoDepartment.DeleteDepartment(params.Id)
	if err != nil {
		return nil, err
	}
	return &dto.DeleteDepartmentResponse{}, nil
}

func (u DepartmentUseCase) GetDepartment(params *dto.GetDepartmentRequest) (*dto.GetDepartmentResponse, error) {
	result, err := u.repoDepartment.GetDepartmentById(params.Id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, constants.ErrDepartmentDBNotFound
	}
	return &dto.GetDepartmentResponse{Department: convert.Department2DTO(result)}, nil
}

func (u DepartmentUseCase) GetDepartments(params *dto.GetDepartmentsRequest) (*dto.GetDepartmentsResponse, error) {
	list, length, err := u.repoDepartment.GetDepartments(params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}
	if list == nil {
		return &dto.GetDepartmentsResponse{}, nil
	}
	var result []dto.Department
	for _, i := range *list {
		result = append(result, convert.Department2DTO(&i))
	}
	return &dto.GetDepartmentsResponse{Departments: result, Length: length}, nil
}
