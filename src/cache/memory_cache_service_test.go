package cache

import (
	"testing"
)

func TestMemoryCacheService(t *testing.T) {
	testCases := []struct {
		key         string
		value       int
		expected    bool
		expectedErr error
	}{
		{"key1", 10, true, nil}, // Cache hit
		{"key2", 20, true, nil}, // Cache hit
	}

	cache := NewMemoryCacheService()

	// Test Set and Get methods
	for _, tc := range testCases {
		err := cache.Set(tc.key, tc.value)
		if err != nil {
			t.Errorf("Unexpected error while setting cache value: %v", err)
		}

		value, found, err := cache.Get(tc.key)
		if err != nil {
			t.Errorf("Unexpected error while getting cache value: %v", err)
		}

		if found != tc.expected {
			t.Errorf("Expected cache hit for key '%s', but cache miss occurred", tc.key)
		}

		if value != tc.value {
			t.Errorf("For key '%s', expected cache value %d, but got %d", tc.key, tc.value, value)
		}
	}

	// Test Get method for non-existing key
	_, found, err := cache.Get("non-existing-key")
	if err != nil {
		t.Errorf("Unexpected error while getting cache value: %v", err)
	}

	if found {
		t.Error("Expected cache miss, but cache hit occurred")
	}
}
