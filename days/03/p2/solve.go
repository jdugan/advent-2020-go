package p2

import (
  "../helper/wire"
)


// ========== PUBLIC FNS ==================================

func Solve(wires []map[string]int) int {
  shortest := 0
  wire0    := wires[0]
  wire1    := wires[1]
  coords   := wire.FindIntersections(wire0, wire1)

  for _, coord := range coords {
    sdist := wire0[coord] + wire1[coord]

    if shortest == 0 || sdist < shortest {
      shortest = sdist
    }
  }

  return shortest
}
