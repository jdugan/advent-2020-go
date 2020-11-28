package p1

import (
  "../../../util/slice"
  "../helper/wire"
)


// ========== PUBLIC FNS ==================================

func Solve(wires []map[string]int) int {
  shortest := 0
  coords   := wire.FindIntersections(wires[0], wires[1])

  for _, coord := range coords {
    dims  := slice.CastListToInts(coord)
    mdist := abs(dims[0]) + abs(dims[1])

    if shortest == 0 || mdist < shortest {
      shortest = mdist
    }
  }

  return shortest
}


// ========== PRIVATE FNS =================================

func abs (n int) int {
  if n < 0 {
    return -n
  } else {
    return n
  }
}
