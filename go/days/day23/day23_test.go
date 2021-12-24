package day23

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPath(t *testing.T) {
	var tests = []struct {
		start, end int
		points     []int
		cost       int
	}{
		{7, 0, []int{7, 1, 0}, 3},
		{11, 0, []int{11, 7, 1, 0}, 4},
		{19, 0, []int{19, 15, 11, 7, 1, 0}, 6},
		{8, 2, []int{8, 2}, 2},
		{8, 3, []int{8, 3}, 2},
		{8, 0, []int{8, 2, 1, 0}, 5},
		{8, 5, []int{8, 3, 4, 5}, 6},
		{8, 6, []int{8, 3, 4, 5, 6}, 7},
		{7, 8, []int{7, 2, 8}, 4},
		{8, 7, []int{8, 2, 7}, 4},
		{0, 7, []int{0, 1, 7}, 3},
		{0, 14, []int{0, 1, 2, 3, 4, 10, 14}, 10},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%d to %d", tt.start, tt.end)
		t.Run(testName, func(t *testing.T) {
			store := PathStore{map[[2]int]Path{}}
			result := store.Get(tt.start, tt.end)
			if !reflect.DeepEqual(result.points, tt.points) {
				t.Errorf("got points %v, want %v", result.points, tt.points)
			}
			if result.cost != tt.cost {
				t.Errorf("got cost %d, want %d", result.cost, tt.cost)
			}
		})
	}
}
