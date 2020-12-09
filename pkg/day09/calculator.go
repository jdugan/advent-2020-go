package day09

import (
  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Calculator struct {}


// ========== RECEIVERS ===================================

func (c Calculator) AnySequence (numbers pie.Ints, target int) pie.Ints {
  sequence := pie.Ints{}
  curr     := 0

  for i := 0; i < numbers.Len(); i++ {
    farr   := pie.Ints{}
    list   := numbers[i:]
    for farr.Sum() < target {
      curr, list = list.Shift()
      farr       = append(farr, curr)
    }
    if farr.Sum() == target && farr.Len() > 1 {
      sequence = farr
      break
    }
  }

  return sequence
}

func (c Calculator) AnyTwo (numbers pie.Ints, target int) pie.Ints {
  factors := pie.Ints{}
  f0      := 0

  for factors.Len() == 0 && numbers.Len() > 0 {
    f0, numbers = numbers.Shift()
    f1 := target - f0
    if numbers.Contains(f1) {
      factors = pie.Ints{f0, f1}
    }
  }

  return factors
}
