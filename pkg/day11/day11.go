package day11

import (
  "fmt"
  "strings"

  "github.com/elliotchance/pie/pie"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================
// TODO: Refactor both methods for performance?
// They take around 1,5s combined so far. :(
//

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
    curr = area.TotalOccupiedSeatCount()
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
    curr = area.TotalOccupiedSeatCount()
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

  return SeatingArea{seats: seats, width: width, height: height, slopes: baseSlopes()}
}

func baseSlopes () []pie.Ints {
  slopes := make([]pie.Ints, 8)
  slopes[0] = pie.Ints{-1, -1}      // northwest
  slopes[1] = pie.Ints{0, -1}       // north
  slopes[2] = pie.Ints{1, -1}       // northeast
  slopes[3] = pie.Ints{1, 0}        // east
  slopes[4] = pie.Ints{1, 1}        // southeast
  slopes[5] = pie.Ints{0, 1}        // south
  slopes[6] = pie.Ints{-1, 1}       // southwest
  slopes[7] = pie.Ints{-1, 0}       // west

  return slopes
}
