package pool

import (
	"testing"
)

func TestInitPool(t *testing.T) {
	pool := InitPool(2)
	if pool == nil {
		t.Errorf("Expected pool to be initialized but got nil")
	}
	if pool.Cap() != 2 {
		t.Errorf("Expected pool cap to be 2 but got %v", pool.Cap())
	}
}
