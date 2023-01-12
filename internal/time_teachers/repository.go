package timeTeacher

import (
	"github.com/rinatkh/artstudio_back/internal/time_teachers/models/core"
	"time"
)

type TimeTeacherRepository interface {
	GetTimeTeacherById(id int64) (*core.TimeTeacher, error)
	GetTimeTeachers(teacherId string, startTime, finishTime time.Time) (*[]core.TimeTeacher, error)
	DeleteTimeTeacher(id int64) error
	CreateTimeTeacher(timeTeacher *core.TimeTeacher) (*core.TimeTeacher, error)
	UpdateTimeTeacher(timeTeacher *core.TimeTeacher) (*core.TimeTeacher, error)
}
