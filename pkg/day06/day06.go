package day06

import (
  "fmt"

  "github.com/elliotchance/pie/pie"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 6)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  sum := 0
  for _, g := range groups() {
    da := g.DistinctAnswers()
    sum = sum + len(da)
  }
  return sum
}

func Puzzle2() int {
  sum := 0
  for _, g := range groups() {
    ua := g.UnanimousAnswers()
    sum = sum + len(ua)
  }
  return sum
}


// ========== PRIVATE FNS =================================

func groups () []Group {
  lines   := reader.Lines("./data/day06/input.txt")
  groups  := make([]Group, 0)
  answers := make([]string, 0)

  for _, line := range lines {
    if line == "" {
      groups  = append(groups, Group{answers: pie.Strings(answers)})
      answers = make([]string, 0)
    } else {
      answers = append(answers, line)
    }
  }
  groups  = append(groups, Group{answers: pie.Strings(answers)})

  return groups
}
