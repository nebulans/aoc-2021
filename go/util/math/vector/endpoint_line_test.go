package vector

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDirection(t *testing.T) {
	var tests = []struct {
		start, end Vec2
		want       Vec2
	}{
		{Vec2{0, 0}, Vec2{0, 1}, Vec2{0, 1}},
		{Vec2{0, 0}, Vec2{1, 0}, Vec2{1, 0}},
		{Vec2{0, 0}, Vec2{1, 1}, Vec2{1, 1}},
		{Vec2{1, 1}, Vec2{0, 0}, Vec2{-1, -1}},
		{Vec2{10, 10}, Vec2{0, 0}, Vec2{-1, -1}},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%d->%d:%d", tt.start, tt.end, tt.want)
		t.Run(testName, func(t *testing.T) {
			line := NewEndpointLine(tt.start, tt.end)
			if line.Direction != tt.want {
				t.Errorf("got %d, want %d", line.Direction, tt.want)
			}
		})
	}
}

func TestPoints(t *testing.T) {
	var tests = []struct {
		start, end Vec2
		want       []Vec2
	}{
		{Vec2{0, 0}, Vec2{0, 1}, []Vec2{
			{0, 0},
			{0, 1},
		}},
		{Vec2{0, 0}, Vec2{1, 0}, []Vec2{
			{0, 0},
			{1, 0},
		}},
		{Vec2{0, 0}, Vec2{1, 1}, []Vec2{
			{0, 0},
			{1, 1},
		}},
		{Vec2{0, 0}, Vec2{0, 3}, []Vec2{
			{0, 0},
			{0, 1},
			{0, 2},
			{0, 3},
		}},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%d->%d", tt.start, tt.end)
		t.Run(testName, func(t *testing.T) {
			line := NewEndpointLine(tt.start, tt.end)
			if !reflect.DeepEqual(line.Points(), tt.want) {
				t.Errorf("got %v, want %v", line.Points(), tt.want)
			}
		})
	}
}
