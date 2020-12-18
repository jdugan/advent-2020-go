package day18

import (
  "fmt"

  "github.com/elliotchance/pie/pie"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 18)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  sum := 0
  for _, eq := range equations() {
    sum = sum + NewMath{}.Solve(eq)
  }

  return sum
}

func Puzzle2() int {
  sum := 0
  for _, eq := range equations() {
    sum = sum + AdvancedMath{}.Solve(eq)
  }

  return sum
}


// ========== PRIVATE FNS =================================

func equations () pie.Strings {
  lines := reader.Lines("./data/day18/input.txt")

  return lines
}
