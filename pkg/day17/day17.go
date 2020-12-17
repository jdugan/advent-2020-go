package day17

import (
  "fmt"
  "strings"

  // "github.com/elliotchance/pie/pie"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 17)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  ps := basePowerSource()

  return ps.Run(6, 3)
}

func Puzzle2() int {
  ps := basePowerSource()

  return ps.Run(6, 4)
}


// ========== PRIVATE FNS =================================

func basePowerSource () PowerSource {
  lines := reader.Lines("./data/day17/input.txt")
  cubes := make(map[string]Cube)

  for y, row := range lines {
    cols := strings.Split(row, "")
    for x, col := range cols {
      if col == "#" {
        coord := Coord{x: x, y: y, z: 0, w:0}
        cube  := Cube{coord: coord, state: true}
        cubes[cube.Id()] = cube
      }
    }
  }

  return PowerSource{cubes: cubes}
}
