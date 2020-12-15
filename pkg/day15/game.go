package day15

import (
  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Game struct {
  seeds   pie.Ints
}


// ========== RECEIVERS ===================================

// TODO: Good lord, refactor this.
//
func (g Game) Play (limit int) int {
  nmap := make(map[int]pie.Ints)

  nums := g.seeds
  turn := nums.Len()
  for i, n := range nums {
    v := i + 1
    nmap[n] = pie.Ints{v}
  }

  prev := nums.Last()
  curr := 0
  for turn < limit {
    turn = turn + 1

    ages, _ := nmap[prev]
    if ages.Len() == 1 {
      curr = 0
    } else {
      size   := ages.Len()
      final  := ages[size - 1]
      penult := ages[size - 2]
      curr    = final - penult
    }
    nmap[curr] = append(nmap[curr], turn)

    cages := nmap[curr]
    if cages.Len() > 2 {
      size   := cages.Len()
      final  := cages[size - 1]
      penult := cages[size - 2]
      nmap[curr] = pie.Ints{penult, final}
    }

    prev = curr
  }

  return curr
}


// ---------- UTILITIES -----------------------------------
