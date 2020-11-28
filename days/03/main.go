package main

import (
  "../../util/reader"
  "../../util/slice"
  "./helper/wire"
  "./p1"
  "./p2"
  "fmt"
)


// ========== RUNNER ======================================

func main() {
  lines := reader.Lines("./data/input.txt")
  wires := make([]map[string]int, len(lines))
  for index, line := range lines {
    instructions := slice.CastListToStrings(line)
    wires[index]  = wire.BuildPath(instructions)
  }

  fmt.Println(" ")
  fmt.Println("DAY", 3)
  fmt.Println(" ")
  fmt.Println("Puzzle 1", "=>", p1.Solve(wires))
  fmt.Println("Puzzle 2", "=>", p2.Solve(wires))
  fmt.Println(" ")
}
