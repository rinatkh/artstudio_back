package departmentSubjects

type UseCase interface {
	AddDepartmentSubjects(params *AddDepartmentSubjectsRequest) (*AddDepartmentSubjectsResponse, error)
	DeleteDepartmentSubjects(params *DeleteDepartmentSubjectsRequest) (*DeleteDepartmentSubjectsResponse, error)
	GetDepartmentSubjects(params *GetDepartmentSubjectsRequest) (*GetDepartmentSubjectsResponse, error)
	DeleteAll(params *DeleteAllRequest) (*DeleteAllResponse, error)
}
