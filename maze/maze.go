package maze

import (
	"fmt"
	"math"

	"github.com/gookit/color"
	"golang.org/x/exp/slices"
)

type Maze struct {
	Width    int
	Length   int
	cells    [][]*Cell
	Solution *solution
	visited  map[*Cell]struct{}
}

func NewMaze(x, y int) *Maze {
	maze := &Maze{Width: x, Length: y}
	maze.visited = make(map[*Cell]struct{})
	maze.cells = make([][]*Cell, maze.Width)
	for x := 0; x < maze.Width; x++ {
		for y := 0; y < maze.Length; y++ {
			maze.cells[x] = append(maze.cells[x], newCell(maze, x, y))
		}
	}
	for x := 0; x < maze.Width; x++ {
		for y := maze.Length - 1; y >= 0; y-- {
			cell := maze.cells[x][y]
			if y < maze.Length-1 {
				cell.north = maze.cells[x][y+1]
			}
			if y > 0 {
				cell.south = maze.cells[x][y-1]
			}
			if x < maze.Width-1 {
				cell.east = maze.cells[x+1][y]
			}
			if x > 0 {
				cell.west = maze.cells[x-1][y]
			}
		}
	}
	return maze
}

func (m *Maze) Visit(cell *Cell) {
	m.visited[cell] = struct{}{}
}

func (m *Maze) Visited(cell *Cell) bool {
	_, ok := m.visited[cell]
	return ok
}

func (m *Maze) Unvisit(cell *Cell) {
	delete(m.visited, cell)
}

func (m *Maze) IsFullyVisited() bool {
	return len(m.visited) == m.Width*m.Length
}

func (m *Maze) GetCell(x, y int) *Cell {
	return m.cells[x][y]
}

func (m *Maze) GetDistance(x, y int) int {
	return m.Solution.distances[x][y]
}

func (m *Maze) getDeadends() []*Cell {
	var deadends []*Cell
	for y := m.Length - 1; y >= 0; y-- {
		for x := 0; x < m.Width; x++ {
			if len(m.cells[x][y].links) == 1 {
				deadends = append(deadends, m.cells[x][y])
			}
		}
	}
	return deadends
}

func (m *Maze) Braid(percent int) {
	deadends := m.getDeadends()
	for _, deadend := range deadends {
		if rnd.Int()%100 < percent {
			continue
		}
		var link *Cell
		unlinked := deadend.unlinkedNeighbors()
		if len(unlinked) == 0 {
			continue
		}
		for _, cell := range unlinked {
			if slices.Contains(deadends, cell) {
				link = cell
				break
			}
		}
		if link == nil {
			link = unlinked[rnd.Int()%len(unlinked)]
		}
		deadend.Link(link)
	}
}

func (m *Maze) Print(solve bool) string {
	out := "+"
	for i := 0; i < m.Width; i++ {
		out = fmt.Sprintf("%s%s", out, "---+")
	}
	out = fmt.Sprintf("%s\n", out)
	for y := m.Length - 1; y >= 0; y-- {
		top := "|"
		bottom := "+"
		for x := 0; x < m.Width; x++ {
			body := "   "
			if m.Solution != nil && solve {
				if _, ok := m.Solution.path[point{x: x, y: y}]; ok {
					if x == m.Solution.EntranceX && y == m.Solution.EntranceY {
						body = color.RGB(205, 0, 0).Sprintf(" o ")
					} else if x == m.Solution.ExitX && y == m.Solution.ExitY {
						body = color.RGB(0, 205, 0).Sprintf(" o ")
					} else {
						dist := float32(m.GetDistance(x, y)) / float32(m.Solution.MaxDistance)
						R := math.Max(208*(1-float64(dist))+48, 0)
						G := math.Max(208*(0.9-1.4*float64(dist))+48, 0)
						B := math.Max(208*(0.9-1.4*float64(dist))+48, 0)
						body = color.RGB(uint8(R), uint8(G), uint8(B)).Sprintf(" o ")
					}
				}
			}
			cell := m.cells[x][y]
			east := "|"
			if cell.IsLinkedEast() {
				east = " "
			}
			top = fmt.Sprintf("%s%s%s", top, body, east)
			south := "---"
			if cell.IsLinkedSouth() {
				south = "   "
			}
			bottom = fmt.Sprintf("%s%s+", bottom, south)
		}
		out = fmt.Sprintf("%s%s\n%s\n", out, top, bottom)
	}
	return out
}

func (m *Maze) String() string {
	return m.Print(false)
}
