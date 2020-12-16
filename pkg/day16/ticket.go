package day16

import (
  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Ticket struct {
  errorcode   int
  values      pie.Ints
}


// ========== RECEIVERS ===================================

func (t Ticket) DefaultIndices () pie.Ints {
  indices := pie.Ints{}
  for i, _ := range t.values {
    indices = append(indices, i)
  }
  return indices
}
