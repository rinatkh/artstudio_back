package departmentSubjects

type DepartmentSubjects struct {
	DepartmentId int64 `db:"department_id"`
	SubjectId    int64 `db:"subject_id"`
}

type BasicResponse struct{}

type AddDepartmentSubjectsRequest struct {
	DepartmentId int64
	SubjectId    int64
}
type AddDepartmentSubjectsResponse struct {
	BasicResponse
}
type DeleteDepartmentSubjectsRequest struct {
	DepartmentId int64
	SubjectId    int64
}
type DeleteDepartmentSubjectsResponse struct {
	BasicResponse
}
type GetDepartmentSubjectsRequest struct {
	DepartmentId int64
	Limit        int64
	Offset       int64
}
type GetDepartmentSubjectsResponse struct {
	DepartmentSubjects *[]DepartmentSubjects
	Length             int64
}

type DeleteAllRequest struct {
	DepartmentId int64
}
type DeleteAllResponse struct {
	BasicResponse
}
