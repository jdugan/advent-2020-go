package day99

import (
  "fmt"

  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Coord struct {
  x     int
  y     int
  steps int
}


// ========== RECEIVERS ===================================

func (c1 Coord) Equals (c2 Coord) bool {
  return c1.x == c2.x && c1.y == c2.y
}

func (c Coord) Key () string {
  return fmt.Sprintf("%d,%d", c.x, c.y)
}

func (c Coord) Manhattan () int {
  return pie.Ints{c.x, c.y}.Abs().Sum()
}

func (c Coord) Move (dir string) Coord {
  c1 := Coord{x: c.x, y: c.y, steps: c.steps + 1}

  switch dir {
  case "D":
    c1.y = c1.y - 1
  case "L":
    c1.x = c1.x - 1
  case "R":
    c1.x = c1.x + 1
  case "U":
    c1.y = c1.y + 1
  }

  return c1
}
