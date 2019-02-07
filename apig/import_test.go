package apig

import "testing"

// TestFormatImportDir ...
func TestFormatImportDir(t *testing.T) {
	importPaths := generateImportSlice(
		"github.com/vharish836/api-server/db",
		"github.com/vharish836/api-server/models",
		"github.com/vharish836/api-server/server",
		"fmt",
	)

	result := formatImportDir(importPaths)
	if len(result) != 1 {
		t.Fatalf("Number of import dir is incorrect. expected: 1, actual: %d", len(result))
	}

	expect := "github.com/vharish836/api-server"
	if result[0] != expect {
		t.Fatalf("Incorrect import dir. expected: %s, actual: %s", expect, result[0])
	}
}

func generateImportSlice(paths ...string) []string {
	var importPaths []string
	for _, path := range paths {
		importPaths = append(importPaths, path)
	}
	return importPaths
}
