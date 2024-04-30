package domain

import "time"

type OnedriveFileActionResponse struct {
	OdataContext              string          `json:"@odata.context"`
	MicrosoftGraphDownloadURL string          `json:"@microsoft.graph.downloadUrl"`
	CreatedDateTime           time.Time       `json:"createdDateTime"`
	ETag                      string          `json:"eTag"`
	ID                        string          `json:"id"`
	LastModifiedDateTime      time.Time       `json:"lastModifiedDateTime"`
	Name                      string          `json:"name"`
	WebURL                    string          `json:"webUrl"`
	CTag                      string          `json:"cTag"`
	Size                      int64           `json:"size"`
	CreatedBy                 EdBy            `json:"createdBy"`
	LastModifiedBy            EdBy            `json:"lastModifiedBy"`
	ParentReference           ParentReference `json:"parentReference"`
	File                      File            `json:"file"`
	FileSystemInfo            FileSystemInfo  `json:"fileSystemInfo"`
	Shared                    Shared          `json:"shared"`
}

type EdBy struct {
	Application Application `json:"application"`
	User        Application `json:"user"`
}

type Application struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type File struct {
	MIMEType string `json:"mimeType"`
	Hashes   Hashes `json:"hashes"`
}

type Hashes struct {
	QuickXorHash string `json:"quickXorHash"`
}

type FileSystemInfo struct {
	CreatedDateTime      time.Time `json:"createdDateTime"`
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
}

type ParentReference struct {
	DriveType string `json:"driveType"`
	DriveID   string `json:"driveId"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	SiteID    string `json:"siteId"`
}

type Shared struct {
	Scope string `json:"scope"`
}
