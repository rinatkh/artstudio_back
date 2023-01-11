package usecase

import (
	"github.com/rinatkh/artstudio_back/config"
	"github.com/rinatkh/artstudio_back/internal/static"
	"github.com/rinatkh/artstudio_back/pkg/utils"
	"github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type StaticUseCase struct {
	cfg *config.Config
	log *logrus.Entry
}

func NewStaticUC(cfg *config.Config, log *logrus.Entry) static.UseCase {
	return &StaticUseCase{
		cfg: cfg,
		log: log,
	}
}

func (u StaticUseCase) UploadFile(fileHeader *multipart.FileHeader) (*static.UploadFileResponse, error) {
	src, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	uuid, err := utils.GenUUID()
	if err != nil {
		return nil, err
	}

	filename := uuid + filepath.Ext(fileHeader.Filename)

	dst, err := os.Create("/opt/files/" + filename)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}

	url := "/" + filename
	return &static.UploadFileResponse{URL: url}, nil
}

func (u StaticUseCase) UploadImage(fileHeader *multipart.FileHeader) (*static.UploadImageResponse, error) {
	src, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	uuid, err := utils.GenUUID()
	if err != nil {
		return nil, err
	}

	filename := uuid + filepath.Ext(fileHeader.Filename)

	dst, err := os.Create("/opt/pics/" + filename)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}

	url := "/" + filename
	return &static.UploadImageResponse{URL: url}, nil
}
