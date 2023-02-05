package core

type StatisticStudent struct {
	Id                 int64  `db:"id"`
	StudentId          string `db:"student_id"`
	BalanceLessons     int64  `db:"balance_lessons"`
	DoneLessons        int64  `db:"done_lessons"`
	NeedPaymentLessons int64  `db:"need_payment_lessons"`
}
