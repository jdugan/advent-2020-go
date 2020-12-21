package day21

import (
  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Product struct {
  ingredients   pie.Strings
}

type Products []Product


// ========== RECEIVERS ===================================

func (s Products) CountOccurrences (inames pie.Strings) int {
  count := 0
  for _, p := range s {
    for _, i := range p.ingredients {
      if inames.Contains(i) {
        count = count + 1
      }
    }
  }
  return count
}
