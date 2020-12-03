package day03

import (
  "fmt"
)


// ========== DEFINITION ==================================

type Coord struct {
  x int
  y int
}


// ========== RECEIVERS ===================================

func (c Coord) Base (period int) Coord {
  return Coord{x: c.x % period, y: c.y}
}

func (c1 Coord) Equals (c2 Coord) bool {
  return c1.x == c2.x && c1.y == c2.y
}

func (c Coord) Key () string {
  return fmt.Sprintf("%d,%d", c.x, c.y)
}

func (c Coord) Move (dx int, dy int) Coord {
  return Coord{x: c.x + dx, y: c.y + dy}
}
