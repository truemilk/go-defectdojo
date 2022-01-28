package defectdojo

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

type ImportScanService struct {
	client *Client
}

type ImportScan struct {
	ScanDate             *string   `json:"scan_date,omitempty"`
	MinimumSeverity      *string   `json:"minimum_severity,omitempty"`
	Active               *bool     `json:"active,omitempty"`
	Verified             *bool     `json:"verified,omitempty"`
	ScanType             *string   `json:"scan_type,omitempty"`
	EndpointToAdd        *int      `json:"endpoint_to_add,omitempty"`
	File                 *string   `json:"file,omitempty"`
	ProductTypeName      *string   `json:"product_type_name,omitempty"`
	ProductName          *string   `json:"product_name,omitempty"`
	EngagementName       *string   `json:"engagement_name,omitempty"`
	Engagement           *int      `json:"engagement,omitempty"`
	TestTitle            *string   `json:"test_title,omitempty"`
	AutoCreateContext    *bool     `json:"auto_create_context,omitempty"`
	Lead                 *int      `json:"lead,omitempty"`
	Tags                 []*string `json:"tags,omitempty"`
	CloseOldFindings     *bool     `json:"close_old_findings,omitempty"`
	PushToJira           *bool     `json:"push_to_jira,omitempty"`
	Environment          *string   `json:"environment,omitempty"`
	Version              *string   `json:"version,omitempty"`
	BuildId              *string   `json:"build_id,omitempty"`
	BranchTag            *string   `json:"branch_tag,omitempty"`
	CommitHash           *string   `json:"commit_hash,omitempty"`
	ApiScanConfiguration *int      `json:"api_scan_configuration,omitempty"`
	Service              *string   `json:"service,omitempty"`
	GroupBy              *string   `json:"group_by,omitempty"`
	Test                 *int      `json:"test,omitempty"`
	TestId               *int      `json:"test_id,omitempty"`
	EngagementId         *int      `json:"engagement_id,omitempty"`
	ProductId            *int      `json:"product_id,omitempty"`
	ProductTypeId        *int      `json:"product_type_id,omitempty"`
}

type ImportScanMap map[string]string

func (c *ImportScanService) Create(ctx context.Context, m *ImportScanMap) (*ImportScan, error) {
	path := fmt.Sprintf("%s/import-scan/", c.client.BaseURL)

	req, err := newFileUploadRequest(path, *m)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(ImportScan)
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func newFileUploadRequest(uri string, params ImportScanMap) (*http.Request, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	for key, val := range params {
		if key == "file" {
			file, err := os.Open(val)
			if err != nil {
				return nil, err
			}
			fileContents, err := ioutil.ReadAll(file)
			if err != nil {
				return nil, err
			}
			fi, err := file.Stat()
			if err != nil {
				return nil, err
			}
			err = file.Close()
			if err != nil {
				return nil, err
			}
			part, err := writer.CreateFormFile(key, fi.Name())
			if err != nil {
				return nil, err
			}
			_, err = part.Write(fileContents)
			if err != nil {
				return nil, err
			}
		} else {
			_ = writer.WriteField(key, val)
		}
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", writer.FormDataContentType())

	return r, nil
}
