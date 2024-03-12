package calc

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	patterns := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 1, 3},
		{2, 2, 4},
	}

	for idx, pattern := range patterns {
		actual := Add(pattern.a, pattern.b)
		if actual != pattern.expected {
			t.Errorf("pattern %d: expected %d, got %d", idx, pattern.expected, actual)
		} else {
			log.Printf("pattern %d: expected %d, got %d", idx, pattern.expected, actual)
		}
	}
}
