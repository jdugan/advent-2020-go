package day14

import (
  "fmt"

  "github.com/elliotchance/pie/pie"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 14)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  computer, cmds := data()
  memsum         := computer.Initialise(cmds)

  return memsum
}

func Puzzle2() int {
  computer, cmds := data()
  memsum         := computer.Decode(cmds)

  return memsum
}


// ========== PRIVATE FNS =================================

func data () (Computer, pie.Strings) {
  memory   := make(map[int]int)
  computer := Computer{memory: memory}
  cmds     := reader.Lines("./data/day14/input.txt")

  return computer, cmds
}
