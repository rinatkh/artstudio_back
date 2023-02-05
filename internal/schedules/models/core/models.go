package core

type Schedule struct {
	Id          int64  `db:"id"`
	SubjectId   int64  `db:"subject_id"`
	StudentId   string `db:"student_id"`
	CabinetId   int64  `db:"cabinet_id"`
	Description string `db:"description"`
	IsPaid      bool   `db:"is_paid"`
	IsConfirmed bool   `db:"is_confirmed"`
	IsFinished  bool   `db:"is_finished"`
	WhenTime    int64  `db:"when_time"`
}
