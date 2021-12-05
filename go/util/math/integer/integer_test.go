package integer

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 0, 0},
		{1, 5, 5},
		{1, -1, 1},
		{0, -1, 0},
		{-2, -3, -2},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("Max(%d,%d)==%d", tt.a, tt.b, tt.want)
		t.Run(testName, func(t *testing.T) {
			ans := Max(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	var tests = []struct {
		a    int
		want int
	}{
		{0, 0},
		{1, 1},
		{-1, 1},
		{10, 10},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("Abs(%d)==%d", tt.a, tt.want)
		t.Run(testName, func(t *testing.T) {
			ans := Abs(tt.a)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestUnitStep(t *testing.T) {
	var tests = []struct {
		a    int
		want int
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{10, 1},
		{-5, -1},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("UnitStep(%d)==%d", tt.a, tt.want)
		t.Run(testName, func(t *testing.T) {
			ans := UnitStep(tt.a)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
