package integer

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func UnitStep(i int) int {
	if i > 0 {
		return 1
	}
	if i < 0 {
		return -1
	}
	return 0
}
