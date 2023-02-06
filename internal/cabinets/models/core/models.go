package core

type Cabinet struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}

type CabinetTime struct {
	Id         int64 `db:"id"`
	SubjectId  int64 `db:"subject_id"`
	CabinetId  int64 `db:"cabinet_id"`
	StartTime  int64 `db:"start_time"`
	FinishTime int64 `db:"finish_time"`
}
