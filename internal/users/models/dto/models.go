package dto

// User Only for Responses
type User struct {
	Id         string `json:"id"`
	Firstname  string `json:"firstname"`
	Surname    string `json:"surname"`
	Middlename string `json:"middlename,omitempty"`
	Age        int    `json:"age"`
	Sex        string `json:"sex"`
	Image      string `json:"image"`
	Role       string `json:"role"`
}

type BasicResponse struct{}

type GetUserRequest struct {
	Id string `path:"user_id"`
}

type GetUsersRequest struct {
	Limit  int64 `query:"limit"`
	Offset int64 `query:"offset"`
}

type CreateUserRequest struct {
	Firstname  string `json:"firstname"`
	Surname    string `json:"surname"`
	Middlename string `json:"middlename,omitempty"`
	BirthDate  string `json:"birth_date"`
	Sex        string `json:"sex"`
	Image      string `json:"image"`
	Role       string `json:"role"`
}

type UpdateUserRequest struct {
	Id         string `path:"user_id"`
	Firstname  string `json:"firstname"`
	Surname    string `json:"surname"`
	Middlename string `json:"middlename,omitempty"`
	BirthDate  string `json:"birth_date"`
	Sex        string `json:"sex"`
	Image      string `json:"image"`
	Role       string `json:"role"`
}

type DeleteUserRequest struct {
	Id string `path:"user_id"`
}

type CreateUserResponse struct {
	User
}

type UpdateUserResponse struct {
	User
}

type DeleteUserResponse struct {
	BasicResponse
}

type GetUserResponse struct {
	User
}

type GetUsersResponse struct {
	Users  []User `json:"users"`
	Length int64  `json:"length"`
}
