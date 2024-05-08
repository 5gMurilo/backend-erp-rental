package service

import (
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"america-rental-backend/internal/core/util"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

type StorageService struct {
	config domain.StorageAuthentication
	repo   ports.StorageRepository
}

type JsonGraphSdkResponse struct {
	OdataContext       string `json:"@odata.context"`
	OdataNextLink      string `json:"@odata.nextLink"`
	MicrosoftGraphTips string `json:"@microsoft.graph.tips"`
	Value              []struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		WebURL string `json:"webUrl"`
		Folder struct {
			ChildCount int `json:"childCount"`
		} `json:"folder,omitempty"`
	} `json:"value"`
}

var client *msgraphsdkgo.GraphServiceClient
var token azcore.AccessToken
var driveId string
var httpClient = http.Client{}

func NewStorageService(config domain.StorageAuthentication, repo ports.StorageRepository) ports.StorageServices {
	return &StorageService{
		config,
		repo,
	}
}

func (s StorageService) GetDriveItemId(employeeName string) (string, error) {
	if client == nil {
		err := s.InitializeGraph()
		if err != nil {
			return "", err
		}
	}

	req, err := http.NewRequest("GET", "https://graph.microsoft.com/v1.0/me/drive/root/search(q='"+employeeName+"')?select=name,id,webUrl,folder", nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Bearer "+token.Token)

	res, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var j JsonGraphSdkResponse
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(respBody, &j)
	if err != nil {
		return "", err
	}

	var driveItemId string
	for _, v := range j.Value {
		if v.Folder.ChildCount == 0 && strings.Contains(v.WebURL, "/sistema/") {
			driveItemId = v.ID
			break
		}
	}

	return driveItemId, nil
}

func (s StorageService) InitializeGraph() error {
	credentials, err := azidentity.NewUsernamePasswordCredential(s.config.ObjectId, s.config.ClientId, s.config.Username, s.config.Password, nil)
	if err != nil {
		return err
	}

	token, err = credentials.GetToken(context.TODO(), policy.TokenRequestOptions{
		Claims:    "",
		EnableCAE: false,
		Scopes:    s.config.Scopes,
		TenantID:  s.config.ObjectId,
	})
	if err != nil {
		return err
	}

	client, err = msgraphsdkgo.NewGraphServiceClientWithCredentials(credentials, s.config.Scopes)
	if err != nil {
		return err
	}

	me, err := client.Me().Get(context.Background(), nil)
	if err != nil {
		return err
	}

	driveId = *me.GetId()

	return nil
}

func (s StorageService) SendFile(multipartFile *multipart.FileHeader, employeeName string, filetype string, actor string) (*domain.OnedriveFile, error) {
	var fileBytes []byte
	if client == nil {
		err := s.InitializeGraph()
		if err != nil {
			return nil, err
		}
	}

	file, err := multipartFile.Open()
	if err != nil {
		return nil, err
	}

	_, err = file.Read(fileBytes)
	if err != nil {
		return nil, err
	}

	url := util.NewPathFile(employeeName, filetype, multipartFile.Filename)

	httpClient := &http.Client{}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	defer func(writer *multipart.Writer) {
		err := writer.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(writer)

	part, err := writer.CreateFormFile("file", multipartFile.Filename)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", url, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+token.Token)

	rst, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	var response domain.OnedriveFileActionResponse
	rstBody, err := io.ReadAll(rst.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(rstBody, &response)
	if err != nil {
		return nil, err
	}

	newOnedriveFile := domain.OnedriveFile{
		Filename:    response.Name,
		FileUrl:     response.WebURL,
		DriveItemId: response.ID,
		Employee:    employeeName,
		Type:        filetype,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		UpdatedBy:   actor,
	}

	information, err := s.repo.RegisterUpdateInformation(context.TODO(), newOnedriveFile, actor)
	if err != nil {
		return nil, err
	}

	return information, nil
}

func (s StorageService) ListFiles(employeeName string) (*[]domain.OnedriveFile, error) {
	result, err := s.repo.GetOnedriveFilesByEmployee(context.TODO(), employeeName)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s StorageService) DeleteFile(driveItemId string) error {
	if client == nil {
		err := s.InitializeGraph()
		if err != nil {
			return err
		}
	}

	err := client.Drives().ByDriveId(driveId).Items().ByDriveItemId(driveItemId).Delete(context.Background(), nil)
	if err != nil {
		return err
	}

	err = s.repo.DeleteOnedriveFile(context.TODO(), driveItemId)
	if err != nil {
		return err
	}

	return nil
}
