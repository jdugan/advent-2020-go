package day09

import (
  "fmt"

  "github.com/elliotchance/pie/pie"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 9)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  preamble, numbers := data()
  safecracker       := Safecracker{preamble: preamble, numbers: numbers}

  return safecracker.Probe()
}

func Puzzle2() int {
  preamble, numbers := data()
  safecracker       := Safecracker{preamble: preamble, numbers: numbers}

  return safecracker.Crack()
}


// ========== PRIVATE FNS =================================

func data () (pie.Ints, pie.Ints) {
  numbers  := reader.Lines("./data/day09/input.txt").Ints()
  preamble := pie.Ints(make([]int, 25))
  head     := 0

  for i := 0; i < 25; i++ {
    head, numbers = numbers.Shift()
    preamble[i] = head
  }

  return preamble, numbers
}
