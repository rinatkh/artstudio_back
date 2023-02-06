package dto

// Cabinet Only for Responses
type Cabinet struct {
	CabinetId int64  `json:"cabinet_id"`
	Name      string `json:"name"`
	Time      []Time `json:"time"`
}

type Time struct {
	Id         int64 `json:"id"`
	StartTime  int64 `json:"start_time"`
	FinishTime int64 `json:"finish_time"`
}

type TimeRequest struct {
	StartTime  int64 `json:"start_time"`
	FinishTime int64 `json:"finish_time"`
}

type GetCabinetRequest struct {
	CabinetId int64 `path:"cabinet_id"`
}

type GetCabinetResponse struct {
	Cabinet
}

type CreateCabinetRequest struct {
	Name string        `json:"name"`
	Time []TimeRequest `json:"time"`
}

type CreateCabinetResponse struct {
	Cabinet
}

type DeleteCabinetRequest struct {
	CabinetId int64 `path:"cabinet_id"`
}

type DeleteCabinetResponse struct {
}

type UpdateCabinetRequest struct {
	CabinetId int64  `path:"cabinet_id"`
	Name      string `json:"name"`
}

type UpdateCabinetResponse struct {
	Cabinet
}

type UpdateCabinetAddTimeRequest struct {
	CabinetId int64       `path:"cabinet_id"`
	Time      TimeRequest `json:"time"`
}

type UpdateCabinetDeleteTimeRequest struct {
	CabinetId     int64 `path:"cabinet_id"`
	CabinetTimeId int64 `json:"time_id"`
}
