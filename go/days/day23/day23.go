package day23

import (
	"aoc-2021/framework"
	"aoc-2021/util/datastructure/priorityqueue"
	"bufio"
	"container/heap"
	"fmt"
)

var moveCosts = map[byte]int{
	'A': 1,
	'B': 10,
	'C': 100,
	'D': 1000,
}

var roomTargets = map[byte]int{
	'A': 7,
	'B': 8,
	'C': 9,
	'D': 10,
}

var roomConnections = map[int][2]int{
	7:  {1, 2},
	8:  {2, 3},
	9:  {3, 4},
	10: {4, 5},
}

var hallwayConnections = map[int][]int{
	1: {7},
	2: {7, 8},
	3: {8, 9},
	4: {9, 10},
	5: {10},
}

type Path struct {
	points []int
	cost   int
}

type PathStore struct {
	paths map[[2]int]Path
}

func (s *PathStore) Get(from, to int) Path {
	coord := [2]int{from, to}
	_, found := s.paths[coord]
	if !found {
		s.paths[coord] = s.generatePath(from, to)
	}
	return s.paths[coord]
}

func (s *PathStore) generatePath(start, end int) Path {
	steps := make([]int, 1, 7)
	steps[0] = start
	cost := 0

	next := func(pos, target int) (int, int) {
		if pos == 0 {
			return 1, 1
		}
		if pos == 6 {
			return 5, 1
		}
		if pos > 6 {
			if target > 6 && pos%4 == target%4 {
				// In final room, so descend towards position
				return pos + 4, 1
			}
			if pos > 10 {
				// Straightforward ascent
				return pos - 4, 1
			}
			routes := roomConnections[pos]
			if target < 7 {
				if target <= routes[0] {
					return routes[0], 2
				} else {
					return routes[1], 2
				}
			} else {
				if (pos+1)%4 < (target+1)%4 {
					return routes[1], 2
				} else {
					return routes[0], 2
				}
			}
		}
		if target <= 6 {
			// Handle ends of corridor
			if pos == 1 && target == 0 {
				return 0, 1
			}
			if pos == 5 && target == 6 {
				return 6, 1
			}
			if pos > target {
				return pos - 1, 2
			} else {
				return pos + 1, 2
			}
		}
		routes := hallwayConnections[pos]
		for _, r := range routes {
			if r%4 == target%4 {
				return r, 2
			}
		}
		if pos == 1 {
			return 2, 2
		}
		if pos == 2 {
			return 3, 2
		}
		if pos == 3 && target%4 == 3 {
			return 2, 2
		}
		if pos == 3 && target%4 == 2 {
			return 4, 2
		}
		if pos == 4 {
			return 3, 2
		}
		if pos == 5 {
			return 4, 2
		}
		panic("No route found")
	}

	i := 0
	for steps[len(steps)-1] != end {
		n, c := next(steps[len(steps)-1], end)
		steps = append(steps, n)
		cost += c
		if i > 20 {
			panic(fmt.Sprintf("Runaway step generation, path is %v", steps))
		}
		i++
	}
	return Path{steps, cost}
}

type Burrow struct {
	positions []byte
	energy    int
}

func (b *Burrow) Copy() Burrow {
	positions := make([]byte, len(b.positions))
	copy(positions, b.positions)
	return Burrow{positions: positions, energy: b.energy}
}

func (b *Burrow) Equals(other Burrow) bool {
	if len(b.positions) != len(other.positions) {
		return false
	}
	for i, t := range b.positions {
		if t != other.positions[i] {
			return false
		}
	}
	return true
}

func (b *Burrow) Format() string {
	o := make([]byte, len(b.positions))
	for i, c := range b.positions {
		if c == byte(0) {
			o[i] = ' '
		} else {
			o[i] = c
		}
	}
	out := fmt.Sprintf(".............\n"+
		".%c%c %c %c %c %c%c.\n"+
		"...%c.%c.%c.%c...\n",
		o[0], o[1], o[2], o[3], o[4], o[5], o[6],
		o[7], o[8], o[9], o[10],
	)
	for i := 11; i < len(b.positions); i += 4 {
		out += fmt.Sprintf("  .%c.%c.%c.%c.\n", o[i], o[i+1], o[i+2], o[i+3])
	}
	out += "  ........."
	return out
}

