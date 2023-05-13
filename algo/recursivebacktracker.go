package algo

import (
	"github.com/fred1268/maze/maze"
)

func recursiveBacktracker_NextCell(m *maze.Maze, path *path) *maze.Cell {
	var prev *maze.Cell
	var unvisitedNeighbors []*maze.Cell
	for {
		prev = path.backtrack()
		if prev == nil {
			break
		}
		unvisitedNeighbors = prev.UnvisitedNeighbors()
		if len(unvisitedNeighbors) != 0 {
			break
		}
	}
	if len(unvisitedNeighbors) != 0 {
		next := unvisitedNeighbors[rnd.Int()%len(unvisitedNeighbors)]
		prev.Link(next)
		return next
	}
	return nil
}

func RecursiveBacktracker(m *maze.Maze) {
	path := newPath()
	cell := m.GetCell(rnd.Int()%m.Width, rnd.Int()%m.Length)
	m.Visit(cell)
	path.visit(cell)
	for {
		var next *maze.Cell
		unvisitedNeighbors := cell.UnvisitedNeighbors()
		if len(unvisitedNeighbors) != 0 {
			next = unvisitedNeighbors[rnd.Int()%len(unvisitedNeighbors)]
			cell.Link(next)
		} else {
			next = recursiveBacktracker_NextCell(m, path)
		}
		m.Visit(next)
		path.visit(next)
		cell = next
		if m.IsFullyVisited() {
			break
		}
	}
}
