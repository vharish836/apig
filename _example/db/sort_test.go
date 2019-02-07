package db

import "testing"

// TestConvertPrefixToQueryPlus ...
func TestConvertPrefixToQueryPlus(t *testing.T) {
	value := convertPrefixToQuery("id")

	if value != "id asc" {
		t.Fatalf("Expected: `id asc`, actual: %s", value)
	}

	value = convertPrefixToQuery(" id")

	if value != "id asc" {
		t.Fatalf("Expected: `id asc`, actual: %s", value)
	}
}

// TestConvertPrefixToQueryMinus ...
func TestConvertPrefixToQueryMinus(t *testing.T) {
	value := convertPrefixToQuery("-id")

	if value != "id desc" {
		t.Fatalf("Expected: `id desc`, actual: %s", value)
	}
}
