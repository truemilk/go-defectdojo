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

func TestProductsService_List(t *testing.T) {
	response := `{
		"count": 2,
		"next": null,
		"previous": null,
		"results": [
			{
				"id": 1,
				"name": "Test Product 1",
				"description": "A test product",
				"created": "2022-01-01T12:00:00Z",
				"prod_numeric_grade": 95,
				"business_criticality": "high",
				"platform": "web",
				"lifecycle": "production",
				"origin": "internal",
				"user_records": 1000,
				"revenue": "1000000",
				"external_audience": true,
				"internet_accessible": true,
				"enable_simple_risk_acceptance": false,
				"enable_full_risk_acceptance": true,
				"product_manager": 1,
				"technical_contact": 2,
				"team_manager": 3,
				"prod_type": 1,
				"members": [1, 2, 3],
				"authorization_groups": [1, 2],
				"regulations": [1],
				"findings_count": 5,
				"findings_list": [1, 2, 3, 4, 5],
				"tags": ["security", "web"],
				"product_meta": [
					{
						"name": "version",
						"value": "1.0.0"
					}
				]
			},
			{
				"id": 2,
				"name": "Test Product 2",
				"description": "Another test product",
				"created": "2022-01-02T12:00:00Z",
				"prod_numeric_grade": 87,
				"business_criticality": "medium",
				"platform": "mobile",
				"lifecycle": "development",
				"origin": "third_party",
				"user_records": 500,
				"revenue": "500000",
				"external_audience": false,
				"internet_accessible": false,
				"enable_simple_risk_acceptance": true,
				"enable_full_risk_acceptance": false,
				"product_manager": 2,
				"technical_contact": 3,
				"team_manager": 1,
				"prod_type": 2,
				"members": [2, 3],
				"authorization_groups": [2],
				"regulations": [],
				"findings_count": 3,
				"findings_list": [6, 7, 8],
				"tags": ["mobile", "development"],
				"product_meta": []
			}
		]
	}`

	expected := Products{
		Count:    Int(2),
		Next:     nil,
		Previous: nil,
		Results: &[]Product{
			{
				ID:                         Int(1),
				Name:                       Str("Test Product 1"),
				Description:                Str("A test product"),
				Created:                    Date(time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC)),
				ProdNumericGrade:           Int(95),
				BusinessCriticality:        Str("high"),
				Platform:                   Str("web"),
				Lifecycle:                  Str("production"),
				Origin:                     Str("internal"),
				UserRecords:                Int(1000),
				Revenue:                    Str("1000000"),
				ExternalAudience:           Bool(true),
				InternetAccessible:         Bool(true),
				EnableSimpleRiskAcceptance: Bool(false),
				EnableFullRiskAcceptance:   Bool(true),
				ProductManager:             Int(1),
				TechnicalContact:           Int(2),
				TeamManager:                Int(3),
				ProdType:                   Int(1),
				Members:                    &[]int{1, 2, 3},
				AuthorizationGroups:        &[]int{1, 2},
				Regulations:                &[]int{1},
				FindingsCount:              Int(5),
				FindingsList:               &[]int{1, 2, 3, 4, 5},
				Tags:                       &[]string{"security", "web"},
				ProductMeta: &[]struct {
					Name  *string `json:"name,omitempty"`
					Value *string `json:"value,omitempty"`
				}{
					{
						Name:  Str("version"),
						Value: Str("1.0.0"),
					},
				},
			},
			{
				ID:                         Int(2),
				Name:                       Str("Test Product 2"),
				Description:                Str("Another test product"),
				Created:                    Date(time.Date(2022, 1, 2, 12, 0, 0, 0, time.UTC)),
				ProdNumericGrade:           Int(87),
				BusinessCriticality:        Str("medium"),
				Platform:                   Str("mobile"),
				Lifecycle:                  Str("development"),
				Origin:                     Str("third_party"),
				UserRecords:                Int(500),
				Revenue:                    Str("500000"),
				ExternalAudience:           Bool(false),
				InternetAccessible:         Bool(false),
				EnableSimpleRiskAcceptance: Bool(true),
				EnableFullRiskAcceptance:   Bool(false),
				ProductManager:             Int(2),
				TechnicalContact:           Int(3),
				TeamManager:                Int(1),
				ProdType:                   Int(2),
				Members:                    &[]int{2, 3},
				AuthorizationGroups:        &[]int{2},
				Regulations:                &[]int{},
				FindingsCount:              Int(3),
				FindingsList:               &[]int{6, 7, 8},
				Tags:                       &[]string{"mobile", "development"},
				ProductMeta:                &[]struct {
					Name  *string `json:"name,omitempty"`
					Value *string `json:"value,omitempty"`
				}{},
			},
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/products/") {
			t.Errorf("Expected /products/ in path, got %s", r.URL.Path)
		}
		fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.Products.List(context.Background(), &ProductsOptions{
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

func TestProductsService_Read(t *testing.T) {
	response := `{
		"id": 123,
		"name": "Test Product",
		"description": "A test product",
		"created": "2022-01-01T12:00:00Z",
		"prod_numeric_grade": 95,
		"business_criticality": "high",
		"platform": "web",
		"lifecycle": "production",
		"origin": "internal",
		"user_records": 1000,
		"revenue": "1000000",
		"external_audience": true,
		"internet_accessible": true,
		"enable_simple_risk_acceptance": false,
		"enable_full_risk_acceptance": true,
		"product_manager": 1,
		"technical_contact": 2,
		"team_manager": 3,
		"prod_type": 1,
		"members": [1, 2, 3],
		"authorization_groups": [1, 2],
		"regulations": [1],
		"findings_count": 5,
		"findings_list": [1, 2, 3, 4, 5],
		"tags": ["security", "web"],
		"product_meta": [
			{
				"name": "version",
				"value": "1.0.0"
			}
		]
	}`

	expected := Product{
		ID:                         Int(123),
		Name:                       Str("Test Product"),
		Description:                Str("A test product"),
		Created:                    Date(time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC)),
		ProdNumericGrade:           Int(95),
		BusinessCriticality:        Str("high"),
		Platform:                   Str("web"),
		Lifecycle:                  Str("production"),
		Origin:                     Str("internal"),
		UserRecords:                Int(1000),
		Revenue:                    Str("1000000"),
		ExternalAudience:           Bool(true),
		InternetAccessible:         Bool(true),
		EnableSimpleRiskAcceptance: Bool(false),
		EnableFullRiskAcceptance:   Bool(true),
		ProductManager:             Int(1),
		TechnicalContact:           Int(2),
		TeamManager:                Int(3),
		ProdType:                   Int(1),
		Members:                    &[]int{1, 2, 3},
		AuthorizationGroups:        &[]int{1, 2},
		Regulations:                &[]int{1},
		FindingsCount:              Int(5),
		FindingsList:               &[]int{1, 2, 3, 4, 5},
		Tags:                       &[]string{"security", "web"},
		ProductMeta: &[]struct {
			Name  *string `json:"name,omitempty"`
			Value *string `json:"value,omitempty"`
		}{
			{
				Name:  Str("version"),
				Value: Str("1.0.0"),
			},
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/products/123/") {
			t.Errorf("Expected /products/123/ in path, got %s", r.URL.Path)
		}
		fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.Products.Read(context.Background(), 123)
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestProductsService_Create(t *testing.T) {
	response := `{
		"id": 456,
		"name": "New Product",
		"description": "A new product",
		"created": "2022-01-01T12:00:00Z",
		"prod_numeric_grade": 80,
		"business_criticality": "medium",
		"platform": "web",
		"lifecycle": "development",
		"origin": "internal",
		"user_records": 100,
		"revenue": "100000",
		"external_audience": false,
		"internet_accessible": true,
		"enable_simple_risk_acceptance": true,
		"enable_full_risk_acceptance": false,
		"product_manager": 1,
		"technical_contact": 2,
		"team_manager": 3,
		"prod_type": 1,
		"members": [1, 2],
		"authorization_groups": [1],
		"regulations": [],
		"findings_count": 0,
		"findings_list": [],
		"tags": ["new", "development"],
		"product_meta": []
	}`

	expected := Product{
		ID:                         Int(456),
		Name:                       Str("New Product"),
		Description:                Str("A new product"),
		Created:                    Date(time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC)),
		ProdNumericGrade:           Int(80),
		BusinessCriticality:        Str("medium"),
		Platform:                   Str("web"),
		Lifecycle:                  Str("development"),
		Origin:                     Str("internal"),
		UserRecords:                Int(100),
		Revenue:                    Str("100000"),
		ExternalAudience:           Bool(false),
		InternetAccessible:         Bool(true),
		EnableSimpleRiskAcceptance: Bool(true),
		EnableFullRiskAcceptance:   Bool(false),
		ProductManager:             Int(1),
		TechnicalContact:           Int(2),
		TeamManager:                Int(3),
		ProdType:                   Int(1),
		Members:                    &[]int{1, 2},
		AuthorizationGroups:        &[]int{1},
		Regulations:                &[]int{},
		FindingsCount:              Int(0),
		FindingsList:               &[]int{},
		Tags:                       &[]string{"new", "development"},
		ProductMeta:                &[]struct {
			Name  *string `json:"name,omitempty"`
			Value *string `json:"value,omitempty"`
		}{},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/products/") {
			t.Errorf("Expected /products/ in path, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.Products.Create(context.Background(), &Product{
		Name:                Str("New Product"),
		Description:         Str("A new product"),
		BusinessCriticality: Str("medium"),
		Platform:            Str("web"),
		Lifecycle:           Str("development"),
		ProdType:            Int(1),
	})
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestProductsService_Delete(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/products/789/") {
			t.Errorf("Expected /products/789/ in path, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, "{}")
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.Products.Delete(context.Background(), 789)
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if actual == nil {
		t.Errorf("expected non-nil response")
	}
}

func TestProductsOptions_ToString(t *testing.T) {
	tests := []struct {
		name     string
		options  *ProductsOptions
		expected string
	}{
		{
			name:     "nil options",
			options:  nil,
			expected: "",
		},
		{
			name:     "empty options",
			options:  &ProductsOptions{},
			expected: "?",
		},
		{
			name: "limit only",
			options: &ProductsOptions{
				Limit: 10,
			},
			expected: "?limit=10",
		},
		{
			name: "offset only",
			options: &ProductsOptions{
				Offset: 20,
			},
			expected: "?offset=20",
		},
		{
			name: "name only",
			options: &ProductsOptions{
				Name: "test",
			},
			expected: "?name=test",
		},
		{
			name: "prefetch only",
			options: &ProductsOptions{
				Prefetch: "prod_type",
			},
			expected: "?prefetch=prod_type",
		},
		{
			name: "all options",
			options: &ProductsOptions{
				Limit:    10,
				Offset:   20,
				Name:     "test",
				Prefetch: "prod_type",
			},
			expected: "?limit=10&offset=20&name=test&prefetch=prod_type",
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