package day08


// ========== DEFINITION ==================================

type Command struct {
  id          int
  operation   string
  argument    int
}


// ========== RECEIVERS ===================================

func (c Command) Invert () Command {
  operation := c.operation

  switch operation {
  case "jmp":
    operation = "nop"
  case "nop":
    operation = "jmp"
  }

  return Command{id: c.id, operation: operation, argument: c.argument}
}
