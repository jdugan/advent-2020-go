package day12

import (
  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Coord struct {
  x         int
  y         int
  heading   string
}


// ========== RECEIVERS ===================================

func (c Coord) DistanceFromOrigin () int {
  factors := pie.Ints{c.x, c.y}.Abs()
  
  return factors.Sum()
}


// ---------- MOVE HELPERS --------------------------------

func (c Coord) MoveCardinally (cmd Command) Coord {
  x := c.x
  y := c.y

  switch cmd.dir {
  case "N":
    y = y + cmd.dist
  case "E":
    x = x + cmd.dist
  case "S":
    y = y - cmd.dist
  case "W":
    x = x - cmd.dist
  }

  return Coord{x: x, y: y, heading: c.heading}
}

func (c Coord) MoveForward (cmd Command) Coord {
  cmd1 := Command{dir: c.heading, dist: cmd.dist}

  return c.MoveCardinally(cmd1)
}

func (c Coord) MoveRelatively (cmd Command, waypoint Coord) Coord {
  dx := waypoint.x * cmd.dist
  dy := waypoint.y * cmd.dist

  return Coord{x: c.x + dx, y: c.y + dy, heading: c.heading}
}


// ---------- ROTATE HELPERS ------------------------------

func (c Coord) RotateLeft (cmd Command) Coord {
  x     := c.x
  y     := c.y
  turns := (cmd.dist / 90) % 4

  switch turns {
  case 1:
    x = c.y * -1
    y = c.x
  case 2:
    x = c.x * -1
    y = c.y * -1
  case 3:
    x = c.y
    y = c.x * -1
  }

  return Coord{x: x, y: y, heading: c.heading}
}

func (c Coord) RotateRight (cmd Command) Coord {
  x     := c.x
  y     := c.y
  turns := (cmd.dist / 90) % 4

  switch turns {
  case 1:
    x = c.y
    y = c.x * -1
  case 2:
    x = c.x * -1
    y = c.y * -1
  case 3:
    x = c.y * -1
    y = c.x
  }

  return Coord{x: x, y: y, heading: c.heading}
}


// ---------- TURN HELPERS --------------------------------

func (c Coord) Turn (headings pie.Strings, degrees int) Coord {
  turns := (degrees / 90) % 4
  index := headings.FindFirstUsing(func (h string) bool {
    return c.heading == h
  })
  index  = (index + turns) % 4

  return Coord{x: c.x, y: c.y, heading: headings[index]}
}

func (c Coord) TurnLeft (cmd Command) Coord {
  headings := pie.Strings{"N", "W", "S", "E"}

  return c.Turn(headings, cmd.dist)
}

func (c Coord) TurnRight (cmd Command) Coord {
  headings := pie.Strings{"N", "E", "S", "W"}

  return c.Turn(headings, cmd.dist)
}
