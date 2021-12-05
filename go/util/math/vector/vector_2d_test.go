package vector

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	var tests = []struct {
		a, b Vec2
		want Vec2
	}{
		{Vec2{0, 0}, Vec2{1, 1}, Vec2{1, 1}},
		{Vec2{1, 0}, Vec2{0, 1}, Vec2{1, 1}},
		{Vec2{1, 1}, Vec2{1, 1}, Vec2{2, 2}},
		{Vec2{10, 1}, Vec2{-2, -5}, Vec2{8, -4}},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%d+%d==%d", tt.a, tt.b, tt.want)
		t.Run(testName, func(t *testing.T) {
			ans := tt.a.Add(tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
