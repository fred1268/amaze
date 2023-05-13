package maze

type point struct {
	x int
	y int
}

type solution struct {
	distances   [][]int
	path        map[point]struct{}
	MaxDistance int
	EntranceX   int
	EntranceY   int
	ExitX       int
	ExitY       int
}

func newSolution(maze *Maze) *solution {
	d := &solution{}
	d.distances = make([][]int, maze.Width)
	for x := 0; x < maze.Width; x++ {
		d.distances[x] = make([]int, maze.Length)
	}
	for x := 0; x < maze.Width; x++ {
		for y := 0; y < maze.Length; y++ {
			d.distances[x][y] = -1
		}
	}
	d.path = make(map[point]struct{})
	return d
}
