package defectdojo

import (
	"time"
)

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(b bool) *bool { return &b }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(i int) *int { return &i }

// Str is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func Str(s string) *string { return &s }

// Date is a helper routine that allocates a new date value
// to store v and returns a pointer to it.
func Date(d time.Time) *time.Time { return &d }

func Slice(v []string) *[]string { return &v }
