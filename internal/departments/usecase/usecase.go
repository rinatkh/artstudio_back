package usecase

import (
	"github.com/rinatkh/artstudio_back/config"
	departmentSubjects "github.com/rinatkh/artstudio_back/internal/DepartmentSubjects"
	"github.com/rinatkh/artstudio_back/internal/departments"
	"github.com/rinatkh/artstudio_back/internal/departments/models/convert"
	"github.com/rinatkh/artstudio_back/internal/departments/models/core"
	"github.com/rinatkh/artstudio_back/internal/departments/models/dto"
	"github.com/rinatkh/artstudio_back/internal/subjects"
	dtoSubject "github.com/rinatkh/artstudio_back/internal/subjects/models/dto"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/sirupsen/logrus"
)

type DepartmentUseCase struct {
	cfg                  *config.Config
	log                  *logrus.Entry
	repoDepartment       departments.DepartmentRepository
	subjectUC            subjects.UseCase
	departmentSubjectsUC departmentSubjects.UseCase
}

func NewDepartmentUC(cfg *config.Config, log *logrus.Entry, repoDepartment departments.DepartmentRepository, subjectUC subjects.UseCase, departmentSubjectsUC departmentSubjects.UseCase) departments.UseCase {
	return &DepartmentUseCase{
		cfg:                  cfg,
		log:                  log,
		repoDepartment:       repoDepartment,
		subjectUC:            subjectUC,
		departmentSubjectsUC: departmentSubjectsUC,
	}
}

func (u DepartmentUseCase) getSubjects(subjectIDs *[]departmentSubjects.DepartmentSubjects) (*[]dtoSubject.Subject, error) {
	var elems []dtoSubject.Subject
	for _, i := range *subjectIDs {
		subject, err := u.subjectUC.GetSubject(&dtoSubject.GetSubjectRequest{Id: i.SubjectId})
		if err != nil {
			return nil, err
		}
		elems = append(elems, subject.Subject)
	}
	return &elems, nil
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
	var elems []dtoSubject.Subject
	for _, i := range params.SubjectIDs {
		subject, err := u.subjectUC.GetSubject(&dtoSubject.GetSubjectRequest{
			Id: i,
		})
		if err != nil {
			return nil, constants.ErrSubjectDBNotFound
		}
		_, err = u.departmentSubjectsUC.AddDepartmentSubjects(&departmentSubjects.AddDepartmentSubjectsRequest{
			DepartmentId: result.Id,
			SubjectId:    i,
		})
		if err != nil {
			return nil, err
		}
		elems = append(elems, subject.Subject)
	}
	return &dto.CreateDepartmentResponse{Department: convert.Department2DTO(result, &elems, int64(len(elems)))}, nil
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
		Id:          params.Id,
		Name:        params.Name,
		Description: &params.Description,
		Image:       params.Image,
	}
	result, err := u.repoDepartment.UpdateDepartment(&department)
	if err != nil {
		return nil, err
	}
	elems, err := u.departmentSubjectsUC.GetDepartmentSubjects(&departmentSubjects.GetDepartmentSubjectsRequest{
		DepartmentId: result.Id,
		Limit:        0,
		Offset:       0,
	})
	list, err := u.getSubjects(elems.DepartmentSubjects)
	if err != nil {
		return nil, err
	}
	return &dto.UpdateDepartmentResponse{Department: convert.Department2DTO(result, list, elems.Length)}, nil
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
	_, err = u.departmentSubjectsUC.DeleteAll(&departmentSubjects.DeleteAllRequest{DepartmentId: params.Id})
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
	elems, err := u.departmentSubjectsUC.GetDepartmentSubjects(&departmentSubjects.GetDepartmentSubjectsRequest{
		DepartmentId: result.Id,
		Limit:        params.LimitSubjects,
		Offset:       params.OffsetSubjects,
	})
	list, err := u.getSubjects(elems.DepartmentSubjects)
	if err != nil {
		return nil, err
	}
	return &dto.GetDepartmentResponse{Department: convert.Department2DTO(result, list, elems.Length)}, nil
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
		elems, err := u.departmentSubjectsUC.GetDepartmentSubjects(&departmentSubjects.GetDepartmentSubjectsRequest{
			DepartmentId: i.Id,
			Limit:        params.LimitSubjects,
			Offset:       params.OffsetSubjects,
		})
		listSubjects, err := u.getSubjects(elems.DepartmentSubjects)
		if err != nil {
			return nil, err
		}
		result = append(result, convert.Department2DTO(&i, listSubjects, elems.Length))
	}
	return &dto.GetDepartmentsResponse{Departments: result, Length: length}, nil
}
