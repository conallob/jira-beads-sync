package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestVersionDefaults(t *testing.T) {
	// Test that default values are set correctly when not using ldflags
	if version != "dev" {
		t.Errorf("expected default version to be 'dev', got %q", version)
	}
	if commit != "none" {
		t.Errorf("expected default commit to be 'none', got %q", commit)
	}
	if date != "unknown" {
		t.Errorf("expected default date to be 'unknown', got %q", date)
	}
}

func TestVersionLdflags(t *testing.T) {
	// Build binary with custom ldflags and verify output
	tempDir := t.TempDir()
	binaryPath := filepath.Join(tempDir, "jira-beads-sync-test")

	testVersion := "v1.2.3"
	testCommit := "abc123def"
	testDate := "2024-01-15T10:30:00Z"

	ldflags := strings.Join([]string{
		"-X main.version=" + testVersion,
		"-X main.commit=" + testCommit,
		"-X main.date=" + testDate,
	}, " ")

	// Get the directory containing this test file
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("failed to get current file path")
	}
	pkgDir := filepath.Dir(filename)

	// Build the binary with ldflags
	buildCmd := exec.Command("go", "build", "-ldflags", ldflags, "-o", binaryPath, ".")
	buildCmd.Dir = pkgDir
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("failed to build binary with ldflags: %v", err)
	}

	// Run the version command
	runCmd := exec.Command(binaryPath, "version")
	output, err := runCmd.Output()
	if err != nil {
		t.Fatalf("failed to run version command: %v", err)
	}

	outputStr := string(output)

	// Verify version is in output
	if !strings.Contains(outputStr, testVersion) {
		t.Errorf("expected output to contain version %q, got:\n%s", testVersion, outputStr)
	}

	// Verify commit is in output
	if !strings.Contains(outputStr, testCommit) {
		t.Errorf("expected output to contain commit %q, got:\n%s", testCommit, outputStr)
	}

	// Verify date is in output
	if !strings.Contains(outputStr, testDate) {
		t.Errorf("expected output to contain date %q, got:\n%s", testDate, outputStr)
	}
}

func TestVersionOutputFormat(t *testing.T) {
	// Build binary with ldflags and verify output format
	tempDir := t.TempDir()
	binaryPath := filepath.Join(tempDir, "jira-beads-sync-test")

	testVersion := "v2.0.0"
	testCommit := "deadbeef"
	testDate := "2024-06-01T00:00:00Z"

	ldflags := strings.Join([]string{
		"-X main.version=" + testVersion,
		"-X main.commit=" + testCommit,
		"-X main.date=" + testDate,
	}, " ")

	// Get the directory containing this test file
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("failed to get current file path")
	}
	pkgDir := filepath.Dir(filename)

	// Build the binary with ldflags
	buildCmd := exec.Command("go", "build", "-ldflags", ldflags, "-o", binaryPath, ".")
	buildCmd.Dir = pkgDir
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("failed to build binary with ldflags: %v", err)
	}

	// Run the version command
	runCmd := exec.Command(binaryPath, "version")
	output, err := runCmd.Output()
	if err != nil {
		t.Fatalf("failed to run version command: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines of output, got %d:\n%s", len(lines), output)
	}

	// Verify first line format: "jira-beads-sync <version>"
	expectedFirstLine := "jira-beads-sync " + testVersion
	if lines[0] != expectedFirstLine {
		t.Errorf("expected first line %q, got %q", expectedFirstLine, lines[0])
	}

	// Verify second line contains commit
	if !strings.Contains(lines[1], "commit:") || !strings.Contains(lines[1], testCommit) {
		t.Errorf("expected second line to contain 'commit: %s', got %q", testCommit, lines[1])
	}

	// Verify third line contains date
	if !strings.Contains(lines[2], "built:") || !strings.Contains(lines[2], testDate) {
		t.Errorf("expected third line to contain 'built: %s', got %q", testDate, lines[2])
	}
}
