package day08

import (
  "fmt"
  "strconv"
  "strings"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 8)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  program := Program{commands: commands()}
  val, _  := program.Run()

  return val
}

func Puzzle2() int {
  program := Program{commands: commands()}
  val, _  := program.Repair()

  return val
}


// ========== PRIVATE FNS =================================

func commands () []Command {
  lines := reader.Lines("./data/day08/input.txt")
  cmds  := make([]Command, len(lines))

  for i, line := range lines {
    elements     := strings.Split(line, " ")
    operation    := elements[0]
    argument, _  := strconv.Atoi(elements[1])
    command      := Command{id: i, operation: operation, argument: argument}

    cmds[i] = command
  }

  return cmds
}
