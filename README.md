# :door: amaze :door:

A program to create, solve and draw mazes in your terminal or in 2D/3D

## Gallery

2D Terminal view
`10x10, Recursive Backtracker, Seed 1684354595222685`

<p align="center">
  <img style="float: right;" src="https://github.com/fred1268/assets/blob/development/amaze/10x10_recursivebacktracker_1684354595222685.png" alt="Amaze in Action"/>
</p>

2D Graphical view
`20x20, Growing Tree, Mix, Braid 80%, Seed 1684354465956585`

<p align="center">
  <img style="float: right;" src="https://github.com/fred1268/assets/blob/development/amaze/20x20_growingtree_mix_braid80_1684354465956585.png" alt="Amaze in Action"/>
</p>

3D Graphical view
`40x40, Hunt and Kill, Braid 50%, Seed 1684354734176462`

<p align="center">
  <img style="float: right;" src="https://github.com/fred1268/assets/blob/development/amaze/40x40_huntandkill_braid50_1684354734176462.png" alt="Amaze in Action"/>
</p>

---

## Setup

This program uses G3n, so it has [the same prerequisites](https://github.com/g3n/engine#dependencies).
It is pretty straightforward, since most of the computers nowadays, have an OpenGL driver and a C
compiler. I tested both on Linux (manjaro) and mac M1.

> Also, note that Amaze uses [**clap :clap:, the Command Line Argument Parser**](https://github.com/fred1268/go-clap).
> clap is a non intrusive, lightweight command line parsing library you may want to try out in your
> own projects. Feel free to give it a try!

---

## Installation

```
git clone github.com/fred1268/amaze
cd amaze
go install
```

---

## Running the program

You can run the program with several flags, although it has decent defaults. Here are the available flags:

```
amaze: creates, solves and displays mazes
        --width, -w <width>: set maze's width
        --length, -l <length>: set maze's length
        --algo, -a <name>: choose the algorithm used to create the maze:
           binarytree
           sidewinder
           aldousbroder
           wilson
           huntandkill
           recursivebacktracker
           growingtree
              use --choice, -c to chose sub algorithm:
              random, last, or mix
        --seed, -d <seed>: use seed to generate identical maze
        --braid, -b <percent>: braid to remove deadends
        --solve, -s: solve the maze
        --print, -p: print maze in the terminal
```

Description of the command line parameters

- width and length allow you to chose the size of your map (defaults to 20).

- Amaze uses various algorithms to generate mazes, so you can chose the one you prefer (defaults to binarytree).

> Please note that **Growing Tree** requires the choice of a sub-algorithm.
> You can chose between `random`, `last` and `mix` (defaults to `random`)

- seed allows you to replay a maze you liked in the past. The current maze's seed is displayed in the console
  when the program starts.

- braid allows you to remove none (0%), some (1-99%) or all (100%) of the deadends in the maze
  (defaults to 0, keep deadends). Do **not** include the % sign like in `--braid 50`.

- solve allows you to solve the maze, and will display the maze with a color gradient.
  The entrance (red dot) will be in lighter colors whereas the exit (green dot) will be in darker colors.
  By default, mazes are unresolved, just in case you want to solve them :wink:.

> Please note: in the terminal, only the exit path is highlighted in a gradient of color.

- print allows you to display the maze in the terminal only (no 3D).

## Feedback

Feel free to send feedback, star, etc.
