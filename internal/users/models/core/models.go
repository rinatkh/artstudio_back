package core

import "time"

type User struct {
	Id         string    `db:"uuid"`
	Firstname  string    `db:"firstname"`
	Surname    string    `db:"surname"`
	Middlename string    `db:"middlename"`
	Sex        string    `db:"sex"`
	BirthDate  time.Time `db:"birth_date"`
	Role       string    `db:"role"`
	Image      string    `db:"image"`
	CreateAt   time.Time `db:"created_at"`
}
