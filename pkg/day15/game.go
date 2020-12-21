package day15

import (
  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Game struct {
  seeds   pie.Ints
}


// ========== RECEIVERS ===================================

func (g Game) Play (limit int) int {
  cache := Cache{}.Initialise(g.seeds)
  turn  := g.seeds.Len()
  prev  := g.seeds.Last()
  curr  := 0

  for turn < limit {
    turn        = turn + 1
    cache, curr = cache.Cycle(prev, turn)
    prev        = curr
  }

  return curr
}
