package day08

import (
  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Program struct {
  commands    []Command
}


// ========== RECEIVERS ===================================

func (p Program) Repair () (int, bool) {
  final := 0
  cmds  := p.commands
  cmds1 := make([]Command, len(cmds))

  for i, cmd := range cmds {
    copy(cmds1, cmds)
    cmds1[i] = cmd.Invert()

    program      := Program{commands: cmds1}
    val, success := program.Run()

    if success {
      final = val
      break
    }
  }

  return final, final > 0
}

func (p Program) Run () (int, bool) {
  visited    := pie.Ints(make([]int, 0))
  terminus   := len(p.commands)
  val        := 0

  error := false
  pos   := 0

  for !error && pos != terminus {
    visited = append(visited, pos)
    cmd    := p.commands[pos]

    switch cmd.operation {
    case "acc":
      val = val + cmd.argument
      pos = pos + 1
    case "jmp":
      pos = pos + cmd.argument
    case "nop":
      pos = pos + 1
    }

    if visited.Contains(pos) {
      error = true
    }
  }

  return val, !error
}
