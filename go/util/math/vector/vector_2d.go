package vector

type Vec2 struct {
	X int
	Y int
}

func (vec Vec2) Add(other Vec2) Vec2 {
	return Vec2{vec.X + other.X, vec.Y + other.Y}
}

func (vec Vec2) Mul(val int) Vec2 {
	return Vec2{X: vec.X * val, Y: vec.Y * val}
}
