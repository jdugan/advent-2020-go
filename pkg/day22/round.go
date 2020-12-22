package day22

import (
  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Round struct {
  s1   int
  s2   int
}


// ========== RECEIVERS ===================================

func (r Round) Id () string {
  scores := pie.Ints{r.s1, r.s2}.Strings()

  return scores.Join(",")
}

func (r Round) HasWinner () bool {
  return r.Winner() > 0
}

func (r Round) HighScore () int {
  return pie.Ints{r.s1, r.s2}.Max()
}

func (r Round) Winner () int {
  winner := -1
  switch {
  case r.s1 > 0 && r.s2 == 0:
    winner = 1
  case r.s2 > 0 && r.s1 == 0:
    winner = 2
  }

  return winner
}
