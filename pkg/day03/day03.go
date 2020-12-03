package day03

import (
  "fmt"
  "strings"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 3)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  slope := data()
  count := slope.Count(3, 1)
  return count
}

func Puzzle2() int {
  slope := data()
  c0 := slope.Count(1, 1)
  c1 := slope.Count(3, 1)
  c2 := slope.Count(5, 1)
  c3 := slope.Count(7, 1)
  c4 := slope.Count(1, 2)
  return c0 * c1 * c2 * c3 * c4
}


// ========== PRIVATE FNS =================================

func data () Slope  {
  rows   := reader.Lines("./data/day03/input.txt")
  length := len(rows)
  period := len(strings.Split(rows[0], ""))
  trees  := make(map[string]Coord)
  for y, row := range rows {
     cols := strings.Split(row, "")
     for x, col := range cols {
       if col == "#" {
         c := Coord{x: x, y: y}
         trees[c.Key()] = c
       }
     }
  }

  return Slope{trees: trees, length: length, period: period}
}
