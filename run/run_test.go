package run

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	manifestPath := filepath.Join("testdata", "manifest.k")
	brewfilePath := filepath.Join("testdata", "Brewfile")
	checkerPath := filepath.Join("testdata", "version_checker.sh")

	err := os.MkdirAll("testdata", os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create testdata directory: %v", err)
	}
	defer os.RemoveAll("testdata")

	err = copyFile("../testdata/manifest.k", manifestPath)
	if err != nil {
		t.Fatalf("Failed to copy manifest file: %v", err)
	}

	Brewfile(manifestPath, brewfilePath, checkerPath)

	if _, err := os.Stat(brewfilePath); os.IsNotExist(err) {
		t.Errorf("Expected Brewfile to be generated, but it was not found")
	}

	if _, err := os.Stat(checkerPath); os.IsNotExist(err) {
		t.Errorf("Expected version_checker.sh to be generated, but it was not found")
	}
}

func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0o644)
}
