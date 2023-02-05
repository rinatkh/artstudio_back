package schedules

import "github.com/rinatkh/artstudio_back/internal/schedules/models/dto"

type UseCase interface {
	CreateSchedule(params *dto.CreateScheduleRequest) (*dto.CreateScheduleResponse, error)
	UpdateSchedule(params *dto.UpdateScheduleRequest) (*dto.UpdateScheduleResponse, error)
	DeleteSchedule(params *dto.DeleteScheduleRequest) (*dto.DeleteScheduleResponse, error)
	GetSchedule(params *dto.GetScheduleRequest) (*dto.GetScheduleResponse, error)
	GetStudentSchedules(params *dto.GetStudentSchedulesRequest) (*dto.GetStudentSchedulesResponse, error)
	GetTeacherSchedules(params *dto.GetTeacherSchedulesRequest) (*dto.GetTeacherSchedulesResponse, error)
	GetSlotsSchedules(params *dto.GetSlotsRequest) (*dto.GetSlotsResponse, error)
}
