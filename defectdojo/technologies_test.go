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

func TestTechnologiesService_List(t *testing.T) {
	response := `{
		"count": 1,
		"next": null,
		"previous": null,
		"results": [
			{
				"id": 1,
				"tags": ["python"],
				"name": "Django",
				"confidence": 90,
				"version": "4.2",
				"icon": "",
				"website": "https://djangoproject.com",
				"website_found": "https://example.com",
				"created": "2022-01-01T12:00:00Z",
				"product": 1,
				"user": 1
			}
		]
	}`

	expected := Technologies{
		Count:    Int(1),
		Next:     nil,
		Previous: nil,
		Results: &[]Technology{
			{
				Id:           Int(1),
				Tags:         &[]string{"python"},
				Name:         Str("Django"),
				Confidence:   Int(90),
				Version:      Str("4.2"),
				Icon:         Str(""),
				Website:      Str("https://djangoproject.com"),
				WebsiteFound: Str("https://example.com"),
				Created:      Date(time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC)),
				Product:      Int(1),
				User:         Int(1),
			},
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/technologies/") {
			t.Errorf("Expected /technologies/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.Technologies.List(context.Background(), &TechnologiesOptions{
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

func TestTechnologiesService_Read(t *testing.T) {
	response := `{
		"id": 123,
		"tags": ["python"],
		"name": "Django",
		"confidence": 90,
		"version": "4.2",
		"icon": "",
		"website": "https://djangoproject.com",
		"website_found": "https://example.com",
		"created": "2022-01-01T12:00:00Z",
		"product": 1,
		"user": 1
	}`

	expected := Technology{
		Id:           Int(123),
		Tags:         &[]string{"python"},
		Name:         Str("Django"),
		Confidence:   Int(90),
		Version:      Str("4.2"),
		Icon:         Str(""),
		Website:      Str("https://djangoproject.com"),
		WebsiteFound: Str("https://example.com"),
		Created:      Date(time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC)),
		Product:      Int(1),
		User:         Int(1),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/technologies/123/") {
			t.Errorf("Expected /technologies/123/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.Technologies.Read(context.Background(), 123)
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestTechnologiesService_Create(t *testing.T) {
	response := `{
		"id": 456,
		"tags": ["go"],
		"name": "Go",
		"confidence": 100,
		"version": "1.21",
		"icon": "",
		"website": "https://go.dev",
		"website_found": "https://example.com",
		"created": "2022-01-01T12:00:00Z",
		"product": 1,
		"user": 1
	}`

	expected := Technology{
		Id:           Int(456),
		Tags:         &[]string{"go"},
		Name:         Str("Go"),
		Confidence:   Int(100),
		Version:      Str("1.21"),
		Icon:         Str(""),
		Website:      Str("https://go.dev"),
		WebsiteFound: Str("https://example.com"),
		Created:      Date(time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC)),
		Product:      Int(1),
		User:         Int(1),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/technologies/") {
			t.Errorf("Expected /technologies/ in path, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.Technologies.Create(context.Background(), &Technology{
		Name:       Str("Go"),
		Confidence: Int(100),
		Version:    Str("1.21"),
		Product:    Int(1),
	})
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestTechnologiesService_Delete(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/technologies/789/") {
			t.Errorf("Expected /technologies/789/ in path, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintln(w, "{}")
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.Technologies.Delete(context.Background(), 789)
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if actual == nil {
		t.Errorf("expected non-nil response")
	}
}

func TestTechnologiesOptions_ToString(t *testing.T) {
	tests := []struct {
		name     string
		options  *TechnologiesOptions
		expected string
	}{
		{
			name: "limit only",
			options: &TechnologiesOptions{
				Limit: 10,
			},
			expected: "?limit=10",
		},
		{
			name: "all fields",
			options: &TechnologiesOptions{
				Limit:  10,
				Offset: 20,
			},
			expected: "?limit=10&offset=20",
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
