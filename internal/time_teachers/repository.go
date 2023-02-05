package timeTeacher

import (
	"github.com/rinatkh/artstudio_back/internal/time_teachers/models/core"
)

type TimeTeacherRepository interface {
	GetTimeTeacherById(id int64) (*core.TimeTeacher, error)
	GetTimeTeachers(teacherId string, startTime, finishTime int64) (*[]core.TimeTeacher, error)
	DeleteTimeTeacher(id int64) error
	DeleteTimeTeacherByUserId(id string) error
	CreateTimeTeacher(timeTeacher *core.TimeTeacher) (*core.TimeTeacher, error)
	UpdateTimeTeacher(timeTeacher *core.TimeTeacher) (*core.TimeTeacher, error)
}
