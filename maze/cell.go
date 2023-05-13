package maze

type Cell struct {
	maze  *Maze
	x     int
	y     int
	north *Cell
	south *Cell
	east  *Cell
	west  *Cell
	links map[*Cell]struct{}
}

func newCell(maze *Maze, x, y int) *Cell {
	return &Cell{maze: maze, x: x, y: y, links: make(map[*Cell]struct{})}
}

func (c *Cell) Link(other *Cell) {
	c.links[other] = struct{}{}
	other.links[c] = struct{}{}
}

func (c *Cell) LinkNorth() bool {
	if c.y < c.maze.Length-1 {
		c.Link(c.maze.cells[c.x][c.y+1])
		return true
	}
	return false
}

func (c *Cell) LinkSouth() bool {
	if c.y > 0 {
		c.Link(c.maze.cells[c.x][c.y-1])
		return true
	}
	return false
}

func (c *Cell) LinkEast() bool {
	if c.x < c.maze.Width-1 {
		c.Link(c.maze.cells[c.x+1][c.y])
		return true
	}
	return false
}

func (c *Cell) LinkWest() bool {
	if c.x > 0 {
		c.Link(c.maze.cells[c.x-1][c.y])
		return true
	}
	return false
}

func (c *Cell) IsLinkedNorth() bool {
	_, ok := c.links[c.north]
	return ok
}

func (c *Cell) IsLinkedSouth() bool {
	_, ok := c.links[c.south]
	return ok
}

func (c *Cell) IsLinkedEast() bool {
	_, ok := c.links[c.east]
	return ok
}

func (c *Cell) IsLinkedWest() bool {
	_, ok := c.links[c.west]
	return ok
}

func (c *Cell) Neighbors() []*Cell {
	var neighbors []*Cell
	if c.x > 0 {
		neighbors = append(neighbors, c.maze.cells[c.x-1][c.y])
	}
	if c.x < c.maze.Width-1 {
		neighbors = append(neighbors, c.maze.cells[c.x+1][c.y])
	}
	if c.y > 0 {
		neighbors = append(neighbors, c.maze.cells[c.x][c.y-1])
	}
	if c.y < c.maze.Length-1 {
		neighbors = append(neighbors, c.maze.cells[c.x][c.y+1])
	}
	return neighbors
}

func (c *Cell) unlinkedNeighbors() []*Cell {
	var unlinked []*Cell
	neighbors := c.Neighbors()
	for _, cell := range neighbors {
		if _, ok := c.links[cell]; !ok {
			unlinked = append(unlinked, cell)
		}
	}
	return unlinked
}

func (c *Cell) neighbors(visited bool) []*Cell {
	var neighbors []*Cell
	if c.x > 0 {
		cell := c.maze.cells[c.x-1][c.y]
		if c.maze.Visited(cell) == visited {
			neighbors = append(neighbors, cell)
		}
	}
	if c.x < c.maze.Width-1 {
		cell := c.maze.cells[c.x+1][c.y]
		if c.maze.Visited(cell) == visited {
			neighbors = append(neighbors, cell)
		}
	}
	if c.y > 0 {
		cell := c.maze.cells[c.x][c.y-1]
		if c.maze.Visited(cell) == visited {
			neighbors = append(neighbors, cell)
		}
	}
	if c.y < c.maze.Length-1 {
		cell := c.maze.cells[c.x][c.y+1]
		if c.maze.Visited(cell) == visited {
			neighbors = append(neighbors, cell)
		}
	}
	return neighbors
}

func (c *Cell) VisitedNeighbors() []*Cell {
	return c.neighbors(true)
}

func (c *Cell) UnvisitedNeighbors() []*Cell {
	return c.neighbors(false)
}
