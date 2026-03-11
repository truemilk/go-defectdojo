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

func TestToolTypesService_List(t *testing.T) {
	response := `{
		"count": 1,
		"next": null,
		"previous": null,
		"results": [
			{
				"id": 1,
				"name": "Burp Suite",
				"description": "Web vulnerability scanner"
			}
		]
	}`

	expected := ToolTypes{
		Count:    Int(1),
		Next:     nil,
		Previous: nil,
		Results: &[]ToolType{
			{
				Id:          Int(1),
				Name:        Str("Burp Suite"),
				Description: Str("Web vulnerability scanner"),
			},
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/tool_types/") {
			t.Errorf("Expected /tool_types/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.ToolTypes.List(context.Background(), &ToolTypesOptions{
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

func TestToolTypesService_Read(t *testing.T) {
	response := `{
		"id": 123,
		"name": "Burp Suite",
		"description": "Web vulnerability scanner"
	}`

	expected := ToolType{
		Id:          Int(123),
		Name:        Str("Burp Suite"),
		Description: Str("Web vulnerability scanner"),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/tool_types/123/") {
			t.Errorf("Expected /tool_types/123/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.ToolTypes.Read(context.Background(), 123)
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestToolTypesService_Create(t *testing.T) {
	response := `{
		"id": 456,
		"name": "New Scanner",
		"description": "A new scanning tool"
	}`

	expected := ToolType{
		Id:          Int(456),
		Name:        Str("New Scanner"),
		Description: Str("A new scanning tool"),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/tool_types/") {
			t.Errorf("Expected /tool_types/ in path, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.ToolTypes.Create(context.Background(), &ToolType{
		Name:        Str("New Scanner"),
		Description: Str("A new scanning tool"),
	})
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestToolTypesService_Update(t *testing.T) {
	response := `{
		"id": 789,
		"name": "Updated Scanner",
		"description": "An updated scanning tool"
	}`

	expected := ToolType{
		Id:          Int(789),
		Name:        Str("Updated Scanner"),
		Description: Str("An updated scanning tool"),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/tool_types/789/") {
			t.Errorf("Expected /tool_types/789/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.ToolTypes.Update(context.Background(), 789, &ToolType{
		Name:        Str("Updated Scanner"),
		Description: Str("An updated scanning tool"),
	})
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestToolTypesService_PartialUpdate(t *testing.T) {
	response := `{
		"id": 321,
		"name": "Partially Updated Scanner",
		"description": "Original description"
	}`

	expected := ToolType{
		Id:          Int(321),
		Name:        Str("Partially Updated Scanner"),
		Description: Str("Original description"),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("Expected PATCH request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/tool_types/321/") {
			t.Errorf("Expected /tool_types/321/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.ToolTypes.PartialUpdate(context.Background(), 321, &ToolType{
		Name: Str("Partially Updated Scanner"),
	})
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestToolTypesService_Delete(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/tool_types/654/") {
			t.Errorf("Expected /tool_types/654/ in path, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintln(w, "{}")
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.ToolTypes.Delete(context.Background(), 654)
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if actual == nil {
		t.Errorf("expected non-nil response")
	}
}

func TestToolTypesOptions_ToString(t *testing.T) {
	tests := []struct {
		name     string
		options  *ToolTypesOptions
		expected string
	}{
		{
			name: "name only",
			options: &ToolTypesOptions{
				Name: "Burp",
			},
			expected: "?name=Burp",
		},
		{
			name: "all fields",
			options: &ToolTypesOptions{
				Limit:       10,
				Offset:      20,
				ID:          5,
				Name:        "Burp",
				Description: "scanner",
			},
			expected: "?limit=10&offset=20&id=5&name=Burp&description=scanner",
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
