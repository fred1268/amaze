package algo

import (
	"github.com/fred1268/maze/maze"
)

func huntAndKill_NextCell(m *maze.Maze) *maze.Cell {
	var cell *maze.Cell
	for y := m.Length - 1; y >= 0; y-- {
		for x := 0; x < m.Width; x++ {
			cell = m.GetCell(x, y)
			if !m.Visited(cell) {
				visitedNeighbors := cell.VisitedNeighbors()
				if len(visitedNeighbors) != 0 {
					c := visitedNeighbors[rnd.Int()%len(visitedNeighbors)]
					cell.Link(c)
					m.Visit(cell)
					return cell
				}
			}
		}
	}
	return nil
}

func HuntAndKill(m *maze.Maze) {
	cell := m.GetCell(rnd.Int()%m.Width, rnd.Int()%m.Length)
	m.Visit(cell)
	for {
		unvisitedNeighbors := cell.UnvisitedNeighbors()
		if len(unvisitedNeighbors) != 0 {
			next := unvisitedNeighbors[rnd.Int()%len(unvisitedNeighbors)]
			cell.Link(next)
			m.Visit(next)
			cell = next
		} else {
			cell = huntAndKill_NextCell(m)
		}
		if m.IsFullyVisited() {
			break
		}
	}
}
