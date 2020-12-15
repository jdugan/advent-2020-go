package day15

import (
  "fmt"
  "strings"

  "github.com/elliotchance/pie/pie"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 15)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  game := Game{seeds: seeds()}

  return game.Play(2020)
}

func Puzzle2() int {
  game := Game{seeds: seeds()}

  return game.Play(30000000)
}


// ========== PRIVATE FNS =================================

func seeds () pie.Ints {
  lines := reader.Lines("./data/day15/input.txt")
  nums  := pie.Strings(strings.Split(lines[0], ",")).Ints()

  return nums
}
