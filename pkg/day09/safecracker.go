package day09

import (
  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Safecracker struct {
  preamble  pie.Ints
  numbers   pie.Ints
}


// ========== RECEIVERS ===================================

func (c Safecracker) Crack () int {
  calc    := Calculator{}
  numbers := append(c.preamble, c.numbers...)
  target  := c.Probe()

  factors  := calc.AnySequence(numbers, target).Sort()
  weakness := factors.First() + factors.Last()

  return weakness
}

func (c Safecracker) Probe () int {
  calc    := Calculator{}
  numbers := append(c.preamble, c.numbers...)
  result  := 0

  for i := 25; i < numbers.Len(); i++ {
    preamble := numbers[i-25:i]
    current  := numbers[i]

    if calc.AnyTwo(preamble, current).Len() == 0 {
      result = current
      break
    }
  }

  return result
}
