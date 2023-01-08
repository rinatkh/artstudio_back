package departmentSubjects

type DepartmentSubjectsRepository interface {
	GetDepartmentSubjects(departmentId, limit, offset int64) (*[]DepartmentSubjects, int64, error)
	DeleteDepartmentSubjects(departmentId, subjectId int64) error
	AddDepartmentSubjects(departmentId, subjectId int64) error
	DeleteAll(departmentId int64) error
}
