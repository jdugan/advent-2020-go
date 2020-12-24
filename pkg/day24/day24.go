package day24

import (
  "fmt"
  "strings"

  "github.com/elliotchance/pie/pie"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 24)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  grid, instrs := data()
  grid = grid.Install(instrs)

  return grid.Count()
}

func Puzzle2() int {
  grid, instrs := data()
  grid = grid.Install(instrs)
  grid = grid.Rotate(100)

  return grid.Count()
}


// ========== PRIVATE FNS =================================

func data () (Grid, []pie.Strings) {
  lines  := reader.Lines("./data/day24/input.txt")

  instrs := make([]pie.Strings, 0)
  for _, line := range lines {
    instrs = append(instrs, parseLine(line))
  }

  return Grid{}, instrs
}

func parseLine (line string) pie.Strings {
  ps    := pie.Strings{}
  chars := strings.Split(line, "")

  i := 0
  for i < len(chars) {
    if chars[i] == "n" || chars[i] == "s" {
      ps = append(ps, chars[i] + chars[i+1])
      i = i + 2
    } else {
      ps = append(ps, chars[i])
      i = i + 1
    }
  }

  return ps
}
