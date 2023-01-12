package core

import "time"

type TimeTeacher struct {
	Id         int64     `db:"id"`
	TeacherId  string    `db:"teacher_id"`
	StartTime  time.Time `db:"start_time"`
	FinishTime time.Time `db:"finish_time"`
}
