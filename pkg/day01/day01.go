package day01

import (
  "fmt"

  "github.com/elliotchance/pie/pie"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 1)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  return data().Map(simpleFuel).Sum()
}

func Puzzle2() int {
  return data().Map(totalFuel).Sum()
}


// ========== PRIVATE FNS =================================

func data () pie.Ints {
  return reader.Lines("./data/day01/input.txt").Ints()
}

func simpleFuel (mass int) int {
  f := mass/3 - 2
  return f
}

func totalFuel (mass int) int {
  sum := 0
  for mass > 0 {
    f := simpleFuel(mass)
    if f > 0 {
      sum  = sum + f
    }
    mass = f
  }
  return sum
}