func (b *Burrow) candidateToTarget(pos int) (bool, int) {
	t := 0
	for t = roomTargets[b.positions[pos]]; t < len(b.positions); t += 4 {
		v := b.positions[t]
		if !(v == byte(0) || v == b.positions[pos]) {
			return false, 0
		}
	}
	t -= 4
	for b.positions[t] != byte(0) {
		t -= 4
	}
	return true, t
}

func (b *Burrow) Moves(s *PathStore) []Burrow {
	roomTargets := map[byte]int{
		'A': 7,
		'B': 8,
		'C': 9,
		'D': 10,
	}
	candidates := make([][2]int, 0, 56)
	for a, r := range roomTargets {
		for r < len(b.positions) {
			if b.positions[r] != byte(0) {
				// Moves from top of a room
				canMove := false
				for rr := r; rr < len(b.positions); rr += 4 {
					if b.positions[rr] != a {
						canMove = true
					}
				}
				if canMove {
					// Moves to corridor spaces
					for i := 0; i < 7; i++ {
						if b.positions[i] == byte(0) {
							candidates = append(candidates, [2]int{r, i})
						}
					}
					available, t := b.candidateToTarget(r)
					if available {
						candidates = append(candidates, [2]int{r, t})
					}
				}
				break
			}
			r += 4
		}
	}
	for i := 0; i < 7; i++ {
		if b.positions[i] != byte(0) {
			available, t := b.candidateToTarget(i)
			if available {
				candidates = append(candidates, [2]int{i, t})
			}
		}
	}

	moves := make([]Burrow, 0, len(candidates))
	for _, c := range candidates {
		p := s.Get(c[0], c[1])
		pathClear := true
		for i, p := range p.points {
			if i == 0 {
				continue
			}
			if b.positions[p] != byte(0) {
				pathClear = false
				break
			}
		}
		if pathClear {
			n := b.Copy()
			n.positions[c[0]] = byte(0)
			n.positions[c[1]] = b.positions[c[0]]
			n.energy += p.cost * moveCosts[b.positions[c[0]]]
			moves = append(moves, n)
		}

	}
	return moves
}

type Puzzle struct {
	framework.PuzzleBase
	startingState Burrow
}

func (p *Puzzle) Init() {
	p.Parts = map[string]func() int{
		"1": p.minEnergy2,
		"2": p.minEnergy4,
	}
}

func (p *Puzzle) Parse(scanner *bufio.Scanner) {
	lines := make([]string, 0, 4)
	for i := 0; i < 2; i++ {
		// Don't need first two lines
		scanner.Scan()
	}
	for scanner.Scan() {
		line := scanner.Text()
		if line[3] != '#' {
			lines = append(lines, line)
		}
	}
	states := make([]byte, 6+4*len(lines)+1)
	for i, line := range lines {
		states[7+i*4] = line[3]
		states[8+i*4] = line[5]
		states[9+i*4] = line[7]
		states[10+i*4] = line[9]
	}
	p.startingState = Burrow{positions: states}
}

func (p *Puzzle) solve(target Burrow) int {
	seenStates := map[string]int{}
	discardedStates := 0
	pq := make(priorityqueue.PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &priorityqueue.Item{Value: p.startingState.Copy(), Priority: 0})
	store := &PathStore{map[[2]int]Path{}}
	i := 0
	for pq.Len() > 0 {
		i++
		item := heap.Pop(&pq).(*priorityqueue.Item)
		burrow := item.Value.(Burrow)
		if burrow.Equals(target) {
			fmt.Printf("%d iterations\n%d discarded states\n", i, discardedStates)
			return burrow.energy
		}
		for _, b := range burrow.Moves(store) {
			e, found := seenStates[string(b.positions)]
			if found && e <= b.energy {
				discardedStates++
				continue
			}
			seenStates[string(b.positions)] = b.energy
			heap.Push(&pq, &priorityqueue.Item{Value: b, Priority: -1 * b.energy})
		}
	}
	fmt.Printf("Ran out of states after %d iterations\n", i)
	return 0
}

func (p *Puzzle) minEnergy2() int {
	//fmt.Println(p.startingState.Format())
	n := byte(0)
	target := Burrow{positions: []byte{n, n, n, n, n, n, n, 'A', 'B', 'C', 'D', 'A', 'B', 'C', 'D'}}
	return p.solve(target)

}

func (p *Puzzle) minEnergy4() int {
	n := byte(0)
	target := Burrow{positions: []byte{n, n, n, n, n, n, n, 'A', 'B', 'C', 'D', 'A', 'B', 'C', 'D', 'A', 'B', 'C', 'D', 'A', 'B', 'C', 'D'}}
	return p.solve(target)
}
