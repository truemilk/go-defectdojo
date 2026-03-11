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

func TestTestsService_List(t *testing.T) {
	response := `{
		"count": 1,
		"next": null,
		"previous": null,
		"results": [
			{
				"id": 1,
				"engagement": 1,
				"tags": ["automated"],
				"scan_type": "ZAP Scan",
				"title": "Weekly ZAP Scan",
				"description": "Automated weekly scan",
				"target_start": "2022-01-01",
				"target_end": "2022-01-02",
				"percent_complete": 100,
				"lead": 1,
				"test_type": 1
			}
		]
	}`

	expected := Tests{
		Count:    Int(1),
		Next:     nil,
		Previous: nil,
		Results: &[]Test{
			{
				Id:              Int(1),
				EngagementId:    Int(1),
				Tags:            &[]string{"automated"},
				ScanType:        Str("ZAP Scan"),
				Title:           Str("Weekly ZAP Scan"),
				Description:     Str("Automated weekly scan"),
				TargetStart:     Str("2022-01-01"),
				TargetEnd:       Str("2022-01-02"),
				PercentComplete: Int(100),
				Lead:            Int(1),
				TestType:        Int(1),
			},
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/tests/") {
			t.Errorf("Expected /tests/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.Tests.List(context.Background(), &TestsOptions{
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

func TestTestsService_Read(t *testing.T) {
	response := `{
		"id": 123,
		"engagement": 5,
		"tags": ["manual"],
		"scan_type": "Manual Code Review",
		"title": "Code Review Q1",
		"description": "Manual code review",
		"target_start": "2022-03-01",
		"target_end": "2022-03-15",
		"percent_complete": 50,
		"lead": 2,
		"test_type": 3
	}`

	expected := Test{
		Id:              Int(123),
		EngagementId:    Int(5),
		Tags:            &[]string{"manual"},
		ScanType:        Str("Manual Code Review"),
		Title:           Str("Code Review Q1"),
		Description:     Str("Manual code review"),
		TargetStart:     Str("2022-03-01"),
		TargetEnd:       Str("2022-03-15"),
		PercentComplete: Int(50),
		Lead:            Int(2),
		TestType:        Int(3),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/tests/123/") {
			t.Errorf("Expected /tests/123/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.Tests.Read(context.Background(), 123)
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestTestsService_Create(t *testing.T) {
	response := `{
		"id": 456,
		"engagement": 1,
		"tags": ["new"],
		"scan_type": "SonarQube Scan",
		"title": "New Scan",
		"description": "New automated scan",
		"target_start": "2022-04-01",
		"target_end": "2022-04-02",
		"lead": 1,
		"test_type": 2
	}`

	expected := Test{
		Id:           Int(456),
		EngagementId: Int(1),
		Tags:         &[]string{"new"},
		ScanType:     Str("SonarQube Scan"),
		Title:        Str("New Scan"),
		Description:  Str("New automated scan"),
		TargetStart:  Str("2022-04-01"),
		TargetEnd:    Str("2022-04-02"),
		Lead:         Int(1),
		TestType:     Int(2),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/tests/") {
			t.Errorf("Expected /tests/ in path, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.Tests.Create(context.Background(), &Test{
		EngagementId: Int(1),
		ScanType:     Str("SonarQube Scan"),
		Title:        Str("New Scan"),
		TargetStart:  Str("2022-04-01"),
		TargetEnd:    Str("2022-04-02"),
		TestType:     Int(2),
	})
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestTestsService_Update(t *testing.T) {
	response := `{
		"id": 789,
		"engagement": 1,
		"tags": ["updated"],
		"scan_type": "ZAP Scan",
		"title": "Updated Scan",
		"description": "Updated scan description",
		"target_start": "2022-05-01",
		"target_end": "2022-05-15",
		"percent_complete": 75,
		"lead": 1,
		"test_type": 1
	}`

	expected := Test{
		Id:              Int(789),
		EngagementId:    Int(1),
		Tags:            &[]string{"updated"},
		ScanType:        Str("ZAP Scan"),
		Title:           Str("Updated Scan"),
		Description:     Str("Updated scan description"),
		TargetStart:     Str("2022-05-01"),
		TargetEnd:       Str("2022-05-15"),
		PercentComplete: Int(75),
		Lead:            Int(1),
		TestType:        Int(1),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/tests/789/") {
			t.Errorf("Expected /tests/789/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.Tests.Update(context.Background(), 789, &Test{
		EngagementId: Int(1),
		ScanType:     Str("ZAP Scan"),
		Title:        Str("Updated Scan"),
		Description:  Str("Updated scan description"),
		TargetStart:  Str("2022-05-01"),
		TargetEnd:    Str("2022-05-15"),
		TestType:     Int(1),
	})
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestTestsService_PartialUpdate(t *testing.T) {
	response := `{
		"id": 321,
		"engagement": 1,
		"scan_type": "ZAP Scan",
		"title": "Partially Updated Scan",
		"target_start": "2022-05-01",
		"target_end": "2022-05-15",
		"test_type": 1
	}`

	expected := Test{
		Id:           Int(321),
		EngagementId: Int(1),
		ScanType:     Str("ZAP Scan"),
		Title:        Str("Partially Updated Scan"),
		TargetStart:  Str("2022-05-01"),
		TargetEnd:    Str("2022-05-15"),
		TestType:     Int(1),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("Expected PATCH request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/tests/321/") {
			t.Errorf("Expected /tests/321/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.Tests.PartialUpdate(context.Background(), 321, &Test{
		Title: Str("Partially Updated Scan"),
	})
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestTestsOptions_ToString(t *testing.T) {
	tests := []struct {
		name     string
		options  *TestsOptions
		expected string
	}{
		{
			name: "title only",
			options: &TestsOptions{
				Title: "ZAP",
			},
			expected: "?title=ZAP",
		},
		{
			name: "all fields",
			options: &TestsOptions{
				Limit:      10,
				Offset:     20,
				Title:      "ZAP",
				Engagement: 5,
				TestType:   3,
			},
			expected: "?limit=10&offset=20&title=ZAP&test_type=3&engagement=5",
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
