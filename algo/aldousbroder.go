package algo

import (
	"github.com/fred1268/maze/maze"
)

func AldousBroder(m *maze.Maze) {
	cell := m.GetCell(rnd.Int()%m.Width, rnd.Int()%m.Length)
	m.Visit(cell)
	for {
		next := cell.Neighbors()[rnd.Int()%len(cell.Neighbors())]
		if !m.Visited(next) {
			cell.Link(next)
			m.Visit(next)
		}
		cell = next
		if m.IsFullyVisited() {
			break
		}
	}
}
