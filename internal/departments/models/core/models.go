package core

type Department struct {
	Id          int64   `db:"id"`
	Name        string  `db:"name"`
	Description *string `db:"description"`
	Image       string  `db:"image"`
}
