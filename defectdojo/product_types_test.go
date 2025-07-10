package defectdojo

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestProductTypesService_List(t *testing.T) {
	response := `{
		"count": 2,
		"next": null,
		"previous": null,
		"results": [
			{
				"id": 1,
				"name": "Web Application",
				"description": "Web application product type",
				"critical_product": true,
				"key_product": false,
				"created": "2022-01-01T12:00:00Z",
				"updated": "2022-01-02T12:00:00Z",
				"members": [1, 2, 3],
				"authorization_groups": [1, 2]
			},
			{
				"id": 2,
				"name": "Mobile Application",
				"description": "Mobile application product type",
				"critical_product": false,
				"key_product": true,
				"created": "2022-01-03T12:00:00Z",
				"updated": "2022-01-04T12:00:00Z",
				"members": [2, 3],
				"authorization_groups": [2]
			}
		]
	}`

	expected := ProductTypes{
		Count:    Int(2),
		Next:     nil,
		Previous: nil,
		Results: &[]ProductType{
			{
				Id:                  Int(1),
				Name:                Str("Web Application"),
				Description:         Str("Web application product type"),
				CriticalProduct:     Bool(true),
				KeyProduct:          Bool(false),
				Created:             Date(time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC)),
				Updated:             Date(time.Date(2022, 1, 2, 12, 0, 0, 0, time.UTC)),
				Members:             &[]int{1, 2, 3},
				AuthorizationGroups: &[]int{1, 2},
			},
			{
				Id:                  Int(2),
				Name:                Str("Mobile Application"),
				Description:         Str("Mobile application product type"),
				CriticalProduct:     Bool(false),
				KeyProduct:          Bool(true),
				Created:             Date(time.Date(2022, 1, 3, 12, 0, 0, 0, time.UTC)),
				Updated:             Date(time.Date(2022, 1, 4, 12, 0, 0, 0, time.UTC)),
				Members:             &[]int{2, 3},
				AuthorizationGroups: &[]int{2},
			},
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/product_types/") {
			t.Errorf("Expected /product_types/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.ProductTypes.List(context.Background(), &ProductTypesOptions{
		Limit:  10,
		Offset: 0,
	})
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestProductTypesService_Read(t *testing.T) {
	response := `{
		"id": 123,
		"name": "Test Product Type",
		"description": "A test product type",
		"critical_product": true,
		"key_product": false,
		"created": "2022-01-01T12:00:00Z",
		"updated": "2022-01-02T12:00:00Z",
		"members": [1, 2, 3],
		"authorization_groups": [1, 2]
	}`

	expected := ProductType{
		Id:                  Int(123),
		Name:                Str("Test Product Type"),
		Description:         Str("A test product type"),
		CriticalProduct:     Bool(true),
		KeyProduct:          Bool(false),
		Created:             Date(time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC)),
		Updated:             Date(time.Date(2022, 1, 2, 12, 0, 0, 0, time.UTC)),
		Members:             &[]int{1, 2, 3},
		AuthorizationGroups: &[]int{1, 2},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/product_types/123/") {
			t.Errorf("Expected /product_types/123/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.ProductTypes.Read(context.Background(), 123)
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestProductTypesService_Create(t *testing.T) {
	response := `{
		"id": 456,
		"name": "New Product Type",
		"description": "A new product type",
		"critical_product": false,
		"key_product": true,
		"created": "2022-01-01T12:00:00Z",
		"updated": "2022-01-01T12:00:00Z",
		"members": [1],
		"authorization_groups": [1]
	}`

	expected := ProductType{
		Id:                  Int(456),
		Name:                Str("New Product Type"),
		Description:         Str("A new product type"),
		CriticalProduct:     Bool(false),
		KeyProduct:          Bool(true),
		Created:             Date(time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC)),
		Updated:             Date(time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC)),
		Members:             &[]int{1},
		AuthorizationGroups: &[]int{1},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/product_types/") {
			t.Errorf("Expected /product_types/ in path, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.ProductTypes.Create(context.Background(), &ProductType{
		Name:            Str("New Product Type"),
		Description:     Str("A new product type"),
		CriticalProduct: Bool(false),
		KeyProduct:      Bool(true),
	})
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestProductTypesService_Update(t *testing.T) {
	response := `{
		"id": 789,
		"name": "Updated Product Type",
		"description": "An updated product type",
		"critical_product": true,
		"key_product": false,
		"created": "2022-01-01T12:00:00Z",
		"updated": "2022-01-03T12:00:00Z",
		"members": [1, 2],
		"authorization_groups": [1, 2]
	}`

	expected := ProductType{
		Id:                  Int(789),
		Name:                Str("Updated Product Type"),
		Description:         Str("An updated product type"),
		CriticalProduct:     Bool(true),
		KeyProduct:          Bool(false),
		Created:             Date(time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC)),
		Updated:             Date(time.Date(2022, 1, 3, 12, 0, 0, 0, time.UTC)),
		Members:             &[]int{1, 2},
		AuthorizationGroups: &[]int{1, 2},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/product_types/789/") {
			t.Errorf("Expected /product_types/789/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.ProductTypes.Update(context.Background(), 789, &ProductType{
		Name:            Str("Updated Product Type"),
		Description:     Str("An updated product type"),
		CriticalProduct: Bool(true),
		KeyProduct:      Bool(false),
	})
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestProductTypesService_PartialUpdate(t *testing.T) {
	response := `{
		"id": 321,
		"name": "Partially Updated Product Type",
		"description": "A partially updated product type",
		"critical_product": false,
		"key_product": true,
		"created": "2022-01-01T12:00:00Z",
		"updated": "2022-01-04T12:00:00Z",
		"members": [1],
		"authorization_groups": [1]
	}`

	expected := ProductType{
		Id:                  Int(321),
		Name:                Str("Partially Updated Product Type"),
		Description:         Str("A partially updated product type"),
		CriticalProduct:     Bool(false),
		KeyProduct:          Bool(true),
		Created:             Date(time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC)),
		Updated:             Date(time.Date(2022, 1, 4, 12, 0, 0, 0, time.UTC)),
		Members:             &[]int{1},
		AuthorizationGroups: &[]int{1},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("Expected PATCH request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/product_types/321/") {
			t.Errorf("Expected /product_types/321/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.ProductTypes.PartialUpdate(context.Background(), 321, &ProductType{
		Name: Str("Partially Updated Product Type"),
	})
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestProductTypesService_Delete(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/product_types/654/") {
			t.Errorf("Expected /product_types/654/ in path, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintln(w, "{}")
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.ProductTypes.Delete(context.Background(), 654)
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if actual == nil {
		t.Errorf("expected non-nil response")
	}
}

func TestProductTypesOptions_ToString(t *testing.T) {
	tests := []struct {
		name     string
		options  *ProductTypesOptions
		expected string
	}{
		{
			name:     "nil options",
			options:  nil,
			expected: "",
		},
		{
			name:     "empty options",
			options:  &ProductTypesOptions{},
			expected: "?",
		},
		{
			name: "limit only",
			options: &ProductTypesOptions{
				Limit: 10,
			},
			expected: "?limit=10",
		},
		{
			name: "offset only",
			options: &ProductTypesOptions{
				Offset: 20,
			},
			expected: "?offset=20",
		},
		{
			name: "id only",
			options: &ProductTypesOptions{
				ID: 123,
			},
			expected: "?id=123",
		},
		{
			name: "name only",
			options: &ProductTypesOptions{
				Name: "test",
			},
			expected: "?name=test",
		},
		{
			name: "critical_product only",
			options: &ProductTypesOptions{
				CriticalProduct: "true",
			},
			expected: "?critical_product=true",
		},
		{
			name: "key_product only",
			options: &ProductTypesOptions{
				KeyProduct: "false",
			},
			expected: "?key_product=false",
		},
		{
			name: "created only",
			options: &ProductTypesOptions{
				Created: "2022-01-01",
			},
			expected: "?created=2022-01-01",
		},
		{
			name: "updated only",
			options: &ProductTypesOptions{
				Updated: "2022-01-02",
			},
			expected: "?updated=2022-01-02",
		},
		{
			name: "prefetch only",
			options: &ProductTypesOptions{
				Prefetch: "members",
			},
			expected: "?prefetch=members",
		},
		{
			name: "all options",
			options: &ProductTypesOptions{
				Limit:           10,
				Offset:          20,
				ID:              123,
				Name:            "test",
				CriticalProduct: "true",
				KeyProduct:      "false",
				Created:         "2022-01-01",
				Updated:         "2022-01-02",
				Prefetch:        "members",
			},
			expected: "?limit=10&offset=20&id=123&name=test&critical_product=true&key_product=false&created=2022-01-01&updated=2022-01-02&prefetch=members",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.options.ToString()
			if actual != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, actual)
			}
		})
	}
}