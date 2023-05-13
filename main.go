package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fred1268/go-clap/clap"
	"github.com/fred1268/maze/algo"
	"github.com/fred1268/maze/maze"
	"github.com/fred1268/maze/render"
)

type Config struct {
	Width     int    `clap:"--width,-w"`
	Length    int    `clap:"--length,-l"`
	Algorithm string `clap:"--algo,-a"`
	GTChoice  string `clap:"--choice,-c"`
	Seed      int    `clap:"--seed,-d"`
	Braid     int    `clap:"--braid,-b"`
	Print     bool   `clap:"--print,-p"`
	Solve     bool   `clap:"--solve,-s"`
	Help      bool   `clap:"--help,-h"`
}

func main() {
	var cfg *Config = &Config{
		Width:     20,
		Length:    20,
		Algorithm: "binarytree",
		Seed:      int(time.Now().UnixMicro()),
		Solve:     false,
		Help:      false,
		Braid:     0,
	}
	clap.Parse(os.Args, cfg)
	if cfg.Help {
		fmt.Printf("%s: creates, solves and displays mazes\n", os.Args[0])
		fmt.Printf("\t--width, -w <width>: set maze's width\n")
		fmt.Printf("\t--length, -l <length>: set maze's length\n")
		fmt.Printf("\t--algo, -a <name>: choose the algorithm used to create the maze:\n")
		fmt.Printf("\t   binarytree\n")
		fmt.Printf("\t   sidewinder\n")
		fmt.Printf("\t   aldousbroder\n")
		fmt.Printf("\t   wilson\n")
		fmt.Printf("\t   huntandkill\n")
		fmt.Printf("\t   recursivebacktracker\n")
		fmt.Printf("\t   growingtree\n")
		fmt.Printf("\t      use --choice, -c to chose sub algorithm:\n")
		fmt.Printf("\t      random, last, or mix\n")
		fmt.Printf("\t--seed, -d <seed>: use seed to generate identical maze\n")
		fmt.Printf("\t--braid, -b <percent>: braid to remove deadends\n")
		fmt.Printf("\t--solve, -s: solve the maze\n")
		fmt.Printf("\t--print, -p: print maze in the terminal\n")
		os.Exit(0)
	}
	fmt.Printf("Current seed: %d\n", cfg.Seed)
	algo.SetSeed(int64(cfg.Seed))
	maze.SetSeed(int64(cfg.Seed))
	m := maze.NewMaze(cfg.Width, cfg.Length)
	switch cfg.Algorithm {
	case "binarytree":
		algo.BinaryTree(m)
	case "sidewinder":
		algo.SideWinder(m)
	case "aldousbroder":
		algo.AldousBroder(m)
	case "wilson":
		algo.Wilson(m)
	case "huntandkill":
		algo.HuntAndKill(m)
	case "recursivebacktracker":
		algo.RecursiveBacktracker(m)
	case "growingtree":
		fn := algo.GTRandom
		switch cfg.GTChoice {
		case "random":
			fn = algo.GTRandom
		case "last":
			fn = algo.GTLast
		case "mix":
			fn = algo.GTMix
		}
		algo.GrowingTree(m, fn)
	}
	if cfg.Braid != 0 {
		m.Braid(cfg.Braid)
	}
	maze.Solve(m)
	if cfg.Print {
		render.RenderTerminal(m, cfg.Solve)
		os.Exit(0)
	}
	render.Render3D(m, cfg.Solve)
}
