package computer

import (
  "fmt"

  "github.com/elliotchance/pie/pie"
)


// ========== PUBLIC FUNCTIONS ============================

// This function returns a closure that represents
// a runnable int computer.
//
func Create (initial pie.Ints, limit int) func() pie.Ints {
  return func () pie.Ints {
    memory := make([]int, len(initial))
    copy(memory, initial)
    memory = pie.Ints(memory)

    halt  := false
    loops := 0
    pos   := 0

    for !halt && loops < limit {
      opcode := memory[pos]

      switch opcode {
      case 1:
        memory, pos = opcodeAddByAddress(memory, pos)
      case 2:
        memory, pos = opcodeMultiplyByAddress(memory, pos)
      case 99:
        halt = true
      default:
        halt = true
        fmt.Println("ERROR: Opcode", opcode, "not recognised.")
      }

      loops = loops + 1
    }

    return memory
  }
}


// ========== PRIVATE FUNCTIONS ===========================

func opcodeAddByAddress (memory []int, pos int) ([]int, int) {
  paddr1 := memory[pos + 1]
  paddr2 := memory[pos + 2]
  oaddr  := memory[pos + 3]

  memory[oaddr] = memory[paddr1] + memory[paddr2]
  pos = pos + 4

  return memory, pos
}

func opcodeMultiplyByAddress (memory []int, pos int) ([]int, int) {
  paddr1 := memory[pos + 1]
  paddr2 := memory[pos + 2]
  oaddr  := memory[pos + 3]

  memory[oaddr] = memory[paddr1] * memory[paddr2]
  pos = pos + 4

  return memory, pos
}
