package day11

import (
  "fmt"
  "strings"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 11)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  area := baseSeatingArea().SetAdjacentSeatIds()
  prev := -1
  curr := 0

  for prev != curr {
    prev = curr
    area = area.IterateByAdjacency()
    curr = area.OccupiedSeatCount()
  }

  return curr
}

func Puzzle2() int {
  area := baseSeatingArea().SetVisibleIds()
  prev := -1
  curr := 0

  for prev != curr {
    prev = curr
    area = area.IterateByVisibility()
    curr = area.OccupiedSeatCount()
  }

  return curr
}


// ========== PRIVATE FNS =================================

func baseSeatingArea () SeatingArea {
  lines  := reader.Lines("./data/day11/input.txt")

  seats  := make(map[string]Seat)
  width  := len(strings.Split(lines[0], ""))
  height := len(lines)

  for y, row := range lines {
    cols := strings.Split(row, "")
    for x, col := range cols {
      if col != "." {
        s := Seat{x: x + 1, y: y + 1, occupied: col == "#" }
        seats[s.Id()] = s
      }
    }
  }

  return SeatingArea{seats: seats, width: width, height: height}
}
