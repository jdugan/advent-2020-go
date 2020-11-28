package main

import (
  "../../util/reader"
  "../../util/slice"
  "./p1"
  "./p2"
  "fmt"
)


// ========== RUNNER ======================================

func main() {
  lines := utils.Lines("./data/input.txt")
  data  := utils.CastStringsToInts(lines)

  fmt.Println(" ")
  fmt.Println("DAY", 0)
  fmt.Println(" ")
  fmt.Println("Puzzle 1", "=>", p1.Solve(data))
  fmt.Println("Puzzle 2", "=>", p2.Solve(data))
  fmt.Println(" ")
}
