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
  chars := strings.Split(str, "")

  return pie.Strings(chars).Sort().Unique().Join("")
}

func (g Group) UnanimousAnswers () string {
  str, list := g.answers.Shift()

  arr := pie.Strings(strings.Split(str, "")).Sort()
  for _, item := range list {
    chars := pie.Strings(strings.Split(item, "")).Sort()
    arr    = arr.Intersect(chars)
  }

  return arr.Sort().Join("")
}
