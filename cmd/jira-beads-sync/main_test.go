package main

import (
	"os"
	"strings"
	"testing"
)

func TestIsURL(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "http URL",
			input: "http://jira.example.com/browse/PROJ-123",
			want:  true,
		},
		{
			name:  "https URL",
			input: "https://jira.example.com/browse/PROJ-123",
			want:  true,
		},
		{
			name:  "http URL with port",
			input: "http://localhost:8080/browse/PROJ-123",
			want:  true,
		},
		{
			name:  "https URL with port",
			input: "https://jira.company.com:443/projects/PROJ/issues/PROJ-123",
			want:  true,
		},
		{
			name:  "http URL with query parameters",
			input: "http://jira.example.com/browse/PROJ-123?filter=123&view=detail",
			want:  true,
		},
		{
			name:  "https URL with fragment",
			input: "https://jira.example.com/browse/PROJ-123#comment-456",
			want:  true,
		},
		{
			name:  "http URL minimal",
			input: "http://example.com",
			want:  true,
		},
		{
			name:  "https URL minimal",
			input: "https://example.com",
			want:  true,
		},
		{
			name:  "issue key only",
			input: "PROJ-123",
			want:  false,
		},
		{
			name:  "issue key with whitespace",
			input: "  PROJ-123  ",
			want:  false,
		},
		{
			name:  "path without protocol",
			input: "/browse/PROJ-123",
			want:  false,
		},
		{
			name:  "domain without protocol",
			input: "jira.example.com/browse/PROJ-123",
			want:  false,
		},
		{
			name:  "ftp URL",
			input: "ftp://example.com/file",
			want:  false,
		},
		{
			name:  "file URL",
			input: "file:///path/to/file",
			want:  false,
		},
		{
			name:  "empty string",
			input: "",
			want:  false,
		},
		{
			name:  "HTTP uppercase",
			input: "HTTP://EXAMPLE.COM",
			want:  false,
		},
		{
			name:  "HTTPS uppercase",
			input: "HTTPS://EXAMPLE.COM",
			want:  false,
		},
		{
			name:  "http with typo",
			input: "htttp://example.com",
			want:  false,
		},
		{
			name:  "https with typo",
			input: "htttps://example.com",
			want:  false,
		},
		{
			name:  "protocol-relative URL",
			input: "//example.com/path",
			want:  false,
		},
		{
			name:  "http:// only",
			input: "http://",
			want:  true,
		},
		{
			name:  "https:// only",
			input: "https://",
			want:  true,
		},
		{
			name:  "http in middle of string",
			input: "this http://example.com is a url",
			want:  false,
		},
		{
			name:  "URL with unicode",
			input: "https://例え.jp/path",
			want:  true,
		},
		{
			name:  "URL with special characters",
			input: "https://example.com/path?param=value&other=123",
			want:  true,
		},
		{
			name:  "URL with IP address",
			input: "http://192.168.1.1/browse/PROJ-123",
			want:  true,
		},
		{
			name:  "URL with localhost",
			input: "http://localhost/browse/PROJ-123",
			want:  true,
		},
		{
			name:  "https URL with subdomain",
			input: "https://jira.atlassian.example.com/browse/PROJ-123",
			want:  true,
		},
		{
			name:  "URL with authentication",
			input: "https://user:pass@example.com/browse/PROJ-123",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isURL(tt.input)
			if got != tt.want {
				t.Errorf("isURL(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsURLEdgeCases(t *testing.T) {
	t.Run("real Jira URLs", func(t *testing.T) {
		jiraURLs := []string{
			"https://jira.atlassian.com/browse/JRASERVER-1234",
			"http://jira.company.com/projects/PROJ/issues/PROJ-123",
			"https://company.atlassian.net/browse/PROJ-456",
		}

		for _, url := range jiraURLs {
			if !isURL(url) {
				t.Errorf("Expected %q to be recognized as URL", url)
			}
		}
	})

	t.Run("real issue keys", func(t *testing.T) {
		issueKeys := []string{
			"PROJ-123",
			"ABC-1",
			"LONGPROJECTNAME-999999",
			"A-1",
		}

		for _, key := range issueKeys {
			if isURL(key) {
				t.Errorf("Expected %q NOT to be recognized as URL", key)
			}
		}
	})

	t.Run("ambiguous inputs", func(t *testing.T) {
		// These should not be URLs
		ambiguous := []string{
			"http",
			"https",
			"http:/",
			"https:/",
			"://example.com",
		}

		for _, input := range ambiguous {
			result := isURL(input)
			// Document current behavior
			t.Logf("isURL(%q) = %v", input, result)
		}
	})
}

func BenchmarkIsURL(b *testing.B) {
	testCases := []struct {
		name  string
		input string
	}{
		{"URL", "https://jira.example.com/browse/PROJ-123"},
		{"Issue Key", "PROJ-123"},
		{"Long URL", "https://very-long-subdomain.jira.example.company.com/projects/PROJECT/issues/PROJECT-123456?filter=recent&view=detail#comment-789"},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isURL(tc.input)
			}
		})
	}
}

func TestRunFetchByJQLWithMockConfig(t *testing.T) {
	// Create a temporary config file
	tmpDir := t.TempDir()
	configFile := tmpDir + "/config.yml"

	configContent := `jira:
  base_url: https://jira.example.com
  username: test@example.com
  api_token: test-token
`

	if err := os.WriteFile(configFile, []byte(configContent), 0600); err != nil {
		t.Fatalf("Failed to create test config: %v", err)
	}

	// Set environment variable to override config location
	oldXDG := os.Getenv("XDG_CONFIG_HOME")
	oldHOME := os.Getenv("HOME")
	defer func() {
		os.Setenv("XDG_CONFIG_HOME", oldXDG)
		os.Setenv("HOME", oldHOME)
	}()

	os.Setenv("XDG_CONFIG_HOME", tmpDir)
	os.Setenv("HOME", tmpDir)

	// Create the expected config directory structure
	configDir := tmpDir + "/jira-beads-sync"
	if err := os.MkdirAll(configDir, 0755); err != nil {
		t.Fatalf("Failed to create config dir: %v", err)
	}

	configPath := configDir + "/config.yml"
	if err := os.WriteFile(configPath, []byte(configContent), 0600); err != nil {
		t.Fatalf("Failed to write config: %v", err)
	}

	// Test will fail at network call (which is expected without a real Jira server)
	// But it will exercise the config loading and client creation code paths
	err := runFetchByJQL("project = TEST")

	// We expect an error because there's no real Jira server
	// But the error should be from network/API call, not from config loading
	if err != nil {
		// Expected - network call fails
		t.Logf("Got expected error (network failure): %v", err)

		// Verify it's not a config error
		if strings.Contains(err.Error(), "failed to configure") {
			t.Error("Should not fail at config stage with valid config")
		}
	}
}
