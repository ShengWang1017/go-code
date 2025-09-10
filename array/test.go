package array

import "testing"

func lc53Test(t *testing.T) {
	result := MaxSubArray([]int{1})
	if result != 0 {
		t.Errorf("expected 0, got %d", result)
	}
}
