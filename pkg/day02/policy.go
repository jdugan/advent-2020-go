package day02

import (
  "strings"

  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Policy struct {
  min     int
  max     int
  letter  string
}


// ========== RECEIVERS ===================================

func (p Policy) OldMatch (password string) bool {
  strs    := pie.Strings(strings.Split(password, ""))
  matches := strs.Filter(func (s string) bool { return s == p.letter })
  count   := matches.Len()

  return count >= p.min && count <= p.max
}

func (p Policy) CurrentMatch (password string) bool {
  strs    := pie.Strings(strings.Split(password, ""))
  strs     = pie.Strings{strs[p.min - 1], strs[p.max - 1]}
  matches := strs.Filter(func (s string) bool { return s == p.letter })
  count   := matches.Len()

  return count == 1
}
