package vector

type Vec3 struct {
	X int
	Y int
	Z int
}

func (vec Vec3) Add(other Vec3) Vec3 {
	return Vec3{
		X: vec.X + other.X,
		Y: vec.Y + other.Y,
		Z: vec.Z + other.Z,
	}
}

func (vec Vec3) Sub(other Vec3) Vec3 {
	return Vec3{
		X: vec.X - other.X,
		Y: vec.Y - other.Y,
		Z: vec.Z - other.Z,
	}
}

func (vec Vec3) Mul(other Vec3) Vec3 {
	return Vec3{
		X: vec.X * other.X,
		Y: vec.Y * other.Y,
		Z: vec.Z * other.Z,
	}
}

type Transform3 struct {
	Values [3][3]int
}

func (t Transform3) Apply(v Vec3) Vec3 {
	return Vec3{
		X: v.X*t.Values[0][0] + v.Y*t.Values[1][0] + v.Z*t.Values[2][0],
		Y: v.X*t.Values[0][1] + v.Y*t.Values[1][1] + v.Z*t.Values[2][1],
		Z: v.X*t.Values[0][2] + v.Y*t.Values[1][2] + v.Z*t.Values[2][2],
	}
}

func (t Transform3) Determinant() int {
	return t.Values[0][0]*t.Values[1][1]*t.Values[2][2] +
		t.Values[0][1]*t.Values[1][2]*t.Values[2][0] +
		t.Values[0][2]*t.Values[1][0]*t.Values[2][1] -
		t.Values[0][2]*t.Values[1][1]*t.Values[2][0] -
		t.Values[0][1]*t.Values[1][0]*t.Values[2][2] -
		t.Values[0][0]*t.Values[1][2]*t.Values[2][1]
}

var ProperRotations = [24]Transform3{
	{[3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}},
	{[3][3]int{{1, 0, 0}, {0, 0, 1}, {0, -1, 0}}},
	{[3][3]int{{1, 0, 0}, {0, -1, 0}, {0, 0, -1}}},
	{[3][3]int{{1, 0, 0}, {0, 0, -1}, {0, 1, 0}}},
	{[3][3]int{{-1, 0, 0}, {0, -1, 0}, {0, 0, 1}}},
	{[3][3]int{{-1, 0, 0}, {0, 0, 1}, {0, 1, 0}}},
	{[3][3]int{{-1, 0, 0}, {0, 1, 0}, {0, 0, -1}}},
	{[3][3]int{{-1, 0, 0}, {0, 0, -1}, {0, -1, 0}}},
	{[3][3]int{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}}},
	{[3][3]int{{0, 0, 1}, {0, -1, 0}, {1, 0, 0}}},
	{[3][3]int{{0, -1, 0}, {0, 0, -1}, {1, 0, 0}}},
	{[3][3]int{{0, 0, -1}, {0, 1, 0}, {1, 0, 0}}},
	{[3][3]int{{0, -1, 0}, {0, 0, 1}, {-1, 0, 0}}},
	{[3][3]int{{0, 0, 1}, {0, 1, 0}, {-1, 0, 0}}},
	{[3][3]int{{0, 1, 0}, {0, 0, -1}, {-1, 0, 0}}},
	{[3][3]int{{0, 0, -1}, {0, -1, 0}, {-1, 0, 0}}},
	{[3][3]int{{0, 0, 1}, {1, 0, 0}, {0, 1, 0}}},
	{[3][3]int{{0, -1, 0}, {1, 0, 0}, {0, 0, 1}}},
	{[3][3]int{{0, 0, -1}, {1, 0, 0}, {0, -1, 0}}},
	{[3][3]int{{0, 1, 0}, {1, 0, 0}, {0, 0, -1}}},
	{[3][3]int{{0, 0, 1}, {-1, 0, 0}, {0, -1, 0}}},
	{[3][3]int{{0, 1, 0}, {-1, 0, 0}, {0, 0, 1}}},
	{[3][3]int{{0, 0, -1}, {-1, 0, 0}, {0, 1, 0}}},
	{[3][3]int{{0, -1, 0}, {-1, 0, 0}, {0, 0, -1}}},
}
