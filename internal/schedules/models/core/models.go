package core

import "time"

type Schedule struct {
	SubjectId   string    `db:"subject_id"`
	StudentId   string    `db:"student_id"`
	IsPaid      bool      `db:"is_paid"`
	IsConfirmed bool      `db:"is_confirmed"`
	IsFinished  bool      `db:"is_finished"`
	WhenTime    time.Time `db:"when_time"`
}
