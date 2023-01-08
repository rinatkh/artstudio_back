package usecase

import (
	"github.com/rinatkh/artstudio_back/config"
	"github.com/rinatkh/artstudio_back/internal/DepartmentSubjects"
	"github.com/rinatkh/artstudio_back/internal/departments"
	dtoDepartment "github.com/rinatkh/artstudio_back/internal/departments/models/dto"
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
	departmentsUC          departments.UseCase
}

func NewDepartmentSubjectsUC(cfg *config.Config, log *logrus.Entry, repoDepartmentSubjects departmentSubjects.DepartmentSubjectsRepository, subjectsUC subjects.UseCase, departmentsUC departments.UseCase) departmentSubjects.UseCase {
	return &DepartmentSubjectsUseCase{
		cfg:                    cfg,
		log:                    log,
		repoDepartmentSubjects: repoDepartmentSubjects,
		subjectsUC:             subjectsUC,
		departmentsUC:          departmentsUC,
	}
}

func (u DepartmentSubjectsUseCase) AddDepartmentSubjects(params *departmentSubjects.AddDepartmentSubjectsRequest) (*departmentSubjects.AddDepartmentSubjectsResponse, error) {
	_, err := u.subjectsUC.GetSubject(&dtoSubject.GetSubjectRequest{
		Id: params.SubjectId,
	})
	if err != nil {
		return nil, constants.ErrSubjectDBNotFound
	}
	_, err = u.departmentsUC.GetDepartment(&dtoDepartment.GetDepartmentRequest{
		Id: params.DepartmentId,
	})
	if err != nil {
		return nil, constants.ErrDepartmentDBNotFound
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
	_, err = u.departmentsUC.GetDepartment(&dtoDepartment.GetDepartmentRequest{
		Id: params.DepartmentId,
	})
	if err != nil {
		return nil, constants.ErrDepartmentDBNotFound
	}
	return &departmentSubjects.DeleteDepartmentSubjectsResponse{}, u.repoDepartmentSubjects.AddDepartmentSubjects(params.DepartmentId, params.SubjectId)
}
func (u DepartmentSubjectsUseCase) GetDepartmentSubjects(params *departmentSubjects.GetDepartmentSubjectsRequest) (*departmentSubjects.GetDepartmentSubjectsResponse, error) {
	res, length, err := u.repoDepartmentSubjects.GetDepartmentSubjects(params.DepartmentId, params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}
	return &departmentSubjects.GetDepartmentSubjectsResponse{
		DepartmentSubjects: res,
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
