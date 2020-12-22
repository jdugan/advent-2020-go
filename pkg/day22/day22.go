package day22

import (
  "fmt"
  "strings"

  "github.com/elliotchance/pie/pie"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 22)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  d1, d2 := decks()
  game   := Game{d1: d1, d2: d2}
  result := game.Play()

  return result.score
}

func Puzzle2() int {
  d1, d2 := decks()
  game   := Game{d1: d1, d2: d2}
  result := game.PlayRecursively()

  return result.score
}


// ========== PRIVATE FNS =================================

func decks () (pie.Ints, pie.Ints) {
  lines := reader.Lines("./data/day22/input.txt")
  d1    := parseLine(lines[0])
  d2    := parseLine(lines[1])

  return d1, d2
}

func parseLine (line string) pie.Ints {
  cards := pie.Strings(strings.Split(line, ","))
  return cards.Ints()
}
