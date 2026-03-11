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

func TestFindingsService_List(t *testing.T) {
	response := `{
		"count": 1,
		"next": null,
		"previous": null,
		"results": [
			{
				"id": 1,
				"tags": ["critical"],
				"title": "SQL Injection in Login",
				"date": "2022-01-15",
				"cwe": 89,
				"severity": "Critical",
				"description": "SQL injection vulnerability found",
				"active": true,
				"verified": true,
				"false_p": false,
				"duplicate": false,
				"test": 1,
				"reporter": 1
			}
		]
	}`

	expected := Findings{
		Count:    Int(1),
		Next:     nil,
		Previous: nil,
		Results: &[]Finding{
			{
				Id:          Int(1),
				Tags:        &[]string{"critical"},
				Title:       Str("SQL Injection in Login"),
				Date:        Str("2022-01-15"),
				Cwe:         Int(89),
				Severity:    Str("Critical"),
				Description: Str("SQL injection vulnerability found"),
				Active:      Bool(true),
				Verified:    Bool(true),
				FalseP:      Bool(false),
				Duplicate:   Bool(false),
				Test:        Int(1),
				Reporter:    Int(1),
			},
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/findings/") {
			t.Errorf("Expected /findings/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.Findings.List(context.Background(), &FindingsOptions{
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

func TestFindingsService_Read(t *testing.T) {
	response := `{
		"id": 123,
		"tags": ["xss"],
		"title": "Reflected XSS",
		"date": "2022-02-01",
		"cwe": 79,
		"severity": "High",
		"description": "Reflected XSS in search parameter",
		"active": true,
		"verified": false,
		"false_p": false,
		"duplicate": false,
		"test": 5,
		"reporter": 2
	}`

	expected := Finding{
		Id:          Int(123),
		Tags:        &[]string{"xss"},
		Title:       Str("Reflected XSS"),
		Date:        Str("2022-02-01"),
		Cwe:         Int(79),
		Severity:    Str("High"),
		Description: Str("Reflected XSS in search parameter"),
		Active:      Bool(true),
		Verified:    Bool(false),
		FalseP:      Bool(false),
		Duplicate:   Bool(false),
		Test:        Int(5),
		Reporter:    Int(2),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/findings/123/") {
			t.Errorf("Expected /findings/123/ in path, got %s", r.URL.Path)
		}
		_, _ = fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	dj, _ := NewDojoClient(ts.URL, "token", nil)

	actual, err := dj.Findings.Read(context.Background(), 123)
	if !cmp.Equal(err, nil) {
		t.Errorf("error: %s", err)
	}

	if !cmp.Equal(actual, &expected) {
		t.Errorf("should have been equal, %+v, %+v", actual, &expected)
	}
}

func TestFindingsOptions_ToString(t *testing.T) {
	tests := []struct {
		name     string
		options  *FindingsOptions
		expected string
	}{
		{
			name: "severity only",
			options: &FindingsOptions{
				Severity: "Critical",
			},
			expected: "?severity=Critical",
		},
		{
			name: "all fields",
			options: &FindingsOptions{
				Limit:    10,
				Offset:   20,
				Title:    "SQL",
				Severity: "High",
				Active:   "true",
				Verified: "false",
				Prefetch: "test",
			},
			expected: "?limit=10&offset=20&title=SQL&severity=High&active=true&verified=false&prefetch=test",
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
