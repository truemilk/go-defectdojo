// Package defectdojo provides a Go client library for accessing the DefectDojo API v2.
// This file contains the ProductTypes service implementation for managing product types.
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

// ProductTypesService handles communication with the product types related methods of the DefectDojo API.
// It provides methods to list, read, create, update, and delete product types in DefectDojo.
type ProductTypesService struct {
	client *Client
}

// ProductType represents a product type in DefectDojo, which is a categorization
// that helps organize and classify products based on their characteristics.
type ProductType struct {
	// Id is the unique identifier for the product type
	Id *int `json:"id,omitempty"`
	// Name is the human-readable name of the product type
	Name *string `json:"name,omitempty"`
	// Description provides additional details about the product type
	Description *string `json:"description,omitempty"`
	// CriticalProduct indicates if products of this type are considered critical
	CriticalProduct *bool `json:"critical_product,omitempty"`
	// KeyProduct indicates if products of this type are considered key products
	KeyProduct *bool `json:"key_product,omitempty"`
	// Updated is the timestamp when the product type was last updated
	Updated *time.Time `json:"updated,omitempty"`
	// Created is the timestamp when the product type was created
	Created *time.Time `json:"created,omitempty"`
	// Members contains the IDs of users who are members of this product type
	Members *[]int `json:"members,omitempty"`
	// AuthorizationGroups contains the IDs of authorization groups for this product type
	AuthorizationGroups *[]int `json:"authorization_groups,omitempty"`
}

// ProductTypes represents a paginated response containing multiple product types from the DefectDojo API.
type ProductTypes struct {
	// Count is the total number of product types matching the query
	Count *int `json:"count,omitempty"`
	// Next is the URL for the next page of results
	Next *string `json:"next,omitempty"`
	// Previous is the URL for the previous page of results
	Previous *string `json:"previous,omitempty"`
	// Results contains the actual product type data for the current page
	Results *[]ProductType `json:"results,omitempty"`
	// Prefetch contains related objects that were prefetched to reduce API calls
	Prefetch *struct {
		// AuthorizationGroups maps authorization group IDs to their full objects
		AuthorizationGroups *map[string]DojoGroup `json:"authorization_groups,omitempty"`
		// Members maps user IDs to their full user objects
		Members *map[string]User `json:"members,omitempty"`
	} `json:"prefetch,omitempty"`
}

// ProductTypesOptions contains optional parameters for filtering and paginating product types.
type ProductTypesOptions struct {
	// Limit specifies the maximum number of product types to return (default: 20)
	Limit int
	// Offset specifies the starting position for pagination
	Offset int
	// ID filters product types by their unique identifier
	ID int
	// Name filters product types by name (partial match)
	Name string
	// CriticalProduct filters by critical product status ("true" or "false")
	CriticalProduct string
	// KeyProduct filters by key product status ("true" or "false")
	KeyProduct string
	// Created filters by creation date (ISO 8601 format)
	Created string
	// Updated filters by last updated date (ISO 8601 format)
	Updated string
	// Prefetch specifies related objects to include in the response to reduce API calls
	Prefetch string
}

// ToString converts ProductTypesOptions to a URL query string for API requests.
// It returns an empty string if no options are set, otherwise returns a query string
// starting with '?' and containing the appropriate parameters.
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

// List retrieves a paginated list of product types from DefectDojo.
// It accepts optional filtering and pagination parameters through ProductTypesOptions.
// The returned ProductTypes struct contains the results and pagination information.
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

// Read retrieves a single product type by its ID from DefectDojo.
// It returns the complete product type information including all fields and relationships.
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

// Create creates a new product type in DefectDojo.
// It accepts a ProductType struct with the desired fields set and returns the created product type
// with server-generated fields like ID and timestamps populated.
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

// Update performs a full update of a product type in DefectDojo.
// It replaces all fields of the product type with the provided values.
// Fields not specified in the ProductType struct will be set to their zero values.
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

// PartialUpdate performs a partial update of a product type in DefectDojo.
// It only updates the fields that are specified in the ProductType struct.
// Fields not specified will remain unchanged on the server.
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

// Delete removes a product type from DefectDojo by its ID.
// It returns the deleted product type information. Note that deleting a product type
// may affect associated products and should be done with caution.
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
