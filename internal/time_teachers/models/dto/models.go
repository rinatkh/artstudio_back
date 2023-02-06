package dto

import (
	"github.com/rinatkh/artstudio_back/internal/users/models/dto"
)

// TimeTeacher Only for Responses
type TimeTeacher struct {
	Id         int64 `json:"id"`
	StartTime  int64 `json:"start_time"`
	FinishTime int64 `json:"finish_time"`
}

type AllTimeTeacher struct {
	Teacher dto.User      `json:"teacher"`
	Time    []TimeTeacher `json:"time"`
}
type OneTimeTeacher struct {
	TeacherId dto.User    `json:"teacher_id"`
	Time      TimeTeacher `json:"time"`
}

type BasicResponse struct{}

type GetTimeTeacherRequest struct {
	UserId string `header:"User-Id"`
	Id     int64  `path:"time_teacher_id"`
}

type GetTimeTeacherResponse struct {
	OneTimeTeacher
}

type GetTimeTeachersRequest struct {
	UserId     string `header:"User-Id"`
	TeacherId  string `query:"teacher_id"`
	StartTime  int64  `query:"start_time"`
	FinishTime int64  `query:"finish_time"`
}

type GetTimeTeachersResponse struct {
	AllTimeTeacher
}

type Time struct {
	StartTime  int64 `json:"start_time"`
	FinishTime int64 `json:"finish_time"`
}

type CreateTimeTeacherRequest struct {
	UserId    string `header:"User-Id"`
	Time      []Time `json:"time"`
	TeacherId string `json:"teacher_id,omitempty"`
}

type CreateTimeTeacherResponse struct {
	AllTimeTeacher
}
type UpdateTimeTeacherRequest struct {
	Id     int64  `path:"time_teacher_id"`
	UserId string `header:"User-Id"`
	Time   Time   `json:"time"`
}

type UpdateTimeTeacherResponse struct {
	OneTimeTeacher
}

type DeleteTimeTeacherRequest struct {
	Id     int64  `path:"time_teacher_id"`
	UserId string `header:"User-Id"`
}

type DeleteTimeTeacherResponse struct {
	BasicResponse
}
