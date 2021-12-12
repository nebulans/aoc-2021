package day12

import (
	"aoc-2021/framework"
	"bufio"
	"strings"
)

type Connection struct {
	from string
	to   string
}

type Path struct {
	nodes []string
}

func (p *Path) add(node string) Path {
	n := Path{nodes: make([]string, len(p.nodes))}
	copy(n.nodes, p.nodes)
	n.nodes = append(n.nodes, node)
	return n
}

func (p *Path) Last() string {
	return p.nodes[len(p.nodes)-1]
}

func (p *Path) isComplete() bool {
	return p.Last() == "end"
}

func (p *Path) Format() string {
	return strings.Join(p.nodes, " -> ")
}

type PathStack struct {
	paths    []Path
	position int
}

func (s *PathStack) Push(path Path) {
	if len(s.paths) <= s.position {
		s.paths = append(s.paths, path)
	} else {
		s.paths[s.position] = path
	}
	s.position++
}

func (s *PathStack) Pop() Path {
	s.position--
	return s.paths[s.position]
}

func (s *PathStack) isEmpty() bool {
	return s.position < 1
}

func MakePathStack() *PathStack {
	return &PathStack{paths: []Path{}, position: 0}
}

type Graph struct {
	connections map[string][]string
	acceptFn    func(*Path, string) bool
}

func (g *Graph) connect(conn Connection) {
	g.connections[conn.from] = append(g.connections[conn.from], conn.to)
	g.connections[conn.to] = append(g.connections[conn.to], conn.from)
}

func (g *Graph) explore() []Path {
	complete := make([]Path, 0)
	stack := MakePathStack()
	stack.Push(Path{nodes: []string{"start"}})
	for !stack.isEmpty() {
		path := stack.Pop()
		possibles := g.connections[path.Last()]
		for _, poss := range possibles {
			if g.acceptFn(&path, poss) {
				n := path.add(poss)
				if n.isComplete() {
					complete = append(complete, n)
				} else {
					stack.Push(n)
				}
			}
		}
	}
	return complete
}

func MakeGraph() *Graph {
	return &Graph{
		connections: map[string][]string{},
	}
}

func NoRevisits(path *Path, node string) bool {
	if node == "start" {
		return false
	} // Always reject returning to start
	if node == "end" {
		return true
	} // Always accept arriving at end
	if strings.ToUpper(node) == node {
		return true
	}
	for _, n := range path.nodes {
		if n == node {
			return false
		}
	}
	return true
}

func SingleRevisit(path *Path, node string) bool {
	if node == "start" {
		return false
	} // Always reject returning to start
	if node == "end" {
		return true
	} // Always accept arriving at end
	if strings.ToUpper(node) == node {
		return true
	}
	revisited := false
	visitCounts := map[string]int{node: 1}
	for _, visited := range path.nodes {
		if strings.ToLower(visited) == visited {
			prev := visitCounts[visited]
			if prev > 0 {
				if revisited {
					return false
				}
				revisited = true
			}
			visitCounts[visited] = prev + 1
		}
	}
	return true
}

type Puzzle struct {
	framework.PuzzleBase
	connections []Connection
}

func (p *Puzzle) Init() {
	p.connections = make([]Connection, 0)
	p.Parts = map[string]func() int{
		"1": func() int { return p.countPaths(NoRevisits) },
		"2": func() int { return p.countPaths(SingleRevisit) },
	}
}

func (p *Puzzle) Parse(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		p.connections = append(p.connections, Connection{parts[0], parts[1]})
	}
}

func (p *Puzzle) countPaths(acceptFn func(*Path, string) bool) int {
	g := MakeGraph()
	g.acceptFn = acceptFn
	for _, conn := range p.connections {
		g.connect(conn)
	}
	paths := g.explore()
	return len(paths)
}
