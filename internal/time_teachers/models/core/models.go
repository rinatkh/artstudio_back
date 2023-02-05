package core

type TimeTeacher struct {
	Id         int64  `db:"id"`
	TeacherId  string `db:"teacher_id"`
	StartTime  int64  `db:"start_time"`
	FinishTime int64  `db:"finish_time"`
}
