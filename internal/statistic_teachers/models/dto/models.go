package dto

import "github.com/rinatkh/artstudio_back/internal/users/models/dto"

// StatisticTeacher Only for Responses
type StatisticTeacher struct {
	FutureLessons int64 `json:"future_lessons"`
	NeedDzLessons int64 `json:"need_dz_lessons"`
	PastLessons   int64 `json:"past_lessons"`
}

type GetStatisticTeacherRequest struct {
	UserId    string `header:"User-Id"`
	TeacherId string `path:"teacher_id"`
}

type GetStatisticTeacherResponse struct {
	StatisticTeacher
	Teacher dto.User `json:"teacher"`
}

type CreateStatisticTeacherRequest struct {
	TeacherId string `json:"teacher_id"`
}

type CreateStatisticTeacherResponse struct {
	StatisticTeacher
}

type DeleteStatisticTeacherRequest struct {
	TeacherId string `json:"teacher_id"`
}

type DeleteStatisticTeacherResponse struct {
}

type UpdateStatisticTeacherRequest struct {
	TeacherId string `json:"teacher_id"`
	StatisticTeacher
}

type UpdateStatisticTeacherResponse struct {
	StatisticTeacher
	Teacher dto.User `json:"teacher"`
}
