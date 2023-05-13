package render

import (
	"math"

	"github.com/fred1268/maze/maze"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

func renderCells(maze *maze.Maze, scene *core.Node, solve bool) {
	mazeWidth := float32(maze.Width)*cellWidth + float32(maze.Width-1)*intWallWidth
	mazeLength := float32(maze.Length)*cellLength + float32(maze.Length-1)*intWallLength
	for y := maze.Length - 1; y >= 0; y-- {
		for x := 0; x < maze.Width; x++ {
			var dist float32 = 0
			if maze.GetDistance(x, y) != -1 {
				dist = float32(maze.GetDistance(x, y)) / float32(maze.Solution.MaxDistance)
			}
			mat := material.NewStandard(&math32.Color{R: 0.2, G: 0.06, B: 0})
			if solve {
				dist = float32(4*math.Pow(float64(dist-0.5), 3) + 0.5)
				R := 1 - dist
				G := 0.9 - 1.4*dist
				B := 0.9 - 1.4*dist
				mat = material.NewStandard(&math32.Color{R: R, G: G, B: B})
			}
			floor := geometry.NewBox(cellWidth+intWallWidth, cellLength+intWallLength, cellHeight)
			mesh := graphic.NewMesh(floor, mat)
			xx := float32(x)*(cellWidth+intWallWidth) + cellWidth/2 - mazeWidth/2
			yy := float32(y)*(cellLength+intWallLength) + cellLength/2 - mazeLength/2
			mesh.SetPosition(xx, yy, 0)
			scene.Add(mesh)
			if (x == maze.Solution.EntranceX && y == maze.Solution.EntranceY) ||
				(x == maze.Solution.ExitX && y == maze.Solution.ExitY) {
				var mat *material.Standard
				if x == maze.Solution.EntranceX && y == maze.Solution.EntranceY {
					mat = material.NewStandard(&math32.Color{R: 0.8, G: 0, B: 0})
				} else {
					mat = material.NewStandard(&math32.Color{R: 0, G: 0.8, B: 0})
				}
				floor := geometry.NewBox(cellWidth/2, cellLength/2, cellHeight)
				mesh := graphic.NewMesh(floor, mat)
				xx := float32(x)*(cellWidth+intWallWidth) + cellWidth/2 - mazeWidth/2
				yy := float32(y)*(cellLength+intWallLength) + cellLength/2 - mazeLength/2
				mesh.SetPosition(xx, yy, cellHeight)
				scene.Add(mesh)

			}
		}
	}
}
