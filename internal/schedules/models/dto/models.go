package dto

import (
	dtoSubject "github.com/rinatkh/artstudio_back/internal/subjects/models/dto"
	"github.com/rinatkh/artstudio_back/internal/users/models/dto"
	"time"
)

// Schedule Only for Responses
type Schedule struct {
	Subject  dtoSubject.Subject `json:"subject"`
	WhenTime time.Time          `json:"when_time"`
	Student  dto.User           `json:"student,omitempty"`
	IsPaid   bool               `json:"is_paid"`
}

type BasicResponse struct{}

type GetScheduleRequest struct {
	Id string `path:"schedule_id"`
}

type GetSchedulesRequest struct {
	Limit  int64 `query:"limit"`
	Offset int64 `query:"offset"`
}

type CreateScheduleRequest struct {
	Firstname  string `json:"firstname"`
	Surname    string `json:"surname"`
	Middlename string `json:"middlename,omitempty"`
	BirthDate  string `json:"birth_date"`
	Sex        string `json:"sex"`
	Image      string `json:"image"`
	Role       string `json:"role"`
}

type UpdateScheduleRequest struct {
	Id         string `path:"schedule_id"`
	ScheduleId string `header:"User-Id"`
	Firstname  string `json:"firstname"`
	Surname    string `json:"surname"`
	Middlename string `json:"middlename,omitempty"`
	BirthDate  string `json:"birth_date"`
	Sex        string `json:"sex"`
	Image      string `json:"image"`
	Role       string `json:"role"`
}

type DeleteScheduleRequest struct {
	Id         string `path:"schedule_id"`
	ScheduleId string `header:"User-Id"`
}

type CreateScheduleResponse struct {
	Schedule
}

type UpdateScheduleResponse struct {
	Schedule
}

type DeleteScheduleResponse struct {
	BasicResponse
}

type GetScheduleResponse struct {
	Schedule
}

type GetSchedulesResponse struct {
	Schedules []Schedule `json:"Schedules"`
	Length    int64      `json:"length"`
}
