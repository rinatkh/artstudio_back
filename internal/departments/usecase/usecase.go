package usecase

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/artstudio_back/config"
	departmentSubjects "github.com/rinatkh/artstudio_back/internal/DepartmentSubjects"
	consts "github.com/rinatkh/artstudio_back/internal/constants"
	"github.com/rinatkh/artstudio_back/internal/departments"
	"github.com/rinatkh/artstudio_back/internal/departments/models/convert"
	"github.com/rinatkh/artstudio_back/internal/departments/models/core"
	"github.com/rinatkh/artstudio_back/internal/departments/models/dto"
	"github.com/rinatkh/artstudio_back/internal/subjects"
	dtoSubject "github.com/rinatkh/artstudio_back/internal/subjects/models/dto"
	"github.com/rinatkh/artstudio_back/internal/users"
	dtoUser "github.com/rinatkh/artstudio_back/internal/users/models/dto"
	"github.com/rinatkh/artstudio_back/pkg/constants"
	"github.com/sirupsen/logrus"
)

type DepartmentUseCase struct {
	cfg                  *config.Config
	log                  *logrus.Entry
	repoDepartment       departments.DepartmentRepository
	subjectUC            subjects.UseCase
	departmentSubjectsUC departmentSubjects.UseCase
	userUC               users.UseCase
}

func NewDepartmentUC(cfg *config.Config, log *logrus.Entry, repoDepartment departments.DepartmentRepository, subjectUC subjects.UseCase, departmentSubjectsUC departmentSubjects.UseCase, userUC users.UseCase) departments.UseCase {
	return &DepartmentUseCase{
		cfg:                  cfg,
		log:                  log,
		repoDepartment:       repoDepartment,
		subjectUC:            subjectUC,
		departmentSubjectsUC: departmentSubjectsUC,
		userUC:               userUC,
	}
}
func (u DepartmentUseCase) checkAdmin(admin string) error {
	user, err := u.userUC.GetUser(&dtoUser.GetUserRequest{Id: admin})
	if err != nil {
		return constants.ErrNoPrivileges
	}
	if user.Role != consts.Admin {
		return constants.ErrNoPrivileges
	}
	return nil
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
	err := u.checkAdmin(params.UserId)
	if err != nil {
		return nil, err
	}
	var elems []dtoSubject.Subject
	for _, i := range params.SubjectIDs {
		subject, err := u.subjectUC.GetSubject(&dtoSubject.GetSubjectRequest{
			Id: i,
		})
		if err != nil {
			return nil, constants.NewCodedError(fmt.Sprintf("subject %d not found in the database", i), fiber.StatusBadRequest)
		}
		elems = append(elems, subject.Subject)
	}
	department := core.Department{
		Name:        params.Name,
		Description: params.Description,
		Image:       params.Image,
	}
	result, err := u.repoDepartment.CreateDepartment(&department)
	if err != nil {
		return nil, err
	}
	for _, i := range params.SubjectIDs {
		_, err = u.departmentSubjectsUC.AddDepartmentSubjects(&departmentSubjects.AddDepartmentSubjectsRequest{
			DepartmentId: result.Id,
			SubjectId:    i,
		})
		if err != nil {
			return nil, err
		}
	}
	return &dto.CreateDepartmentResponse{Department: convert.Department2DTO(result, &elems, int64(len(elems)))}, nil
}

func (u DepartmentUseCase) UpdateDepartment(params *dto.UpdateDepartmentRequest) (*dto.UpdateDepartmentResponse, error) {
	err := u.checkAdmin(params.UserId)
	if err != nil {
		return nil, err
	}
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
		Description: params.Description,
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
	list, err := u.getSubjects(&elems.DepartmentSubjects)
	if err != nil {
		return nil, err
	}
	return &dto.UpdateDepartmentResponse{Department: convert.Department2DTO(result, list, elems.Length)}, nil
}

func (u DepartmentUseCase) DeleteDepartment(params *dto.DeleteDepartmentRequest) (*dto.DeleteDepartmentResponse, error) {
	err := u.checkAdmin(params.UserId)
	if err != nil {
		return nil, err
	}
	check, err := u.repoDepartment.GetDepartmentById(params.Id)
	if err != nil {
		return nil, err
	}
	if check == nil {
		return nil, constants.ErrDepartmentDBNotFound
	}
	_, err = u.departmentSubjectsUC.DeleteAll(&departmentSubjects.DeleteAllRequest{DepartmentId: params.Id})
	if err != nil {
		return nil, err
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
	elems, err := u.departmentSubjectsUC.GetDepartmentSubjects(&departmentSubjects.GetDepartmentSubjectsRequest{
		DepartmentId: result.Id,
		Limit:        params.LimitSubjects,
		Offset:       params.OffsetSubjects,
	})
	list, err := u.getSubjects(&elems.DepartmentSubjects)
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
		if err != nil {
			return nil, err
		}
		if elems.DepartmentSubjects == nil {
			result = append(result, convert.Department2DTO(&i, &[]dtoSubject.Subject{}, 0))
			continue
		}
		listSubjects, err := u.getSubjects(&elems.DepartmentSubjects)
		if err != nil {
			return nil, err
		}
		result = append(result, convert.Department2DTO(&i, listSubjects, elems.Length))
	}
	return &dto.GetDepartmentsResponse{Departments: result, Length: length}, nil
}

func (u DepartmentUseCase) AddSubjDepartment(params *dto.UpdateDepartmentSubjectRequest) (*dto.UpdateDepartmentSubjectResponse, error) {
	err := u.checkAdmin(params.UserId)
	if err != nil {
		return nil, err
	}
	_, err = u.departmentSubjectsUC.AddDepartmentSubjects(&departmentSubjects.AddDepartmentSubjectsRequest{
		DepartmentId: params.Id,
		SubjectId:    params.SubjectId,
	})
	if err != nil {
		return nil, err
	}
	res, err := u.GetDepartment(&dto.GetDepartmentRequest{
		Id:             params.Id,
		LimitSubjects:  0,
		OffsetSubjects: 0,
	})
	if err != nil {
		return nil, err
	}
	return &dto.UpdateDepartmentSubjectResponse{Department: res.Department}, nil
}

func (u DepartmentUseCase) DelSubDepartment(params *dto.UpdateDepartmentSubjectRequest) (*dto.UpdateDepartmentSubjectResponse, error) {
	err := u.checkAdmin(params.UserId)
	if err != nil {
		return nil, err
	}
	_, err = u.departmentSubjectsUC.DeleteDepartmentSubjects(&departmentSubjects.DeleteDepartmentSubjectsRequest{
		DepartmentId: params.Id,
		SubjectId:    params.SubjectId,
	})
	if err != nil {
		return nil, err
	}
	res, err := u.GetDepartment(&dto.GetDepartmentRequest{
		Id:             params.Id,
		LimitSubjects:  0,
		OffsetSubjects: 0,
	})
	if err != nil {
		return nil, err
	}
	return &dto.UpdateDepartmentSubjectResponse{Department: res.Department}, nil
}
