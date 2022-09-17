package test

import "testing"

// Test is a test
type Test struct {
	Name      string
	Mock      func()
	Validator func(t *testing.T, value interface{}, err error)
}
