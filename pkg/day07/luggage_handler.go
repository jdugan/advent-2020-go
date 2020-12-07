package day07

import (
  "github.com/elliotchance/pie/pie"
)


// ========== SUB-DEFINITIONS =============================

type BagGroup struct {
  quantity  int
  name      string
}

type BagGroups []BagGroup


// ========== DEFINITION ==================================

type LuggageHandler struct {
  ruleset   map[string]BagGroups
}


// ========== RECEIVERS ===================================

// Counts the number of bags enclosed by the supplied
// bag group. Multiplies contained counts by the quantity
// and adds the quantity itself to account for the
// enclosing bags themselves.
//
func (lh LuggageHandler) CountEnclosedBags (bagGroup BagGroup) int {
  factor := bagGroup.quantity
  bgs    := lh.ruleset[bagGroup.name]

  sum := 0
  for _, bg := range bgs {
    sum = sum + lh.CountEnclosedBags(bg)
  }

  return (factor * sum) + factor
}

// Counts the number of distinct bag groups that could
// eventually enclose at least one bag of the specified
// type. Checks the bag groups in the rules set and
// walks backward, careful to eliminate duplicates as it
// goes to avoid infinite recursion.
//
func (lh LuggageHandler) CountPossibleContainers (matchedNames pie.Strings, requestedNames pie.Strings) pie.Strings {
  // find containers for requested bags
  newBagNames := pie.Strings{}
  for name, bagGroups := range lh.ruleset {
    for _, bg := range bagGroups {
      if requestedNames.Contains(bg.name) {
        newBagNames = append(newBagNames, name)
        break
      }
    }
  }
  newBagNames = newBagNames.Sort().Unique()

  // limit list to containers we haven't checked before
  uniqueBagNames := pie.Strings{}
  for _, bagName := range newBagNames {
    if !matchedNames.Contains(bagName) {
      uniqueBagNames = append(uniqueBagNames, bagName)
    }
  }

  // if more bags to check, do it; else, we're finished
  if len(uniqueBagNames) > 0 {
    matchedNames = matchedNames.Append(uniqueBagNames...).Sort()
    return lh.CountPossibleContainers(matchedNames, uniqueBagNames)
  } else {
    return matchedNames.Sort().Unique()
  }
}
