package defectdojo

import (
	"context"
	"fmt"
)

type ReImportScanService struct {
	client *Client
}

type ReImportScan struct {
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

type ReimportScanMap map[string]string

func (c *ReImportScanService) Create(ctx context.Context, m *ReImportScan) (*ReImportScan, error) {
	path := fmt.Sprintf("%s/reimport-scan/", c.client.BaseURL)

	up, err := structTagToMap(*m)
	if err != nil {
		return nil, err
	}
	req, err := newFileUploadRequest(path, &up)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(ReImportScan)
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
