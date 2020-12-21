package day15

import (
  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Cache map[int]pie.Ints


// ========== RECEIVERS ===================================

func (c Cache) Cycle (prev int, turn int) (Cache, int) {
  ages, _ := c[prev]
  curr    := 0
  if ages.Len() > 1 {
    curr = ages[1] - ages[0]
  }
  c[curr] = append(c[curr], turn).Bottom(2).Reverse()

  return c, curr
}


// ---------- BUILD HELPERS -------------------------------

func (c Cache) Initialise (seeds pie.Ints) Cache {
  for i, seed := range seeds {
    c[seed] = pie.Ints{i + 1}
  }
  return c
}
