package dto

import "github.com/rinatkh/artstudio_back/internal/users/models/dto"

// Subject Only for Responses
type Subject struct {
	Id          int64    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Image       string   `json:"image"`
	Teacher     dto.User `json:"teacher"`
}

type BasicResponse struct{}

type GetSubjectRequest struct {
	Id int64 `path:"subject_id"`
}

type GetSubjectsRequest struct {
	Limit  int64 `query:"limit"`
	Offset int64 `query:"offset"`
}

type CreateSubjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image"`
	TeacherId   string `json:"teacher_id"`
}

type UpdateSubjectRequest struct {
	Id          int64  `path:"subject_id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image"`
	TeacherId   string `json:"teacher_id"`
}

type DeleteSubjectRequest struct {
	Id int64 `path:"subject_id"`
}

type CreateSubjectResponse struct {
	Subject
}

type UpdateSubjectResponse struct {
	Subject
}

type DeleteSubjectResponse struct {
	BasicResponse
}

type GetSubjectResponse struct {
	Subject
}

type GetSubjectsResponse struct {
	Subjects []Subject `json:"subjects"`
	Length   int64     `json:"length"`
}
