package dto

import (
	dtoSubject "github.com/rinatkh/artstudio_back/internal/subjects/models/dto"
	"github.com/rinatkh/artstudio_back/internal/users/models/dto"
)

// Schedule Only for Responses
type Schedule struct {
	Subject     dtoSubject.Subject `json:"subject"`
	WhenTime    int64              `json:"when_time"`
	Student     dto.User           `json:"student,omitempty"`
	IsPaid      bool               `json:"is_paid"`
	Description string             `json:"description"`
	IsConfirmed bool               `json:"is_confirmed"`
	Cabinet     string             `json:"cabinets"`
}

type BasicResponse struct{}

type GetScheduleRequest struct {
	ScheduleId int64  `path:"schedule_id"`
	UserId     string `header:"User-Id"`
}

type GetScheduleResponse struct {
	Schedule
}

type GetTeacherSchedulesRequest struct {
	UserId     string `header:"User-Id"`
	TeacherId  string `query:"teacher_id"`
	StartTime  int64  `query:"start_time"`
	FinishTime int64  `query:"finish_time"`
}

type GetTeacherSchedulesResponse struct {
	Schedules []Schedule `json:"schedules"`
}

type GetStudentSchedulesRequest struct {
	UserId     string `header:"User-Id"`
	StudentId  string `query:"student_id"`
	StartTime  int64  `query:"start_time"`
	FinishTime int64  `query:"finish_time"`
}

type GetStudentSchedulesResponse struct {
	Schedules []Schedule `json:"schedules"`
}

type CreateScheduleRequest struct {
	UserId      string `header:"User-Id"`
	StudentId   string `json:"student_id"`
	SubjectId   int64  `json:"subject_id"`
	Description string `json:"description"`
	WhenTime    int64  `json:"when_time"`
}

type CreateScheduleResponse struct {
	Schedule
}

type UpdateScheduleRequest struct {
	UserId      string `header:"User-Id"`
	ScheduleId  int64  `json:"subject_id"`
	CabinetId   int64  `json:"cabinet_id"`
	WhenTime    int64  `json:"when_time"`
	Description string `json:"description"`
	StudentId   string `json:"student_id"`
	IsPaid      bool   `json:"is_paid"`
	IsConfirmed bool   `json:"is_confirmed"`
	IsFinished  bool   `json:"is_finished"`
}

type UpdateScheduleResponse struct {
	Schedule
}

type DeleteScheduleRequest struct {
	ScheduleId int64  `path:"schedule_id"`
	UserId     string `header:"User-Id"`
}

type DeleteScheduleResponse struct {
	BasicResponse
}

type GetSlotsRequest struct {
	UserId     string `header:"User-Id"`
	SubjectId  int64  `query:"subject_id"`
	StartTime  int64  `query:"start_time"`
	FinishTime int64  `query:"finish_time"`
}

type GetSlotsResponse struct {
	WhenTime []int64 `json:"when_time"`
}
