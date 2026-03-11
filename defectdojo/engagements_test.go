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

func TestEngagementsService_List(t *testing.T) {
	response := `{
		"count": 1,
		"next": null,
		"previous": null,
		"results": [
			{
				"id": 1,
				"tags": ["quarterly"],
				"name": "Q1 2022 Assessment",
				"description": "Quarterly security assessment",
				"version": "1.0",
				"target_start": "2022-01-01",
				"target_end": "2022-03-31",
				"active": true,
				"status": "In Progress",
				"engagement_type": "Interactive",
				"lead": 1,
				"product": 1
			}
		]
	}`

	expected := Engagements{
		Count:    Int(1),
		Next:     nil,
		Previous: nil,
		Results: &[]Engagement{
			{
				Id:             Int(1),
				Tags:           &[]string{"quarterly"},
				Name:           Str("Q1 2022 Assessment"),
				Description:    Str("Quarterly security assessment"),
				Version:        Str("1.0"),
				TargetStart:    Str("2022-01-01"),
				TargetEnd:      Str("2022-03-31"),
				Active:         Bool(true),
				Status:         Str("In Progress"),
				EngagementType: Str("Interactive"),
				Lead:           Int(1),
				Product:        Int(1),
			},
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/engagements/") {
			t.Errorf("Expected /engagements/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.Engagements.List(context.Background(), &EngagementsOptions{
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

func TestEngagementsService_Read(t *testing.T) {
	response := `{
		"id": 123,
		"tags": ["pentest"],
		"name": "Annual Pentest",
		"description": "Annual penetration test",
		"version": "2.0",
		"target_start": "2022-06-01",
		"target_end": "2022-06-30",
		"active": true,
		"status": "In Progress",
		"engagement_type": "CI/CD",
		"lead": 2,
		"product": 3
	}`

	expected := Engagement{
		Id:             Int(123),
		Tags:           &[]string{"pentest"},
		Name:           Str("Annual Pentest"),
		Description:    Str("Annual penetration test"),
		Version:        Str("2.0"),
		TargetStart:    Str("2022-06-01"),
		TargetEnd:      Str("2022-06-30"),
		Active:         Bool(true),
		Status:         Str("In Progress"),
		EngagementType: Str("CI/CD"),
		Lead:           Int(2),
		Product:        Int(3),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/engagements/123/") {
			t.Errorf("Expected /engagements/123/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.Engagements.Read(context.Background(), 123)
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestEngagementsService_Create(t *testing.T) {
	response := `{
		"id": 456,
		"tags": ["cicd"],
		"name": "CI/CD Assessment",
		"description": "Automated CI/CD scan",
		"target_start": "2022-07-01",
		"target_end": "2022-07-31",
		"active": true,
		"status": "Not Started",
		"engagement_type": "CI/CD",
		"product": 1
	}`

	expected := Engagement{
		Id:             Int(456),
		Tags:           &[]string{"cicd"},
		Name:           Str("CI/CD Assessment"),
		Description:    Str("Automated CI/CD scan"),
		TargetStart:    Str("2022-07-01"),
		TargetEnd:      Str("2022-07-31"),
		Active:         Bool(true),
		Status:         Str("Not Started"),
		EngagementType: Str("CI/CD"),
		Product:        Int(1),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/engagements/") {
			t.Errorf("Expected /engagements/ in path, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.Engagements.Create(context.Background(), &Engagement{
		Name:           Str("CI/CD Assessment"),
		Description:    Str("Automated CI/CD scan"),
		TargetStart:    Str("2022-07-01"),
		TargetEnd:      Str("2022-07-31"),
		EngagementType: Str("CI/CD"),
		Product:        Int(1),
	})
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestEngagementsOptions_ToString(t *testing.T) {
	tests := []struct {
		name     string
		options  *EngagementsOptions
		expected string
	}{
		{
			name: "product only",
			options: &EngagementsOptions{
				Product: 5,
			},
			expected: "?product=5",
		},
		{
			name: "all fields",
			options: &EngagementsOptions{
				Limit:   10,
				Offset:  20,
				Product: 5,
				Name:    "test",
			},
			expected: "?limit=10&offset=20&product=5&name=test",
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
