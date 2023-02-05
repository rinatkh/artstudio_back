package core

type User struct {
	Id         string `db:"uuid"`
	Firstname  string `db:"firstname"`
	Surname    string `db:"surname"`
	Middlename string `db:"middlename"`
	Sex        string `db:"sex"`
	BirthDate  int64  `db:"birth_date"`
	Role       string `db:"role"`
	Image      string `db:"image"`
	CreateAt   int64  `db:"created_at"`
}
