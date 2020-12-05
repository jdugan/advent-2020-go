package day05

import (
  "fmt"
  "strconv"
  "strings"
)


// ========== DEFINITION ==================================

type Ticket struct {
  code  string
}


// ========== RECEIVERS ===================================

func (t Ticket) Format () string {
  return fmt.Sprintf("%s: r %d, s %d, id %d", t.code, t.Row(), t.Seat(), t.Id())
}

func (t Ticket) Id () int {
  return (t.Row() * 8) + t.Seat()
}

func (t Ticket) Row () int {
  code := t.code
  code  = strings.Replace(code, "L", "", -1)
  code  = strings.Replace(code, "R", "", -1)
  code  = strings.Replace(code, "F", "0", -1)
  code  = strings.Replace(code, "B", "1", -1)

  pos, _ := strconv.ParseInt(code, 2, 32)
  return int(pos)
}

func (t Ticket) Seat () int {
  code := t.code
  code  = strings.Replace(code, "F", "", -1)
  code  = strings.Replace(code, "B", "", -1)
  code  = strings.Replace(code, "L", "0", -1)
  code  = strings.Replace(code, "R", "1", -1)

  pos, _ := strconv.ParseInt(code, 2, 32)
  return int(pos)
}
