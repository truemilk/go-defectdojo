package defectdojo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Engagement struct {
	Id                         *int      `json:"id,omitempty"`
	Tags                       []*string `json:"tags,omitempty"`
	Name                       *string   `json:"name,omitempty"`
	Description                *string   `json:"description,omitempty"`
	Version                    *string   `json:"version,omitempty"`
	FirstContacted             *string   `json:"first_contacted,omitempty"`
	TargetStart                *string   `json:"target_start,omitempty"`
	TargetEnd                  *string   `json:"target_end,omitempty"`
	Reason                     *string   `json:"reason,omitempty"`
	Updated                    *string   `json:"updated,omitempty"`
	Created                    *string   `json:"created,omitempty"`
	Active                     *bool     `json:"active,omitempty"`
	Tracker                    *string   `json:"tracker,omitempty"`
	TestStrategy               *string   `json:"test_strategy,omitempty"`
	ThreatModel                *bool     `json:"threat_model,omitempty"`
	ApiTest                    *bool     `json:"api_test,omitempty"`
	PenTest                    *bool     `json:"pen_test,omitempty"`
	CheckList                  *bool     `json:"check_list,omitempty"`
	Status                     *string   `json:"status,omitempty"`
	Progress                   *string   `json:"progress,omitempty"`
	TmodelPath                 *string   `json:"tmodel_path,omitempty"`
	DoneTesting                *bool     `json:"done_testing,omitempty"`
	EngagementType             *string   `json:"engagement_type,omitempty"`
	BuildId                    *string   `json:"build_id,omitempty"`
	CommitHash                 *string   `json:"commit_hash,omitempty"`
	BranchTag                  *string   `json:"branch_tag,omitempty"`
	SourceCodeManagementUri    *string   `json:"source_code_management_uri,omitempty"`
	DeduplicationOnEngagement  *bool     `json:"deduplication_on_engagement,omitempty"`
	Lead                       *int      `json:"lead,omitempty"`
	Requester                  *int      `json:"requester,omitempty"`
	Preset                     *int      `json:"preset,omitempty"`
	ReportType                 *int      `json:"report_type,omitempty"`
	Product                    *int      `json:"product,omitempty"`
	BuildServer                *int      `json:"build_server,omitempty"`
	SourceCodeManagementServer *int      `json:"source_code_management_server,omitempty"`
	OrchestrationEngine        *int      `json:"orchestration_engine,omitempty"`
	Notes                      []*struct {
		Id     *int `json:"id,omitempty"`
		Author struct {
			Id        *int    `json:"id,omitempty"`
			Username  *string `json:"username,omitempty"`
			FirstName *string `json:"first_name,omitempty"`
			LastName  *string `json:"last_name,omitempty"`
		} `json:"author,omitempty"`
		Editor struct {
			Id        *int    `json:"id,omitempty"`
			Username  *string `json:"username,omitempty"`
			FirstName *string `json:"first_name,omitempty"`
			LastName  *string `json:"last_name,omitempty"`
		} `json:"editor,omitempty"`
		History []*struct {
			Id            *int `json:"id,omitempty"`
			CurrentEditor struct {
				Id        *int    `json:"id,omitempty"`
				Username  *string `json:"username,omitempty"`
				FirstName *string `json:"first_name,omitempty"`
				LastName  *string `json:"last_name,omitempty"`
			} `json:"current_editor,omitempty"`
			Data     *string    `json:"data,omitempty"`
			Time     *time.Time `json:"time,omitempty"`
			NoteType *int       `json:"note_type,omitempty"`
		} `json:"history,omitempty"`
		Entry    *string    `json:"entry,omitempty"`
		Date     *time.Time `json:"date,omitempty"`
		Private  *bool      `json:"private,omitempty"`
		Edited   *bool      `json:"edited,omitempty"`
		EditTime *time.Time `json:"edit_time,omitempty"`
		NoteType *int       `json:"note_type,omitempty"`
	} `json:"notes,omitempty"`
	Files []*struct {
		Id    *int    `json:"id,omitempty"`
		File  *string `json:"file,omitempty"`
		Title *string `json:"title,omitempty"`
	} `json:"files,omitempty"`
	RiskAcceptance []*int `json:"risk_acceptance,omitempty"`
}

type Engagements struct {
	Count    int          `json:"count,omitempty"`
	Next     string       `json:"next,omitempty"`
	Previous string       `json:"previous,omitempty"`
	Results  []Engagement `json:"results,omitempty"`
}

type EngagementsOptions struct {
	Limit  int
	Offset int
}

func (o *EngagementsOptions) ToString() string {
	var opts []string
	var optsString string
	if o != nil {
		optsString += "?"
		if o.Limit > 0 {
			opts = append(opts, fmt.Sprintf("limit=%d", o.Limit))
		}
		if o.Offset > 0 {
			opts = append(opts, fmt.Sprintf("offset=%d", o.Offset))
		}
		optsString += strings.Join(opts, "&")
	}
	return optsString
}

func (c *Client) EngagementsList(ctx context.Context, options *EngagementsOptions) (*Engagements, error) {
	path := fmt.Sprintf("%s/engagements/%s", c.BaseURL, options.ToString())

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Engagements{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) EngagementsRead(ctx context.Context, id int) (*Engagement, error) {
	path := fmt.Sprintf("%s/engagements/%d/", c.BaseURL, id)

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(Engagement)
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) EngagementsCreate(ctx context.Context, u *Engagement) (*Engagement, error) {
	path := fmt.Sprintf("%s/engagements/", c.BaseURL)

	postJSON, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(Engagement)
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
