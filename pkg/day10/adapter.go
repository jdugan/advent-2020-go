package day10

import (
  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Adapter struct {
  joltage  int
}


// ========== RECEIVERS ===================================

func (a Adapter) CompatibleJoltages (adapters []Adapter) (matches pie.Ints) {
  for _, a1 := range adapters {
    step := a1.Difference(a)
    if step > 0 && step <= 3 {
      matches = append(matches, a1.joltage)
    }
  }
  return matches
}

func (a Adapter) Difference (a1 Adapter) int {
  return a.joltage - a1.joltage
}
