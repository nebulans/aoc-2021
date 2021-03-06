package integer

import "math/big"

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxSlice(s []int) int {
	m := s[0]
	for _, v := range s {
		if v > m {
			m = v
		}
	}
	return m
}

func MinSlice(s []int) int {
	m := s[0]
	for _, v := range s {
		if v < m {
			m = v
		}
	}
	return m
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

func Sum(a []int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

func Product(a []int) int {
	p := 1
	for _, v := range a {
		p *= v
	}
	return p
}

func SumUint64(a []uint64) uint64 {
	s := uint64(0)
	for _, v := range a {
		s += v
	}
	return s
}

func SumBigInt(a []*big.Int) *big.Int {
	s := big.NewInt(0)
	for _, v := range a {
		s.Add(s, v)
	}
	return s
}

func Pow(value int, exponent int) int {
	if exponent == 0 {
		return 1
	}
	v := value
	for i := 0; i < exponent-1; i++ {
		v *= value
	}
	return v
}
