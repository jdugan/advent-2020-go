package day02

import (
  "fmt"
  "strings"

  "github.com/elliotchance/pie/pie"

  "../../pkg/computer"
  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 2)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1 () int {
  return runScenario(initialMemory(), 12, 2)
}

func Puzzle2 () int {
  const required = 19690720
  initial       := initialMemory()
  found         := false
  noun          := 0
  verb          := 0

  for noun = 0; noun < 100; noun++ {
    for verb = 0; verb < 100; verb++ {
      memory := make([]int, len(initial))
      copy(memory, initial)
      memory = pie.Ints(memory)

      solution := runScenario(memory, noun, verb)
      if solution == required {
        found = true
        break
      }
    }
    if found == true {
      break
    }
  }

  return (noun * 100) + verb
}


// ========== PRIVATE FNS =================================

func initialMemory () pie.Ints {
  lines := reader.Lines("./data/day02/input.txt")
  strs  := strings.Split(lines[0], ",")
  return pie.Strings(strs).Ints()
}

func runScenario (memory pie.Ints, noun int, verb int) int {
  memory[1] = noun
  memory[2] = verb

  computerFn := computer.Create(memory, 1000)
  memory      = computerFn()

  return memory[0]
}
