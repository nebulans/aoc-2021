package integer

import (
	"fmt"
	"math/big"
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

func TestSumFunctions(t *testing.T) {
	var tests = []struct {
		a    []int
		want int
	}{
		{[]int{}, 0},
		{[]int{0}, 0},
		{[]int{2, -1, -1}, 0},
		{[]int{10}, 10},
		{[]int{2, 3, 10}, 15},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("sum(%v)==%d", tt.a, tt.want)
		t.Run(testName, func(t *testing.T) {
			ans := Sum(tt.a)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
		uints := make([]uint64, len(tt.a))
		for i, v := range tt.a {
			uints[i] = uint64(v)
		}
		t.Run(fmt.Sprintf("SumUint64(%v)==%d", uints, tt.want), func(t *testing.T) {
			ans := SumUint64(uints)
			if int(ans) != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
		bigints := make([]*big.Int, len(tt.a))
		for i, v := range tt.a {
			bigints[i] = big.NewInt(int64(v))
		}
		t.Run(fmt.Sprintf("SumBigInt(%v)==%d", bigints, tt.want), func(t *testing.T) {
			ans := SumBigInt(bigints)
			if int(ans.Int64()) != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
