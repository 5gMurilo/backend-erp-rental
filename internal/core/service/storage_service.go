package service

import (
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"context"
	"encoding/json"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdkgo "github.com/microsoftgraph/msgraph-sdk-go"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

type StorageService struct {
	config  domain.StorageAuthentication
	storage ports.StorageServices
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

func NewStorageService(storage ports.StorageServices, config domain.StorageAuthentication) ports.StorageServices {
	return &StorageService{
		config,
		storage,
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
		//TenantID:  "eada9b66-9e51-45a5-858e-ad99eddb9c48",
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

func (s StorageService) SendFile(multipartFile *multipart.FileHeader, employeeName string) (bool, error) {
	var fileBytes []byte
	if client == nil {
		err := s.InitializeGraph()
		if err != nil {
			return false, err
		}
	}

	file, err := multipartFile.Open()
	if err != nil {
		return false, err
	}

	_, err = file.Read(fileBytes)
	if err != nil {
		return false, err
	}

	driveItemId, err := s.GetDriveItemId(employeeName)
	if err != nil {
		return false, err
	}

	req := client.Drives().ByDriveId(driveId).Items().ByDriveItemId(driveItemId).Content()

	rst, err := req.Put(context.Background(), fileBytes, nil)
	if err != nil {
		return false, err
	}

	if rst.GetId() == nil {
		return false, errors.New("Falha ao enviar arquivo")
	}

	return true, nil
}

func (s StorageService) ListFiles(employeeName string) ([]domain.StorageResponse, error) {
	if client == nil {
		err := s.InitializeGraph()
		if err != nil {
			return nil, err
		}
	}

	driveItemId, err := s.GetDriveItemId(employeeName)
	if err != nil {
		return nil, err
	}

	children, err := client.Drives().ByDriveId(driveId).Items().ByDriveItemId(driveItemId).Children().Get(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	var rst []domain.StorageResponse
	for _, driveItem := range children.GetValue() {
		gsr := domain.StorageResponse{
			Id:     driveItem.GetId(),
			Name:   driveItem.GetName(),
			WebUrl: driveItem.GetWebUrl(),
		}
		rst = append(rst, gsr)
	}

	return rst, nil
}
