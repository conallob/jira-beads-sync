package converter

import "testing"

func TestContains(t *testing.T) {
	tests := []struct {
		name  string
		slice []string
		value string
		want  bool
	}{
		{
			name:  "found in single element slice",
			slice: []string{"apple"},
			value: "apple",
			want:  true,
		},
		{
			name:  "not found in single element slice",
			slice: []string{"apple"},
			value: "orange",
			want:  false,
		},
		{
			name:  "found at beginning of slice",
			slice: []string{"first", "second", "third"},
			value: "first",
			want:  true,
		},
		{
			name:  "found in middle of slice",
			slice: []string{"first", "second", "third"},
			value: "second",
			want:  true,
		},
		{
			name:  "found at end of slice",
			slice: []string{"first", "second", "third"},
			value: "third",
			want:  true,
		},
		{
			name:  "not found in multi-element slice",
			slice: []string{"first", "second", "third"},
			value: "fourth",
			want:  false,
		},
		{
			name:  "empty slice",
			slice: []string{},
			value: "anything",
			want:  false,
		},
		{
			name:  "nil slice",
			slice: nil,
			value: "anything",
			want:  false,
		},
		{
			name:  "empty string in slice",
			slice: []string{"", "non-empty"},
			value: "",
			want:  true,
		},
		{
			name:  "empty string not in slice",
			slice: []string{"first", "second"},
			value: "",
			want:  false,
		},
		{
			name:  "case sensitive - lowercase vs uppercase",
			slice: []string{"Apple", "Orange"},
			value: "apple",
			want:  false,
		},
		{
			name:  "exact match required",
			slice: []string{"test", "testing"},
			value: "test",
			want:  true,
		},
		{
			name:  "substring does not match",
			slice: []string{"testing"},
			value: "test",
			want:  false,
		},
		{
			name:  "duplicate values in slice - first occurrence",
			slice: []string{"duplicate", "value", "duplicate"},
			value: "duplicate",
			want:  true,
		},
		{
			name:  "whitespace matters",
			slice: []string{"value", "value ", " value"},
			value: "value",
			want:  true,
		},
		{
			name:  "whitespace no match",
			slice: []string{"value ", " value"},
			value: "value",
			want:  false,
		},
		{
			name:  "special characters",
			slice: []string{"issue-123", "issue_456", "issue.789"},
			value: "issue-123",
			want:  true,
		},
		{
			name:  "unicode characters",
			slice: []string{"hello", "世界", "مرحبا"},
			value: "世界",
			want:  true,
		},
		{
			name:  "large slice with match at end",
			slice: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "target"},
			value: "target",
			want:  true,
		},
		{
			name:  "large slice without match",
			slice: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
			value: "target",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := contains(tt.slice, tt.value)
			if got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsRealWorldScenarios(t *testing.T) {
	t.Run("Jira issue keys", func(t *testing.T) {
		visited := []string{"PROJ-123", "PROJ-124", "PROJ-125"}

		if !contains(visited, "PROJ-123") {
			t.Error("Should find PROJ-123 in visited issues")
		}

		if contains(visited, "PROJ-126") {
			t.Error("Should not find PROJ-126 in visited issues")
		}
	})

	t.Run("dependency tracking", func(t *testing.T) {
		dependencies := []string{"auth-service", "database", "cache"}

		if !contains(dependencies, "database") {
			t.Error("Should find database in dependencies")
		}

		if contains(dependencies, "logging") {
			t.Error("Should not find logging in dependencies")
		}
	})

	t.Run("duplicate prevention", func(t *testing.T) {
		// Simulating deduplication logic
		seen := []string{}
		items := []string{"item1", "item2", "item1", "item3", "item2"}

		var unique []string
		for _, item := range items {
			if !contains(seen, item) {
				seen = append(seen, item)
				unique = append(unique, item)
			}
		}

		expectedUnique := []string{"item1", "item2", "item3"}
		if len(unique) != len(expectedUnique) {
			t.Errorf("Expected %d unique items, got %d", len(expectedUnique), len(unique))
		}

		for i, item := range unique {
			if item != expectedUnique[i] {
				t.Errorf("At index %d: expected %s, got %s", i, expectedUnique[i], item)
			}
		}
	})
}

func BenchmarkContains(b *testing.B) {
	// Benchmark for small slice
	b.Run("small slice (10 elements)", func(b *testing.B) {
		slice := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
		value := "j" // Worst case - at end
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			contains(slice, value)
		}
	})

	// Benchmark for medium slice
	b.Run("medium slice (100 elements)", func(b *testing.B) {
		slice := make([]string, 100)
		for i := 0; i < 100; i++ {
			slice[i] = string(rune('a' + i%26))
		}
		value := slice[99] // Worst case - at end
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			contains(slice, value)
		}
	})

	// Benchmark for large slice
	b.Run("large slice (1000 elements)", func(b *testing.B) {
		slice := make([]string, 1000)
		for i := 0; i < 1000; i++ {
			slice[i] = string(rune('a' + i%26))
		}
		value := slice[999] // Worst case - at end
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			contains(slice, value)
		}
	})
}
