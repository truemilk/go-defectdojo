package defectdojo

import (
	"context"
	"fmt"
)

type ReImportScanService struct {
	client *Client
}

type SeverityCount struct {
	Active       int `json:"active"`
	Verified     int `json:"verified"`
	Duplicate    int `json:"duplicate"`
	FalseP       int `json:"false_p"`
	OutOfScope   int `json:"out_of_scope"`
	IsMitigated  int `json:"is_mitigated"`
	RiskAccepted int `json:"risk_accepted"`
	Total        int `json:"total"`
}

type SeverityBreakdown struct {
	Info     SeverityCount `json:"info"`
	Low      SeverityCount `json:"low"`
	Medium   SeverityCount `json:"medium"`
	High     SeverityCount `json:"high"`
	Critical SeverityCount `json:"critical"`
	Total    SeverityCount `json:"total"`
}

type DeltaStatistics struct {
	Created     SeverityBreakdown `json:"created"`
	Closed      SeverityBreakdown `json:"closed"`
	Reactivated SeverityBreakdown `json:"reactivated"`
	Untouched   SeverityBreakdown `json:"untouched"`
}

type ReimportStatistics struct {
	Before SeverityBreakdown `json:"before"`
	Delta  DeltaStatistics   `json:"delta"`
	After  SeverityBreakdown `json:"after"`
}

type ReImportScan struct {
	ScanDate                     *string             `json:"scan_date,omitempty"`
	MinimumSeverity              *string             `json:"minimum_severity,omitempty"`
	Active                       *bool               `json:"active,omitempty"`
	Verified                     *bool               `json:"verified,omitempty"`
	ScanType                     *string             `json:"scan_type,omitempty"`
	EndpointToAdd                *int                `json:"endpoint_to_add,omitempty"`
	File                         *string             `json:"file,omitempty"`
	ProductTypeName              *string             `json:"product_type_name,omitempty"`
	ProductName                  *string             `json:"product_name,omitempty"`
	EngagementName               *string             `json:"engagement_name,omitempty"`
	CloseOldFindingsProductScope *bool               `json:"close_old_findings_product_scope,omitempty"`
	DoNotReactivate              *bool               `json:"do_not_reactivate,omitempty"`
	TestTitle                    *string             `json:"test_title,omitempty"`
	AutoCreateContext            *bool               `json:"auto_create_context,omitempty"`
	DeduplicationOnEngagement    *bool               `json:"deduplication_on_engagement,omitempty"`
	Lead                         *int                `json:"lead,omitempty"`
	Tags                         *[]string           `json:"tags,omitempty"`
	CloseOldFindings             *bool               `json:"close_old_findings,omitempty"`
	PushToJira                   *bool               `json:"push_to_jira,omitempty"`
	Environment                  *string             `json:"environment,omitempty"`
	Version                      *string             `json:"version,omitempty"`
	BuildId                      *string             `json:"build_id,omitempty"`
	BranchTag                    *string             `json:"branch_tag,omitempty"`
	CommitHash                   *string             `json:"commit_hash,omitempty"`
	ApiScanConfiguration         *int                `json:"api_scan_configuration,omitempty"`
	Service                      *string             `json:"service,omitempty"`
	GroupBy                      *string             `json:"group_by,omitempty"`
	Test                         *int                `json:"test,omitempty"`
	TestId                       *int                `json:"test_id,omitempty"`
	EngagementId                 *int                `json:"engagement_id,omitempty"`
	ProductId                    *int                `json:"product_id,omitempty"`
	ProductTypeId                *int                `json:"product_type_id,omitempty"`
	Statistics                   *ReimportStatistics `json:"statistics,omitempty"`
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
