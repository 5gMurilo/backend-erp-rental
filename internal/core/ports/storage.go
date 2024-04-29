package ports

import (
	"america-rental-backend/internal/core/domain"
	"mime/multipart"
)

type StorageServices interface {
	InitializeGraph() error
	SendFile(multipartFile *multipart.FileHeader, employeeName string) (bool, error)
	ListFiles(employeeName string) ([]domain.StorageResponse, error)
	GetDriveItemId(employeeName string) (string, error)
}
