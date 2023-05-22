package algo

import (
	"github.com/fred1268/amaze/maze"
	"golang.org/x/exp/slices"
)

type choiceFunc func([]*maze.Cell) *maze.Cell

func GTRandom(list []*maze.Cell) *maze.Cell {
	return list[rnd.Int()%len(list)]
}

func GTLast(list []*maze.Cell) *maze.Cell {
	return list[len(list)-1]
}

func GTMix(list []*maze.Cell) *maze.Cell {
	if rnd.Int()%2 == 0 {
		return GTRandom(list)
	}
	return GTLast(list)
}

func GrowingTree(m *maze.Maze, fn choiceFunc) {
	var active []*maze.Cell
	cell := m.GetCell(rnd.Int()%m.Width, rnd.Int()%m.Length)
	active = append(active, cell)
	m.Visit(cell)
	for {
		cell := fn(active)
		neighbors := cell.UnvisitedNeighbors()
		if len(neighbors) != 0 {
			target := neighbors[rnd.Int()%len(neighbors)]
			cell.Link(target)
			active = append(active, target)
			m.Visit(target)
		} else {
			n := slices.Index(active, cell)
			newActive := active[0:n]
			newActive = append(newActive, active[n+1:]...)
			active = newActive
			if len(active) == 0 {
				break
			}
		}
	}
}
