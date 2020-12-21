package day21

import (
  "sort"

  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Ingredient struct {
  name      string
  allergen  string
}

type Ingredients []Ingredient


// ========== RECEIVERS ===================================

func (s Ingredients) SortedNames() pie.Strings {
  sort.Sort(s)

  names := pie.Strings{}
  for _, i := range s {
    names = append(names, i.name)
  }

  return names
}

// ---------- SORT INTERFACE ------------------------------

func (s Ingredients) Len() int {
    return len(s)
}

func (s Ingredients) Less(i, j int) bool {
    return s[i].allergen < s[j].allergen
}

func (s Ingredients) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
