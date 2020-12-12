package day12

import (
  "fmt"

  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Ferry struct {
  coord     Coord
  commands  []Command
}


// ========== RECEIVERS ===================================

func (f Ferry) DistanceTraveled () int {
  pie.Ints{f.x, f.y}.Abs().Sum()
}

func (f Ferry) Move (cmd Command) int {
  return f.DistanceTraveled()
}
