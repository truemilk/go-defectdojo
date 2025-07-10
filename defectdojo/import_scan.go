package defectdojo

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"strings"
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
	DeduplicationOnEngagement  *bool     `json:"deduplication_on_engagement,omitempty"`
	Lead                 *int      `json:"lead,omitempty"`
	Tags                 *[]string `json:"tags,omitempty"`
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

type importScanMap map[string]string

func (c *ImportScanService) Create(ctx context.Context, m *ImportScan) (*ImportScan, error) {
	path := fmt.Sprintf("%s/import-scan/", c.client.BaseURL)

	up, err := structTagToMap(*m)
	if err != nil {
		return nil, err
	}
	req, err := newFileUploadRequest(path, &up)
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

func newFileUploadRequest(uri string, params *importScanMap) (*http.Request, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	for key, val := range *params {
		switch key {
		case "file":
			file, err := os.Open(val)
			if err != nil {
				return nil, err
			}
			fileContents, err := io.ReadAll(file)
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
		case "tags":
			t := strings.Trim(val, "[")
			t = strings.Trim(t, "]")
			_ = writer.WriteField(key, t)
		default:
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

func structTagToMap(in interface{}) (importScanMap, error) {
	m := make(importScanMap)

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {

		tag := strings.Split(t.Field(i).Tag.Get("json"), ",")[0]
		if len(tag) == 0 {
			return nil, errors.New("tag not found")
		}

		value := v.Field(i).Interface()
		if v.Field(i).IsZero() {
			continue
		}
		if v.Field(i).Kind() == reflect.Ptr {
			value = v.Field(i).Elem()
		}

		m[tag] = fmt.Sprintf("%v", value)
	}

	return m, nil
}
