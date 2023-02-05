package dto

import "github.com/rinatkh/artstudio_back/internal/users/models/dto"

// StatisticStudent Only for Responses
type StatisticStudent struct {
	BalanceLessons     int64 `json:"balance_lessons"`
	DoneLessons        int64 `json:"done_lessons"`
	NeedPaymentLessons int64 `json:"need_payment_lessons"`
}

type GetStatisticStudentRequest struct {
	UserId    string `header:"User-Id"`
	StudentId string `path:"student_id"`
}

type GetStatisticStudentResponse struct {
	StatisticStudent
	Student dto.User `json:"student"`
}

type CreateStatisticStudentRequest struct {
	StudentId string `json:"student_id"`
}

type CreateStatisticStudentResponse struct {
	StatisticStudent
}

type DeleteStatisticStudentRequest struct {
	StudentId string `json:"student_id"`
}

type DeleteStatisticStudentResponse struct {
}

type UpdateStatisticStudentRequest struct {
	StudentId string `json:"student_id"`
	StatisticStudent
}

type UpdateStatisticStudentResponse struct {
	StatisticStudent
	Student dto.User `json:"student"`
}
