package day11

import (
  "fmt"

  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Seat struct {
  x             int
  y             int
  occupied      bool

  adjacentIds   pie.Strings
  visibleIds    pie.Strings
}


// ========== RECEIVERS ===================================

func (s Seat) Copy () Seat {
  return Seat{x:            s.x,
              y:            s.y,
              occupied:     s.occupied,
              adjacentIds:  s.adjacentIds,
              visibleIds:   s.visibleIds}
}

func (s Seat) Id () string {
  return fmt.Sprintf("%d,%d", s.x, s.y)
}
