package vector

import (
	"fmt"
	"testing"
)

func TestVec3_Add(t *testing.T) {
	var tests = []struct {
		a, b Vec3
		want Vec3
	}{
		{Vec3{0, 0, 0}, Vec3{1, 2, 3}, Vec3{1, 2, 3}},
		{Vec3{1, 2, 3}, Vec3{0, 0, 0}, Vec3{1, 2, 3}},
		{Vec3{1, 10, 100}, Vec3{200, 20, 2}, Vec3{201, 30, 102}},
		{Vec3{1, 2, 3}, Vec3{-3, -2, -1}, Vec3{-2, 0, 2}},
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

func TestTransform3_Determinant(t *testing.T) {
	var tests = []struct {
		m    Transform3
		want int
	}{
		{Transform3{[3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}}, 1},
		{Transform3{[3][3]int{{-1, 0, 0}, {0, -1, 0}, {0, 0, -1}}}, -1},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("det( %v )==%d", tt.m, tt.want)
		t.Run(testName, func(t *testing.T) {
			ans := tt.m.Determinant()
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestProperTransformDeterminant(t *testing.T) {
	for _, r := range ProperRotations {
		t.Run(fmt.Sprintf("Rotation %v", r), func(t *testing.T) {
			ans := r.Determinant()
			if ans != 1 {
				t.Errorf("Got determinant %d", ans)
			}
		})
	}
}

func TestProperTransformsUnique(t *testing.T) {
	for i, r := range ProperRotations {
		t.Run(fmt.Sprintf("Uniqueness of %v", r), func(t *testing.T) {
			for j, o := range ProperRotations {
				if i == j {
					continue
				}
				if r == o {
					t.Errorf("Repeated")
				}
			}
		})
	}
}

func TestTransform3_Apply(t *testing.T) {
	var tests = []struct {
		vec  Vec3
		r    Transform3
		want Vec3
	}{
		{
			Vec3{1, 0, 0},
			Transform3{[3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}},
			Vec3{1, 0, 0},
		},
		{
			Vec3{1, 2, 3},
			Transform3{[3][3]int{{0, 0, 1}, {0, 1, 0}, {-1, 0, 0}}},
			Vec3{-3, 2, 1},
		},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%v by %v == %d", tt.vec, tt.r, tt.want)
		t.Run(testName, func(t *testing.T) {
			ans := tt.r.Apply(tt.vec)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
