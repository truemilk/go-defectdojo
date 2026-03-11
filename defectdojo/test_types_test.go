package defectdojo

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTestTypesService_List(t *testing.T) {
	response := `{
		"count": 1,
		"next": null,
		"previous": null,
		"results": [
			{
				"id": 1,
				"tags": ["sast"],
				"name": "SonarQube Scan",
				"static_tool": true,
				"dynamic_tool": false,
				"active": true
			}
		]
	}`

	expected := TestTypes{
		Count:    Int(1),
		Next:     nil,
		Previous: nil,
		Results: &[]TestType{
			{
				Id:          Int(1),
				Tags:        &[]string{"sast"},
				Name:        Str("SonarQube Scan"),
				StaticTool:  Bool(true),
				DynamicTool: Bool(false),
				Active:      Bool(true),
			},
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/test_types/") {
			t.Errorf("Expected /test_types/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.TestTypes.List(context.Background(), &TestTypesOptions{
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

func TestTestTypesService_Read(t *testing.T) {
	response := `{
		"id": 123,
		"tags": ["dast"],
		"name": "ZAP Scan",
		"static_tool": false,
		"dynamic_tool": true,
		"active": true
	}`

	expected := TestType{
		Id:          Int(123),
		Tags:        &[]string{"dast"},
		Name:        Str("ZAP Scan"),
		StaticTool:  Bool(false),
		DynamicTool: Bool(true),
		Active:      Bool(true),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/test_types/123/") {
			t.Errorf("Expected /test_types/123/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.TestTypes.Read(context.Background(), 123)
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestTestTypesService_Create(t *testing.T) {
	response := `{
		"id": 456,
		"tags": ["sca"],
		"name": "Dependency Check",
		"static_tool": true,
		"dynamic_tool": false,
		"active": true
	}`

	expected := TestType{
		Id:          Int(456),
		Tags:        &[]string{"sca"},
		Name:        Str("Dependency Check"),
		StaticTool:  Bool(true),
		DynamicTool: Bool(false),
		Active:      Bool(true),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/test_types/") {
			t.Errorf("Expected /test_types/ in path, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.TestTypes.Create(context.Background(), &TestType{
		Name:        Str("Dependency Check"),
		StaticTool:  Bool(true),
		DynamicTool: Bool(false),
		Active:      Bool(true),
	})
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestTestTypesService_Update(t *testing.T) {
	response := `{
		"id": 789,
		"tags": ["sast"],
		"name": "Updated Scan Type",
		"static_tool": true,
		"dynamic_tool": true,
		"active": true
	}`

	expected := TestType{
		Id:          Int(789),
		Tags:        &[]string{"sast"},
		Name:        Str("Updated Scan Type"),
		StaticTool:  Bool(true),
		DynamicTool: Bool(true),
		Active:      Bool(true),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/test_types/789/") {
			t.Errorf("Expected /test_types/789/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.TestTypes.Update(context.Background(), 789, &TestType{
		Name:        Str("Updated Scan Type"),
		StaticTool:  Bool(true),
		DynamicTool: Bool(true),
		Active:      Bool(true),
	})
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestTestTypesService_PartialUpdate(t *testing.T) {
	response := `{
		"id": 321,
		"tags": ["sast"],
		"name": "Partially Updated Type",
		"static_tool": true,
		"dynamic_tool": false,
		"active": false
	}`

	expected := TestType{
		Id:          Int(321),
		Tags:        &[]string{"sast"},
		Name:        Str("Partially Updated Type"),
		StaticTool:  Bool(true),
		DynamicTool: Bool(false),
		Active:      Bool(false),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("Expected PATCH request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/test_types/321/") {
			t.Errorf("Expected /test_types/321/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.TestTypes.PartialUpdate(context.Background(), 321, &TestType{
		Active: Bool(false),
	})
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestTestTypesOptions_ToString(t *testing.T) {
	tests := []struct {
		name     string
		options  *TestTypesOptions
		expected string
	}{
		{
			name: "name only",
			options: &TestTypesOptions{
				Name: "ZAP",
			},
			expected: "?name=ZAP",
		},
		{
			name: "all fields",
			options: &TestTypesOptions{
				Limit:  10,
				Offset: 20,
				Name:   "ZAP",
			},
			expected: "?limit=10&offset=20&name=ZAP",
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
