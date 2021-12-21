package stack

type Stack struct {
	items    []interface{}
	position int
}

func (s *Stack) Push(i interface{}) {
	if len(s.items) <= s.position {
		s.items = append(s.items, i)
	} else {
		s.items[s.position] = i
	}
	s.position++
}

func (s *Stack) Pop() interface{} {
	s.position--
	return s.items[s.position]
}

func (s *Stack) Remaining() []interface{} {
	r := make([]interface{}, 0, s.position)
	for !s.IsEmpty() {
		r = append(r, s.Pop())
	}
	return r
}

func (s *Stack) IsEmpty() bool {
	return s.position < 1
}

func (s *Stack) Len() int {
	return s.position
}

func MakeStack(size int) *Stack {
	return &Stack{
		items:    make([]interface{}, size),
		position: 0,
	}
}
