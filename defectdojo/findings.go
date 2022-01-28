package defectdojo

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type FindingsService struct {
	client *Client
}

type Finding struct {
	Id              *int      `json:"id,omitempty"`
	Tags            []*string `json:"tags,omitempty"`
	RequestResponse *struct {
		ReqResp []*struct {
			AdditionalProp1 *string `json:"additionalProp1,omitempty"`
			AdditionalProp2 *string `json:"additionalProp2,omitempty"`
			AdditionalProp3 *string `json:"additionalProp3,omitempty"`
		} `json:"req_resp,omitempty"`
	} `json:"request_response,omitempty"`
	AcceptedRisks []*struct {
		Id                    *int       `json:"id,omitempty"`
		Name                  *string    `json:"name,omitempty"`
		Recommendation        *string    `json:"recommendation,omitempty"`
		RecommendationDetails *string    `json:"recommendation_details,omitempty"`
		Decision              *string    `json:"decision,omitempty"`
		DecisionDetails       *string    `json:"decision_details,omitempty"`
		AcceptedBy            *string    `json:"accepted_by,omitempty"`
		Path                  *string    `json:"path,omitempty"`
		ExpirationDate        *time.Time `json:"expiration_date,omitempty"`
		ExpirationDateWarned  *time.Time `json:"expiration_date_warned,omitempty"`
		ExpirationDateHandled *time.Time `json:"expiration_date_handled,omitempty"`
		ReactivateExpired     *bool      `json:"reactivate_expired,omitempty"`
		RestartSlaExpired     *bool      `json:"restart_sla_expired,omitempty"`
		Created               *time.Time `json:"created,omitempty"`
		Updated               *time.Time `json:"updated,omitempty"`
		Owner                 *int       `json:"owner,omitempty"`
		AcceptedFindings      []*int     `json:"accepted_findings,omitempty"`
		Notes                 []*int     `json:"notes,omitempty"`
	} `json:"accepted_risks,omitempty"`
	PushToJira       *bool `json:"push_to_jira,omitempty"`
	Age              *int  `json:"age,omitempty"`
	SlaDaysRemaining *int  `json:"sla_days_remaining,omitempty"`
	FindingMeta      []*struct {
		Name  *string `json:"name,omitempty"`
		Value *string `json:"value,omitempty"`
	} `json:"finding_meta,omitempty"`
	RelatedFields *struct {
		Test *struct {
			Id       *int    `json:"id,omitempty"`
			Title    *string `json:"title,omitempty"`
			TestType *struct {
				Id   *int    `json:"id,omitempty"`
				Name *string `json:"name,omitempty"`
			} `json:"test_type,omitempty"`
			Engagement *struct {
				Id      *int    `json:"id,omitempty"`
				Name    *string `json:"name,omitempty"`
				Product *struct {
					Id       *int    `json:"id,omitempty"`
					Name     *string `json:"name,omitempty"`
					ProdType *struct {
						Id   *int    `json:"id,omitempty"`
						Name *string `json:"name,omitempty"`
					} `json:"prod_type,omitempty"`
				} `json:"product,omitempty"`
				BranchTag  *string `json:"branch_tag,omitempty"`
				BuildId    *string `json:"build_id,omitempty"`
				CommitHash *string `json:"commit_hash,omitempty"`
				Version    *string `json:"version,omitempty"`
			} `json:"engagement,omitempty"`
			Environment *struct {
				Id   *int    `json:"id,omitempty"`
				Name *string `json:"name,omitempty"`
			} `json:"environment,omitempty"`
			BranchTag  *string `json:"branch_tag,omitempty"`
			BuildId    *string `json:"build_id,omitempty"`
			CommitHash *string `json:"commit_hash,omitempty"`
			Version    *string `json:"version,omitempty"`
		} `json:"test,omitempty"`
		Jira *struct {
			Id           *int       `json:"id,omitempty"`
			Url          *string    `json:"url,omitempty"`
			JiraId       *string    `json:"jira_id,omitempty"`
			JiraKey      *string    `json:"jira_key,omitempty"`
			JiraCreation *time.Time `json:"jira_creation,omitempty"`
			JiraChange   *time.Time `json:"jira_change,omitempty"`
			JiraProject  *int       `json:"jira_project,omitempty"`
			Finding      *int       `json:"finding,omitempty"`
			Engagement   *int       `json:"engagement,omitempty"`
			FindingGroup *int       `json:"finding_group,omitempty"`
		} `json:"jira,omitempty"`
	} `json:"related_fields,omitempty"`
	JiraCreation  *time.Time `json:"jira_creation,omitempty"`
	JiraChange    *time.Time `json:"jira_change,omitempty"`
	DisplayStatus *string    `json:"display_status,omitempty"`
	FindingGroups []struct {
		Id        *int    `json:"id,omitempty"`
		Name      *string `json:"name,omitempty"`
		Test      *int    `json:"test,omitempty"`
		JiraIssue *struct {
			Id           *int       `json:"id,omitempty"`
			Url          *string    `json:"url,omitempty"`
			JiraId       *string    `json:"jira_id,omitempty"`
			JiraKey      *string    `json:"jira_key,omitempty"`
			JiraCreation *time.Time `json:"jira_creation,omitempty"`
			JiraChange   *time.Time `json:"jira_change,omitempty"`
			JiraProject  *int       `json:"jira_project,omitempty"`
			Finding      *int       `json:"finding,omitempty"`
			Engagement   *int       `json:"engagement,omitempty"`
			FindingGroup *int       `json:"finding_group,omitempty"`
		} `json:"jira_issue,omitempty"`
	} `json:"finding_groups,omitempty"`
	Title                   *string    `json:"title,omitempty"`
	Date                    *string    `json:"date,omitempty"`
	SlaStartDate            *string    `json:"sla_start_date,omitempty"`
	Cwe                     *int       `json:"cwe,omitempty"`
	Cve                     *string    `json:"cve,omitempty"`
	Cvssv3                  *string    `json:"cvssv3,omitempty"`
	Cvssv3Score             *float32   `json:"cvssv3_score,omitempty"`
	Url                     *string    `json:"url,omitempty"`
	Severity                *string    `json:"severity,omitempty"`
	Description             *string    `json:"description,omitempty"`
	Mitigation              *string    `json:"mitigation,omitempty"`
	Impact                  *string    `json:"impact,omitempty"`
	StepsToReproduce        *string    `json:"steps_to_reproduce,omitempty"`
	SeverityJustification   *string    `json:"severity_justification,omitempty"`
	References              *string    `json:"references,omitempty"`
	Active                  *bool      `json:"active,omitempty"`
	Verified                *bool      `json:"verified,omitempty"`
	FalseP                  *bool      `json:"false_p,omitempty"`
	Duplicate               *bool      `json:"duplicate,omitempty"`
	OutOfScope              *bool      `json:"out_of_scope,omitempty"`
	RiskAccepted            *bool      `json:"risk_accepted,omitempty"`
	UnderReview             *bool      `json:"under_review,omitempty"`
	LastStatusUpdate        *time.Time `json:"last_status_update,omitempty"`
	UnderDefectReview       *bool      `json:"under_defect_review,omitempty"`
	IsMitigated             *bool      `json:"is_mitigated,omitempty"`
	ThreadId                *int       `json:"thread_id,omitempty"`
	Mitigated               *time.Time `json:"mitigated,omitempty"`
	NumericalSeverity       *string    `json:"numerical_severity,omitempty"`
	LastReviewed            *time.Time `json:"last_reviewed,omitempty"`
	Param                   *string    `json:"param,omitempty"`
	Payload                 *string    `json:"payload,omitempty"`
	HashCode                *string    `json:"hash_code,omitempty"`
	Line                    *int       `json:"line,omitempty"`
	FilePath                *string    `json:"file_path,omitempty"`
	ComponentName           *string    `json:"component_name,omitempty"`
	ComponentVersion        *string    `json:"component_version,omitempty"`
	StaticFinding           *bool      `json:"static_finding,omitempty"`
	DynamicFinding          *bool      `json:"dynamic_finding,omitempty"`
	Created                 *time.Time `json:"created,omitempty"`
	ScannerConfidence       *int       `json:"scanner_confidence,omitempty"`
	UniqueIdFromTool        *string    `json:"unique_id_from_tool,omitempty"`
	VulnIdFromTool          *string    `json:"vuln_id_from_tool,omitempty"`
	SastSourceObject        *string    `json:"sast_source_object,omitempty"`
	SastSinkObject          *string    `json:"sast_sink_object,omitempty"`
	SastSourceLine          *int       `json:"sast_source_line,omitempty"`
	SastSourceFilePath      *string    `json:"sast_source_file_path,omitempty"`
	NbOccurences            *int       `json:"nb_occurences,omitempty"`
	PublishDate             *string    `json:"publish_date,omitempty"`
	Service                 *string    `json:"service,omitempty"`
	Test                    *int       `json:"test,omitempty"`
	DuplicateFinding        *int       `json:"duplicate_finding,omitempty"`
	ReviewRequestedBy       *int       `json:"review_requested_by,omitempty"`
	DefectReviewRequestedBy *int       `json:"defect_review_requested_by,omitempty"`
	MitigatedBy             *int       `json:"mitigated_by,omitempty"`
	Reporter                *int       `json:"reporter,omitempty"`
	LastReviewedBy          *int       `json:"last_reviewed_by,omitempty"`
	SonarqubeIssue          *int       `json:"sonarqube_issue,omitempty"`
	Endpoints               []*int     `json:"endpoints,omitempty"`
	EndpointStatus          []*int     `json:"endpoint_status,omitempty"`
	Reviewers               []*int     `json:"reviewers,omitempty"`
	Notes                   []*struct {
		Id     *int `json:"id,omitempty"`
		Author *struct {
			Id        *int    `json:"id,omitempty"`
			Username  *string `json:"username,omitempty"`
			FirstName *string `json:"first_name,omitempty"`
			LastName  *string `json:"last_name,omitempty"`
		} `json:"author,omitempty"`
		Editor *struct {
			Id        *int    `json:"id,omitempty"`
			Username  *string `json:"username,omitempty"`
			FirstName *string `json:"first_name,omitempty"`
			LastName  *string `json:"last_name,omitempty"`
		} `json:"editor,omitempty"`
		History []*struct {
			Id            *int `json:"id,omitempty"`
			CurrentEditor *struct {
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
	Files   []*int `json:"files,omitempty"`
	FoundBy []*int `json:"found_by,omitempty"`
}

type Findings struct {
	Count    *int       `json:"count,omitempty"`
	Next     *string    `json:"next,omitempty"`
	Previous *string    `json:"previous,omitempty"`
	Results  []*Finding `json:"results,omitempty"`
	Prefetch *struct {
		DuplicateFinding map[string]Finding `json:"duplicate_finding,omitempty"`
	} `json:"prefetch,omitempty"`
}

type FindingsOptions struct {
	Limit    int
	Offset   int
	Prefetch string
}

func (o *FindingsOptions) ToString() string {
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
		if len(o.Prefetch) > 0 {
			opts = append(opts, fmt.Sprintf("prefetch=%s", o.Prefetch))
		}
		optsString += strings.Join(opts, "&")
	}
	return optsString
}

func (c *FindingsService) List(ctx context.Context, options *FindingsOptions) (*Findings, error) {
	path := fmt.Sprintf("%s/findings/%s", c.client.BaseURL, options.ToString())

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Findings{}
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *FindingsService) Read(ctx context.Context, id int) (*Finding, error) {
	path := fmt.Sprintf("%s/findings/%d/", c.client.BaseURL, id)

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(Finding)
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
