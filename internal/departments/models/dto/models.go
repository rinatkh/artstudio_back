package dto

import "github.com/rinatkh/artstudio_back/internal/subjects/models/dto"

// Department Only for Responses
type Department struct {
	Id             int64         `json:"id"`
	Name           string        `json:"name"`
	Description    string        `json:"description,omitempty"`
	Image          string        `json:"image"`
	Subject        []dto.Subject `json:"subjects"`
	LengthSubjects int64         `json:"length_subjects"`
}

type BasicResponse struct{}

type GetDepartmentRequest struct {
	Id             int64 `path:"department_id"`
	LimitSubjects  int64 `query:"limit_subjects"`
	OffsetSubjects int64 `query:"offset_subjects"`
}

type GetDepartmentsRequest struct {
	Limit          int64 `query:"limit"`
	Offset         int64 `query:"offset"`
	LimitSubjects  int64 `query:"limit_subjects"`
	OffsetSubjects int64 `query:"offset_subjects"`
}

type CreateDepartmentRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Image       string  `json:"image"`
	SubjectIDs  []int64 `json:"subjects_ids,omitempty"`
}

type UpdateDepartmentRequest struct {
	Id          int64  `path:"department_id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image"`
}

type DeleteDepartmentRequest struct {
	Id int64 `path:"department_id"`
}

type CreateDepartmentResponse struct {
	Department
}

type UpdateDepartmentResponse struct {
	Department
}

type DeleteDepartmentResponse struct {
	BasicResponse
}

type GetDepartmentResponse struct {
	Department
}

type GetDepartmentsResponse struct {
	Departments []Department `json:"departments"`
	Length      int64        `json:"length"`
}
