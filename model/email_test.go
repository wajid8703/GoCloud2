package model

import (
	is2 "github.com/matryer/is"
	"testing"
)

func TestEmail_IsValid(t *testing.T) {
	tests := []struct {
		address string
		valid   bool
	}{
		{"me@axample.com", true},
		{"@example.com", false},
		{"me@", false},
		{"@", false},
		{"", false},
	}
	t.Run("reports valid email addresses", func(t *testing.T) {
		for _, test := range tests {
			t.Run(test.address, func(t *testing.T) {
				is := is2.New(t)
				e := Email(test.address)
				is.Equal(test.valid, e.IsValid())
			})
		}
	})
}
