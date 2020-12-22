package day23

import (
  "fmt"

  "github.com/elliotchance/pie/pie"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 23)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  return -1
}

func Puzzle2() int {
  return -2
}


// ========== PRIVATE FNS =================================

func data () pie.Strings {
  lines := reader.Lines("./data/day23/input.txt")

  return lines
}
