package day18

import (
  "regexp"
  "strconv"
  "strings"

  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type NewMath struct {}


// ========== RECEIVERS ===================================

func (nm NewMath) Solve (eq string) int {
  re     := regexp.MustCompile(`(\(\d+( [\+\*] \d+)+\))`)
  groups := re.FindStringSubmatch(eq)

  if len(groups) > 0 {
    seq   := groups[0]
    mseq  := strings.Replace(seq, "(", "", 1)
    mseq   = strings.Replace(mseq, ")", "", 1)
    sr    := nm.Compute(mseq)
    ssr   := strconv.Itoa(sr)
    eq     = strings.Replace(eq, seq, ssr, 1)

    return nm.Solve(eq)
  } else {
    return nm.Compute(eq)
  }
}

func (nm NewMath) Compute (eq string) int {
  parts            := pie.Strings(strings.Split(eq, " "))
  allOperands      := parts.FilterNot(func (s string) bool { return s == "+" || s == "*" }).Ints()
  operators        := parts.Filter(func (s string) bool { return s == "+" || s == "*" })
  result, operands := allOperands.Shift()

  for i := 0; i < operators.Len(); i++ {
    if operators[i] == "+" {
      result = nm.Add(result, operands[i])
    } else {
      result = nm.Multiply(result, operands[i])
    }
  }

  return result
}

// ---------- UTILITIES -----------------------------------

func (nm NewMath) Add (x int, y int) int {
  return x + y
}

func (nm NewMath) Multiply (x int, y int) int {
  return x * y
}
