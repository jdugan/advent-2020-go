package day10

import (
  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Analyser struct {
  adapters  []Adapter
}


// ========== RECEIVERS ===================================

// This function calculates the joltage jumps
// between adapters and maps counts by the jump
// amount. The checksum is calculated from the
// values in the map.
//
func (a Analyser) CheckSum () int {
  dmap     := map[int]int{1: 0, 2: 0, 3: 0}
  adapters := a.adapters[1:]
  prev     := a.adapters[0]

  for _, curr := range adapters {
    diff      := curr.Difference(prev)
    dmap[diff] = dmap[diff] + 1
    prev       = curr
  }
  dmap[3] = dmap[3] + 1

  return dmap[1] * dmap[3]
}

// This function maps the compatible joltage
// jumps available to each adapter. The compatible
// lists are then reduced down to non-linear
// groups and permutation factors are assigned
// based on the length and characteristic of each
// group. Possible permutations are 1, 2, 4, and 7.
// (This algorithm ignores 1 since it does not affect
// the final calculation.) The final result is the
// product of all the significant permutation
// factors.
//
func (a Analyser) PermutationCount () int {
  comps := make([]pie.Ints, 0)
  for i, adapter := range a.adapters {
    possibles := a.adapters[i+1:]
    jolts     := adapter.CompatibleJoltages(possibles)
    comps      = append(comps, jolts)
  }

  // reduce compatible lists down to
  // permutation factors
  factors := pie.Ints{}
  base    := pie.Ints{}
  for _, curr := range comps {
    union := append(base, curr...).Sort().Unique()

    blen := base.Len()
    ilen := base.Intersect(curr).Len()
    ulen := union.Len()

    if ulen > blen {
      if ilen == 0 {
        factors = adjustFactors(factors, base)
        base    = curr
      } else if blen > ilen {
        base = union
      }
    }
  }
  factors = adjustFactors(factors, base)

  // multiply all the factors to get
  // the total permutation
  return factors.Product()
}


// ========== HELPERS =====================================

func adjustFactors (factors pie.Ints, permutations pie.Ints) pie.Ints {
  switch permutations.Len() {
  case 4:
    factors = append(factors, 7)
  case 3:
    factors = append(factors, 4)
  case 2:
    factors = append(factors, 2)
  }
  return factors
}
