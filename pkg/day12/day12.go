package day12

import (
  "fmt"
  "strconv"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 12)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  route  := commands()
  origin := Coord{x: 0, y:0, heading: "E"}
  ferry  := Ferry{position: origin, commands: route}

  return ferry.Move()
}

func Puzzle2() int {
  route    := commands()
  origin   := Coord{x: 0, y:0, heading: "E"}
  waypoint := Coord{x: 10, y: 1, heading: ""}
  ferry    := Ferry{position: origin, waypoint: waypoint, commands: route}

  return ferry.Navigate()
}


// ========== PRIVATE FNS =================================

func commands () []Command {
  lines := reader.Lines("./data/day12/input.txt")

  cmds  := make([]Command, 0)
  for _, line := range lines {
    dir     := string(line[0])
    tmp     := string(line[1:])
    dist, _ := strconv.Atoi(tmp)

    cmds = append(cmds, Command{dir: dir, dist: dist})
  }

  return cmds
}
