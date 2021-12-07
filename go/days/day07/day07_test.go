package day07

import (
	"fmt"
	"testing"
)

func TestSumBelowCost(t *testing.T) {
	var tests = []struct {
		pos, target int
		cost        int
	}{
		{0, 0, 0},
		{0, 1, 1},
		{2, 0, 3},
		{3, 0, 6},
		{4, 0, 10},
		{0, 5, 15},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("Cost(%d,%d)==%d", tt.pos, tt.target, tt.cost)
		t.Run(testName, func(t *testing.T) {
			ans := sumBelowCost(tt.pos, tt.target)
			if ans != tt.cost {
				t.Errorf("got %d, want %d", ans, tt.cost)
			}
		})
	}
}
