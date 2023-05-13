package algo

import (
	"github.com/fred1268/maze/maze"
)

func Wilson(m *maze.Maze) {
	cell := m.GetCell(rnd.Int()%m.Width, rnd.Int()%m.Length)
	m.Visit(cell)
	for {
		for {
			cell = m.GetCell(rnd.Int()%m.Width, rnd.Int()%m.Length)
			if !m.Visited(cell) {
				break
			}
		}
		path := newPath()
		path.visit(cell)
		for {
			next := cell.Neighbors()[rnd.Int()%len(cell.Neighbors())]
			if m.Visited(next) {
				path.visit(next)
				p := path.getPath()
				for i := 0; i < len(p)-1; i++ {
					p[i].Link(p[i+1])
					m.Visit(p[i])
				}
				m.Visit(p[len(p)-1])
				break
			}
			if path.visited(next) {
				path.backtrackTo(next)
			} else {
				path.visit(next)
			}
			cell = next
		}
		if m.IsFullyVisited() {
			break
		}
	}
}
