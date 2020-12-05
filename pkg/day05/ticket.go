package day05

import (
  "fmt"
  "regexp"
  "strconv"
)


// ========== DEFINITION ==================================

type Ticket struct {
  code  string
}


// ========== RECEIVERS ===================================

func (t Ticket) Format () string {
  return fmt.Sprintf("%s: id %d", t.code, t.Id())
}

func (t Ticket) Id () int {
  re0 := regexp.MustCompile(`[F|L]`)
  re1 := regexp.MustCompile(`[B|R]`)

  code := t.code
  code  = re0.ReplaceAllString(code, "0")
  code  = re1.ReplaceAllString(code, "1")

  pos, _ := strconv.ParseInt(code, 2, 32)
  return int(pos)
}
