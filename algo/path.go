package algo

import (
	"github.com/fred1268/maze/maze"
)

type path struct {
	list []*maze.Cell
	all  map[*maze.Cell]struct{}
}

func newPath() *path {
	return &path{all: make(map[*maze.Cell]struct{})}
}

func (p *path) visit(cell *maze.Cell) {
	p.list = append(p.list, cell)
	p.all[cell] = struct{}{}
}

func (p *path) backtrackTo(cell *maze.Cell) {
	for i := len(p.list) - 1; i >= 0; i-- {
		current := p.list[i]
		if current == cell {
			p.list = p.list[0 : i+1]
			break
		}
		delete(p.all, current)
	}
}

func (p *path) backtrack() *maze.Cell {
	if len(p.list) == 0 {
		return nil
	}
	cell := p.list[len(p.list)-1]
	p.list = p.list[0 : len(p.list)-1]
	delete(p.all, cell)
	return cell
}

func (p *path) visited(cell *maze.Cell) bool {
	_, ok := p.all[cell]
	return ok
}

func (p *path) getPath() []*maze.Cell {
	return p.list
}
