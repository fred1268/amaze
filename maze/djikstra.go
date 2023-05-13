package maze

func dijkstra(maze *Maze, startX, startY int) *solution {
	solution := newSolution(maze)
	solution.EntranceX = startX
	solution.EntranceY = startY
	solution.MaxDistance = -1
	solution.distances[startX][startY] = 0
	links := map[*Cell]struct{}{maze.GetCell(startX, startY): {}}
	for {
		linksLinks := make(map[*Cell]struct{})
		for cell := range links {
			for link := range cell.links {
				if solution.distances[link.x][link.y] == -1 {
					solution.distances[link.x][link.y] = solution.distances[cell.x][cell.y] + 1
					linksLinks[link] = struct{}{}
					if solution.distances[link.x][link.y] > solution.MaxDistance {
						solution.ExitX = link.x
						solution.ExitY = link.y
						solution.MaxDistance = solution.distances[link.x][link.y]
					}
				}
			}
		}
		if len(linksLinks) == 0 {
			break
		}
		links = linksLinks
	}
	return solution
}

func computeExitPath(maze *Maze) {
	var pt point
	pt.x = maze.Solution.ExitX
	pt.y = maze.Solution.ExitY
	maze.Solution.path[pt] = struct{}{}
	for {
		cell := maze.GetCell(pt.x, pt.y)
		for c := range cell.links {
			if maze.Solution.distances[c.x][c.y] == maze.Solution.distances[pt.x][pt.y]-1 {
				pt.x = c.x
				pt.y = c.y
				maze.Solution.path[pt] = struct{}{}
				break
			}
		}
		if pt.x == maze.Solution.EntranceX && pt.y == maze.Solution.EntranceY {
			break
		}
	}
}

func Solve(maze *Maze) {
	solution := dijkstra(maze, 0, 0)
	maze.Solution = dijkstra(maze, solution.ExitX, solution.ExitY)
	computeExitPath(maze)
}
