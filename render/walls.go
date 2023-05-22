package render

import (
	"github.com/fred1268/amaze/maze"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

func renderExternalWalls(maze *maze.Maze, scene *core.Node) {
	wallWidth := float32(maze.Width)*cellWidth + float32(maze.Width-1)*intWallWidth
	wallLength := float32(maze.Length)*cellLength + float32(maze.Length-1)*intWallLength + 2*extWallLength
	mat := material.NewStandard(&math32.Color{R: 0.3, G: 0.1, B: 0})
	// left wall
	wall := geometry.NewBox(extWallWidth, wallLength, extWallHeight)
	mesh := graphic.NewMesh(wall, mat)
	mesh.SetPosition(-(wallWidth+extWallWidth)/2, 0, extWallLevel)
	scene.Add(mesh)
	// right wall
	mesh = graphic.NewMesh(wall, mat)
	mesh.SetPosition((wallWidth+extWallWidth)/2, 0, extWallLevel)
	scene.Add(mesh)
	// top wall
	wall = geometry.NewBox(wallWidth, extWallLength, extWallHeight)
	mesh = graphic.NewMesh(wall, mat)
	mesh.SetPosition(0, (wallLength-extWallLength)/2, extWallLevel)
	scene.Add(mesh)
	// bottom wall
	mesh = graphic.NewMesh(wall, mat)
	mesh.SetPosition(0, -(wallLength-extWallLength)/2, extWallLevel)
	scene.Add(mesh)
}

func renderInternalWalls(maze *maze.Maze, scene *core.Node) {
	mazeWidth := float32(maze.Width)*cellWidth + float32(maze.Width-1)*intWallWidth
	mazeLength := float32(maze.Length)*cellLength + float32(maze.Length-1)*intWallLength
	wallWidth := cellWidth + intWallWidth
	wallLength := cellLength + intWallLength
	mat := material.NewStandard(&math32.Color{R: 0.15, G: 0.05, B: 0.05})
	for y := maze.Length - 1; y >= 0; y-- {
		for x := 0; x < maze.Width; x++ {
			cell := maze.GetCell(x, y)
			if x < maze.Width-1 && !cell.IsLinkedEast() {
				wall := geometry.NewBox(intWallWidth, wallLength, intWallHeight)
				mesh := graphic.NewMesh(wall, mat)
				xx := float32(x)*(cellWidth+intWallWidth) + cellWidth + intWallWidth/2 - mazeWidth/2
				yy := float32(y)*(cellLength+intWallLength) + cellLength/2 - mazeLength/2
				mesh.SetPosition(xx, yy, intWallLevel)
				scene.Add(mesh)
			}
			if y > 0 && !cell.IsLinkedSouth() {
				wall := geometry.NewBox(wallWidth, intWallLength, intWallHeight)
				mesh := graphic.NewMesh(wall, mat)
				xx := float32(x)*(cellWidth+intWallWidth) + cellWidth/2 - mazeWidth/2
				yy := float32(y-1)*(cellLength+intWallLength) + cellLength + intWallLength/2 - mazeLength/2
				mesh.SetPosition(xx, yy, intWallLevel)
				scene.Add(mesh)
			}
		}
	}
}
