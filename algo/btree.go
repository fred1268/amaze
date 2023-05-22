package algo

import (
	"github.com/fred1268/amaze/maze"
)

func BinaryTree(m *maze.Maze) {
	for y := m.Length - 1; y >= 0; y-- {
		for x := 0; x < m.Width; x++ {
			cell := m.GetCell(x, y)
			if rnd.Int()%2 == 0 {
				if !cell.LinkEast() {
					cell.LinkNorth()
				}
			} else {
				if !cell.LinkNorth() {
					cell.LinkEast()
				}
			}
		}
	}
}
