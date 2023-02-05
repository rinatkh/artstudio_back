package schedules

import (
	"github.com/rinatkh/artstudio_back/internal/schedules/models/core"
)

type ScheduleRepository interface {
	GetScheduleById(id int64) (*core.Schedule, error)
	GetTeacherSchedules(teacherId string, startTime, finishTime int64) (*[]core.Schedule, error)
	GetStudentSchedules(studentId string, startTime, finishTime int64) (*[]core.Schedule, error)
	GetSubjectSchedules(subjectId int64, startTime, finishTime int64) (*[]core.Schedule, error)
	GetSchedules(startTime, finishTime int64) (*[]core.Schedule, error)
	DeleteSchedule(id int64) error
	DeleteScheduleBySubjectId(id int64) error
	DeleteScheduleByStudentId(id string) error
	CreateSchedule(schedule *core.Schedule) (*core.Schedule, error)
	UpdateSchedule(schedule *core.Schedule) (*core.Schedule, error)
}
