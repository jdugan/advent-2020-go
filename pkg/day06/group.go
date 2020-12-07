package day06

import (
  "strings"

  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Group struct {
  answers  pie.Strings
}


// ========== RECEIVERS ===================================

func (g Group) DistinctAnswers () string {
  str   := g.answers.Join("")
  chars := pie.Strings(strings.Split(str, "")).Sort()

  return chars.Unique().Join("")
}

func (g Group) UnanimousAnswers () string {
  head, tail := g.answers.Shift()

  chars := pie.Strings(strings.Split(head, "")).Sort()
  for _, str := range tail {
    tmp  := pie.Strings(strings.Split(str, "")).Sort()
    chars = chars.Intersect(tmp)
    if len(chars) == 0 {
      break
    }
  }

  return chars.Sort().Join("")
}
