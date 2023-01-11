package usecase

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/config"
	"github.com/rinatkh/artstudio_back/internal/DepartmentSubjects"
	"github.com/rinatkh/artstudio_back/internal/departments"
	"github.com/rinatkh/artstudio_back/internal/subjects"
	dtoSubject "github.com/rinatkh/artstudio_back/internal/subjects/models/dto"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/sirupsen/logrus"
)

type DepartmentSubjectsUseCase struct {
	cfg                    *config.Config
	log                    *logrus.Entry
	repoDepartmentSubjects departmentSubjects.DepartmentSubjectsRepository
	subjectsUC             subjects.UseCase
	repoDepartment         departments.DepartmentRepository
}

func NewDepartmentSubjectsUC(cfg *config.Config, log *logrus.Entry, repoDepartmentSubjects departmentSubjects.DepartmentSubjectsRepository, subjectsUC subjects.UseCase, repoDepartment departments.DepartmentRepository) departmentSubjects.UseCase {
	return &DepartmentSubjectsUseCase{
		cfg:                    cfg,
		log:                    log,
		repoDepartmentSubjects: repoDepartmentSubjects,
		subjectsUC:             subjectsUC,
		repoDepartment:         repoDepartment,
	}
}

func (u DepartmentSubjectsUseCase) AddDepartmentSubjects(params *departmentSubjects.AddDepartmentSubjectsRequest) (*departmentSubjects.AddDepartmentSubjectsResponse, error) {
	_, err := u.subjectsUC.GetSubject(&dtoSubject.GetSubjectRequest{
		Id: params.SubjectId,
	})
	if err != nil {
		return nil, constants.ErrSubjectDBNotFound
	}
	dep, err := u.repoDepartment.GetDepartmentById(params.DepartmentId)
	if err != nil {
		return nil, constants.ErrDepartmentDBNotFound
	}
	if dep == nil {
		return nil, constants.ErrDepartmentDBNotFound
	}
	isExist, err := u.repoDepartmentSubjects.IsDepartmentSubjects(params.DepartmentId, params.SubjectId)
	if err != nil {
		return nil, err
	}
	if isExist {
		return nil, constants.NewCodedError(fmt.Sprintf("Departament %d already have Subject %d", params.DepartmentId, params.SubjectId), fiber.StatusConflict)
	}
	return &departmentSubjects.AddDepartmentSubjectsResponse{}, u.repoDepartmentSubjects.AddDepartmentSubjects(params.DepartmentId, params.SubjectId)
}
func (u DepartmentSubjectsUseCase) DeleteDepartmentSubjects(params *departmentSubjects.DeleteDepartmentSubjectsRequest) (*departmentSubjects.DeleteDepartmentSubjectsResponse, error) {
	_, err := u.subjectsUC.GetSubject(&dtoSubject.GetSubjectRequest{
		Id: params.SubjectId,
	})
	if err != nil {
		return nil, constants.ErrSubjectDBNotFound
	}
	dep, err := u.repoDepartment.GetDepartmentById(params.DepartmentId)
	if err != nil {
		return nil, constants.ErrDepartmentDBNotFound
	}
	if dep == nil {
		return nil, constants.ErrDepartmentDBNotFound
	}
	isExist, err := u.repoDepartmentSubjects.IsDepartmentSubjects(params.DepartmentId, params.SubjectId)
	if err != nil {
		return nil, err
	}
	if !isExist {
		return nil, constants.NewCodedError(fmt.Sprintf("Departament %d already don't have Subject %d", params.DepartmentId, params.SubjectId), fiber.StatusConflict)
	}
	return &departmentSubjects.DeleteDepartmentSubjectsResponse{}, u.repoDepartmentSubjects.DeleteDepartmentSubjects(params.DepartmentId, params.SubjectId)
}
func (u DepartmentSubjectsUseCase) GetDepartmentSubjects(params *departmentSubjects.GetDepartmentSubjectsRequest) (*departmentSubjects.GetDepartmentSubjectsResponse, error) {
	res, length, err := u.repoDepartmentSubjects.GetDepartmentSubjects(params.DepartmentId, params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return &departmentSubjects.GetDepartmentSubjectsResponse{}, nil
	}
	return &departmentSubjects.GetDepartmentSubjectsResponse{
		DepartmentSubjects: *res,
		Length:             length,
	}, nil
}

func (u DepartmentSubjectsUseCase) DeleteAll(params *departmentSubjects.DeleteAllRequest) (*departmentSubjects.DeleteAllResponse, error) {
	err := u.repoDepartmentSubjects.DeleteAll(params.DepartmentId)
	if err != nil {
		return nil, err
	}
	return &departmentSubjects.DeleteAllResponse{}, nil
}
