package static

import (
	"mime/multipart"
)

type UseCase interface {
	UploadImage(fileHeader *multipart.FileHeader) (*UploadImageResponse, error)
	UploadFile(fileHeader *multipart.FileHeader) (*UploadFileResponse, error)
}
