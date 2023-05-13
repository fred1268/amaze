package algo

import (
	"math/rand"

	"github.com/fred1268/maze/maze"
)

func SideWinder(m *maze.Maze) {
	for y := m.Length - 1; y >= 0; y-- {
		var visited []*maze.Cell = nil
		for x := 0; x < m.Width; x++ {
			cell := m.GetCell(x, y)
			if rand.Int()%2 == 0 {
				visited = append(visited, cell)
				if !cell.LinkEast() {
					cell.LinkNorth()
					visited = nil
				}
			} else {
				visited = append(visited, cell)
				n := 0
				if len(visited) > 1 {
					n = rnd.Int() % (len(visited) - 1)
				}
				if !visited[n].LinkNorth() {
					cell.LinkEast()
					visited = nil
				}
				visited = nil
			}
		}
	}
}
