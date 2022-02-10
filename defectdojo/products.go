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

type ProductsService struct {
	client *Client
}

type Product struct {
	ID            *int      `json:"id,omitempty"`
	FindingsCount *int      `json:"findings_count,omitempty"`
	FindingsList  *[]int    `json:"findings_list,omitempty"`
	Tags          *[]string `json:"tags,omitempty"`
	ProductMeta   *[]struct {
		Name  *string `json:"name,omitempty"`
		Value *string `json:"value,omitempty"`
	} `json:"product_meta,omitempty"`
	Name                       *string    `json:"name,omitempty"`
	Description                *string    `json:"description,omitempty"`
	Created                    *time.Time `json:"created,omitempty"`
	ProdNumericGrade           *int       `json:"prod_numeric_grade,omitempty"`
	BusinessCriticality        *string    `json:"business_criticality,omitempty"`
	Platform                   *string    `json:"platform,omitempty"`
	Lifecycle                  *string    `json:"lifecycle,omitempty"`
	Origin                     *string    `json:"origin,omitempty"`
	UserRecords                *int       `json:"user_records,omitempty"`
	Revenue                    *string    `json:"revenue,omitempty"`
	ExternalAudience           *bool      `json:"external_audience,omitempty"`
	InternetAccessible         *bool      `json:"internet_accessible,omitempty"`
	EnableSimpleRiskAcceptance *bool      `json:"enable_simple_risk_acceptance,omitempty"`
	EnableFullRiskAcceptance   *bool      `json:"enable_full_risk_acceptance,omitempty"`
	ProductManager             *int       `json:"product_manager,omitempty"`
	TechnicalContact           *int       `json:"technical_contact,omitempty"`
	TeamManager                *int       `json:"team_manager,omitempty"`
	ProdType                   *int       `json:"prod_type,omitempty"`
	Members                    *[]int     `json:"members,omitempty"`
	AuthorizationGroups        *[]int     `json:"authorization_groups,omitempty"`
	Regulations                *[]int     `json:"regulations,omitempty"`
}

type Products struct {
	Count    *int       `json:"count,omitempty"`
	Next     *string    `json:"next,omitempty"`
	Previous *string    `json:"previous,omitempty"`
	Results  *[]Product `json:"results,omitempty"`
	Prefetch *struct {
		AuthorizationGroups *map[string]DojoGroup   `json:"authorization_groups,omitempty"`
		Members             *map[string]User        `json:"members,omitempty"`
		ProdType            *map[string]ProductType `json:"prod_type,omitempty"`
		ProductManager      *map[string]User        `json:"product_manager,omitempty"`
		TeamManager         *map[string]User        `json:"team_manager,omitempty"`
		TechnicalContact    *map[string]User        `json:"technical_contact,omitempty"`
	} `json:"prefetch,omitempty"`
}

type ProductsOptions struct {
	Limit    int
	Offset   int
	Name     string
	Prefetch string
}

func (o *ProductsOptions) ToString() string {
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
		if len(o.Name) > 0 {
			opts = append(opts, fmt.Sprintf("name=%s", o.Name))
		}
		if len(o.Prefetch) > 0 {
			opts = append(opts, fmt.Sprintf("prefetch=%s", o.Prefetch))
		}
		optsString += strings.Join(opts, "&")
	}
	return optsString
}

func (c *ProductsService) List(ctx context.Context, options *ProductsOptions) (*Products, error) {
	path := fmt.Sprintf("%s/products/%s", c.client.BaseURL, options.ToString())

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Products{}
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *ProductsService) Read(ctx context.Context, id int) (*Product, error) {
	path := fmt.Sprintf("%s/products/%d/", c.client.BaseURL, id)

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(Product)
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ProductsService) Create(ctx context.Context, u *Product) (*Product, error) {
	path := fmt.Sprintf("%s/products/", c.client.BaseURL)

	postJSON, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(Product)
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ProductsService) Delete(ctx context.Context, id int) (*Product, error) {
	path := fmt.Sprintf("%s/products/%d/", c.client.BaseURL, id)

	req, err := http.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(Product)
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
