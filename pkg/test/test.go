package test

import "testing"

// Test is a test
type Test struct {
	Name      string
	Mock      func()
	Validator func(t *testing.T, value interface{}, err error)
}

// DefaultAPIKey returns the default testing api key
func DefaultAPIKey() string {
	return "supersecret"
}

// DefaultSecret returns the default testing secret
func DefaultSecret() string {
	return "supersecret"
}
