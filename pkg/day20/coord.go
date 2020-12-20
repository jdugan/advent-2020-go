package day20

import (
  "fmt"
)


// ========== DEFINITION ==================================

type Coord struct {
  x       int
  y       int
  display string
}


// ========== RECEIVERS ===================================

func (c Coord) GenerateId (x int, y int) string {
  return fmt.Sprintf("%d,%d", x, y)
}

func (c Coord) Id () string {
  return c.GenerateId(c.x, c.y)
}


// ---------- UTILITIES -----------------------------------
