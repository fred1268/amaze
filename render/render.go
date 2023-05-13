package render

import (
	"fmt"
	"time"

	"github.com/fred1268/maze/maze"
	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/window"
)

func Render3D(maze *maze.Maze, solve bool) {
	a := app.App()
	scene := core.NewNode()
	gui.Manager().Set(scene)

	cam := camera.New(1)
	cam.SetPosition(0, 0, 250)
	scene.Add(cam)

	camera.NewOrbitControl(cam)
	onResize := func(evname string, ev interface{}) {
		width, height := a.GetSize()
		a.Gls().Viewport(0, 0, int32(width), int32(height))
		cam.SetAspect(float32(width) / float32(height))
	}
	a.Subscribe(window.OnWindowSize, onResize)
	onResize("", nil)

	renderExternalWalls(maze, scene)
	renderInternalWalls(maze, scene)
	renderCells(maze, scene, solve)

	scene.Add(light.NewAmbient(&math32.Color{R: 1.0, G: 1.0, B: 1.0}, 2.0))
	pointLight := light.NewPoint(&math32.Color{R: 1, G: 1, B: 1}, 2.0)
	pointLight.SetPosition(1, 1, 3)
	scene.Add(pointLight)

	a.Gls().ClearColor(0.5, 0.5, 1, 1.0)

	a.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		a.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
		renderer.Render(scene, cam)
	})
}

func RenderTerminal(maze *maze.Maze, solve bool) {
	fmt.Println(maze.Print(solve))
}
