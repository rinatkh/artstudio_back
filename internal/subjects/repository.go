package subjects

import "github.com/rinatkh/artstudio_back/internal/subjects/models/core"

type SubjectRepository interface {
	GetSubjectById(id int64) (*core.Subject, error)
	GetSubjects(limit, offset int64) (*[]core.Subject, int64, error)
	DeleteSubject(id int64) error
	CreateSubject(subject *core.Subject) (*core.Subject, error)
	UpdateSubject(subject *core.Subject) (*core.Subject, error)
}
