package ports

import (
	"america-rental-backend/internal/core/domain"
	"context"
	"mime/multipart"
)

type StorageServices interface {
	InitializeGraph() error
	SendFile(multipartFile *multipart.FileHeader, employeeName string, filetype string, actor string) (*domain.OnedriveFile, error)
	ListFiles(employeeName string) (*[]domain.OnedriveFile, error)
	GetDriveItemId(employeeName string) (string, error)
	DeleteFile(driveItemId string) error
}

type StorageRepository interface {
	RegisterUpdateInformation(ctx context.Context, onedriveFile domain.OnedriveFile, actor string) (*domain.OnedriveFile, error)
	GetOnedriveFilesByEmployee(ctx context.Context, employeeName string) (*[]domain.OnedriveFile, error)
	UpdateOnedriveFile(ctx context.Context, file domain.OnedriveFile, actor string) (*domain.OnedriveFile, error)
	DeleteOnedriveFile(ctx context.Context, driveItemid string) error
}
