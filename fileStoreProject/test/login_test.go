package test

import (
	"file/ui"
	"testing"
)

// Test for IsValidPassword function
func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{
		{"P@ssw0rd", true},     // Valid password
		{"password", false},    // Missing uppercase, digit, special character
		{"PASSWORD123", false}, // Missing lowercase, special character
		{"Pass123", false},     // Missing special character
		{"P@ssw0", false},      // Too short
		{"P@ssw0rd!", true},    // Valid password with special character at end
		{"P@ss1", false},       // Too short
		{"P@SSW0RD", false},    // Missing lowercase
	}

	for _, test := range tests {
		result := ui.IsValidPassword(test.password)
		if result != test.expected {
			t.Errorf("IsValidPassword(%q) = %v; want %v", test.password, result, test.expected)
		}
	}
}

// Test for isValidMobile function
func TestIsValidMobile(t *testing.T) {
	tests := []struct {
		mobileNumber string
		expected     bool
	}{
		{"1234567890", true},     // Valid phone number
		{"123456789", false},     // Too short
		{"12345678901", false},   // Too long
		{"abcdefghij", false},    // Non-numeric
		{"1234abc890", false},    // Contains letters
		{"1234567890abc", false}, // Contains extra characters
	}

	for _, test := range tests {
		result := ui.IsValidMobile(test.mobileNumber)
		if result != test.expected {
			t.Errorf("isValidMobile(%q) = %v; want %v", test.mobileNumber, result, test.expected)
		}
	}
}
