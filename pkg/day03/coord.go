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

func (c Coord) Key () string {
  return fmt.Sprintf("%d,%d", c.x, c.y)
}

func (c Coord) Move (dx int, dy int, period int) Coord {
  x := (c.x + dx) % period
  y := c.y + dy

  return Coord{x: x, y: y}
}
