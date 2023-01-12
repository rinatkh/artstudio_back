package core

type StatisticTeacher struct {
	Id            int64  `db:"id"`
	TeacherId     string `db:"teacher_id"`
	FutureLessons int64  `db:"future_lessons"`
	NeedDzLessons int64  `db:"need_dz_lessons"`
	PastLessons   int64  `db:"past_lessons"`
}
