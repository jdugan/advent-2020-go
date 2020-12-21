package day21

import (
  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type AllergenMap map[string]pie.Strings


// ========== RECEIVERS ===================================

func (am AllergenMap) RegisterEntriesAndReduce (anames pie.Strings, inames pie.Strings) AllergenMap {
  am = am.RegisterEntries(anames, inames)
  am = am.Reduce()

  return am
}


// ---------- BUILD HELPERS -------------------------------

func (am AllergenMap) Reduce () AllergenMap {
  ikeys := pie.Strings{}
  akeys := pie.Strings{}

  for k, v := range am {
    switch v.Len() {
    case 0:
      // nooop
    case 1:
      ikeys = append(ikeys, v.First())
    default:
      akeys = append(akeys, k)
    }
  }

  for _, ikey := range ikeys {
    for _, akey := range akeys {
      list, _ := am[akey]
      list     = list.FilterNot(func (s string) bool { return s == ikey })
      am[akey] = list
    }
  }

  return am
}

func (am AllergenMap) RegisterEntries (anames pie.Strings, inames pie.Strings) AllergenMap {
  for _, k := range anames {
    enames, exists := am[k]
    if exists {
      am[k] = enames.Intersect(inames)
    } else {
      am[k] = inames
    }
  }

  return am
}
