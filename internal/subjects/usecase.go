package subjects

import "github.com/rinatkh/artstudio_back/internal/subjects/models/dto"

type UseCase interface {
	CreateSubject(params *dto.CreateSubjectRequest) (*dto.CreateSubjectResponse, error)
	UpdateSubject(params *dto.UpdateSubjectRequest) (*dto.UpdateSubjectResponse, error)
	DeleteSubject(params *dto.DeleteSubjectRequest) (*dto.DeleteSubjectResponse, error)
	GetSubject(params *dto.GetSubjectRequest) (*dto.GetSubjectResponse, error)
	GetSubjects(params *dto.GetSubjectsRequest) (*dto.GetSubjectsResponse, error)
}
