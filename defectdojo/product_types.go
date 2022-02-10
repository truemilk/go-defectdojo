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

type ProductTypesService struct {
	client *Client
}

type ProductType struct {
	Id                  *int       `json:"id,omitempty"`
	Name                *string    `json:"name,omitempty"`
	Description         *string    `json:"description,omitempty"`
	CriticalProduct     *bool      `json:"critical_product,omitempty"`
	KeyProduct          *bool      `json:"key_product,omitempty"`
	Updated             *time.Time `json:"updated,omitempty"`
	Created             *time.Time `json:"created,omitempty"`
	Members             *[]int     `json:"members,omitempty"`
	AuthorizationGroups *[]int     `json:"authorization_groups,omitempty"`
}

type ProductTypes struct {
	Count    *int           `json:"count,omitempty"`
	Next     *string        `json:"next,omitempty"`
	Previous *string        `json:"previous,omitempty"`
	Results  *[]ProductType `json:"results,omitempty"`
	Prefetch *struct {
		AuthorizationGroups *map[string]DojoGroup `json:"authorization_groups,omitempty"`
		Members             *map[string]User      `json:"members,omitempty"`
	} `json:"prefetch,omitempty"`
}

type ProductTypesOptions struct {
	Limit           int
	Offset          int
	ID              int
	Name            string
	CriticalProduct string
	KeyProduct      string
	Created         string
	Updated         string
	Prefetch        string
}

func (o *ProductTypesOptions) ToString() string {
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
		if o.ID > 0 {
			opts = append(opts, fmt.Sprintf("id=%d", o.ID))
		}
		if len(o.Name) > 0 {
			opts = append(opts, fmt.Sprintf("name=%s", o.Name))
		}
		if len(o.CriticalProduct) > 0 {
			opts = append(opts, fmt.Sprintf("critical_product=%s", o.CriticalProduct))
		}
		if len(o.KeyProduct) > 0 {
			opts = append(opts, fmt.Sprintf("key_product=%s", o.KeyProduct))
		}
		if len(o.Created) > 0 {
			opts = append(opts, fmt.Sprintf("created=%s", o.Created))
		}
		if len(o.Updated) > 0 {
			opts = append(opts, fmt.Sprintf("updated=%s", o.Updated))
		}
		if len(o.Prefetch) > 0 {
			opts = append(opts, fmt.Sprintf("prefetch=%s", o.Prefetch))
		}
		optsString += strings.Join(opts, "&")
	}
	return optsString
}

func (c *ProductTypesService) List(ctx context.Context, options *ProductTypesOptions) (*ProductTypes, error) {
	path := fmt.Sprintf("%s/product_types/%s", c.client.BaseURL, options.ToString())

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := ProductTypes{}
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *ProductTypesService) Read(ctx context.Context, id int) (*ProductType, error) {
	path := fmt.Sprintf("%s/product_types/%d/", c.client.BaseURL, id)

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(ProductType)
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ProductTypesService) Create(ctx context.Context, u *ProductType) (*ProductType, error) {
	path := fmt.Sprintf("%s/product_types/", c.client.BaseURL)

	postJSON, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(ProductType)
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ProductTypesService) Update(ctx context.Context, id int, u *ProductType) (*ProductType, error) {
	path := fmt.Sprintf("%s/product_types/%d/", c.client.BaseURL, id)

	postJSON, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPut, path, bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(ProductType)
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ProductTypesService) PartialUpdate(ctx context.Context, id int, u *ProductType) (*ProductType, error) {
	path := fmt.Sprintf("%s/product_types/%d/", c.client.BaseURL, id)

	postJSON, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPatch, path, bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(ProductType)
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ProductTypesService) Delete(ctx context.Context, id int) (*ProductType, error) {
	path := fmt.Sprintf("%s/product_types/%d/", c.client.BaseURL, id)

	req, err := http.NewRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(ProductType)
	if err := c.client.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
