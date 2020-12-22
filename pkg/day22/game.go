package day22

import (
  "fmt"

  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Game struct {
  d1  pie.Ints
  d2  pie.Ints
}


// ========== RECEIVERS ===================================

// This function plays a simple game without
// any special rules or recursion.
//
func (g Game) Play () Result {
  var round Round

  d1    := g.d1[0:]
  d2    := g.d2[0:]
  halt  := false

  for !halt {
    var c1, c2 int

    c1, d1 = d1.Shift()
    c2, d2 = d2.Shift()

    if c1 > c2 {
      d1 = append(d1, c1, c2)
    } else {
      d2 = append(d2, c2, c1)
    }

    s1   := g.ScoreDeck(d1)
    s2   := g.ScoreDeck(d2)
    round = Round{s1: s1, s2: s2}

    if round.HasWinner() {
      halt = true
    }
  }

  winner := round.Winner()
  score  := round.HighScore()
  result := Result{winner: winner, score: score}

  return result
}

// This function plays a more complex version
// of the game, considering things like previous
// play states, recursive decks, etc.
//
// To be runnable, we have to record the result
// not just for the original game conditions but
// also for every round of that game, which could
// be the original state of other games and will,
// we know now, yield the same result.
//
func (g Game) PlayRecursively () Result {
  var result Result

  d1     := g.d1[0:]
  d2     := g.d2[0:]
  halt   := false

  rindex := RoundIndexer{}
  s1     := g.ScoreDeck(d1)
  s2     := g.ScoreDeck(d2)
  round  := Round{s1: s1, s2: s2}

  for !halt {
    var c1, c2 int
    var ok bool

    rindex, ok = rindex.Register(round)
    if !ok {
      result = Result{}.Forfeit()
      halt   = true
      break
    }

    c1, d1 = d1.Shift()
    c2, d2 = d2.Shift()

    if g.CanRecurse(c1, d1) && g.CanRecurse(c2, d2) {
      rd1 := d1.Top(c1)
      rd2 := d2.Top(c2)

      rg  := Game{d1: rd1, d2: rd2}
      rr  := rg.PlayRecursively()

      switch rr.winner {
      case 1:
        d1 = append(d1, c1, c2)
      case 2:
        d2 = append(d2, c2, c1)
      default:
        fmt.Println("ERROR: Game ended with invalid winner", rr.winner)
        halt = true
        break
      }
    } else {
      if c1 > c2 {
        d1 = append(d1, c1, c2)
      } else {
        d2 = append(d2, c2, c1)
      }
    }

    s1    = g.ScoreDeck(d1)
    s2    = g.ScoreDeck(d2)
    round = Round{s1: s1, s2: s2}

    if round.HasWinner() {
      winner := round.Winner()
      score  := round.HighScore()
      result  = Result{winner: winner, score: score}
      halt    = true
    }
  }

  return result
}


// ---------- UTILITIES -----------------------------------

func (g Game) CanRecurse (card int, deck pie.Ints) bool {
  return deck.Len() >= card
}

func (g Game) ScoreDeck (deck pie.Ints) int {
  score := 0

  cards := deck.Reverse()
  for i, c := range cards {
    factor := i + 1
    score   = score + (c * factor)
  }

  return score
}
