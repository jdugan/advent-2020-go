package day10

import (
  "fmt"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 10)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  analyser := Analyser{adapters: adapters()}

  return analyser.CheckSum()
}

func Puzzle2() int {
  analyser := Analyser{adapters: adapters()}

  return analyser.PermutationCount()
}


// ========== PRIVATE FNS =================================

func adapters () []Adapter {
  joltages := reader.Lines("./data/day10/input.txt").Ints().Sort()

  devices   := make([]Adapter, joltages.Len() + 1)
  devices[0] = Adapter{joltage: 0}      // wall outlet

  for i, joltage := range joltages {
    devices[i + 1] = Adapter{joltage: joltage}
  }

  return devices
}
