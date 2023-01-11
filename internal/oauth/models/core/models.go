package core

type Oauth struct {
	Id    string `db:"uuid"`
	Email string `db:"email"`
}
