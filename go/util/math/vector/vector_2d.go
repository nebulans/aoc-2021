package vector

type Vec2 struct {
	X int
	Y int
}

func (vec *Vec2) Add(other Vec2) Vec2 {
	return Vec2{vec.X + other.X, vec.Y + other.Y}
}
