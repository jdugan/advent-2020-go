package day18

import (
  // "fmt"
  "regexp"
  "strconv"
  "strings"

  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type AdvancedMath struct {}


// ========== RECEIVERS ===================================

func (am AdvancedMath) Solve (eq string) int {
  re     := regexp.MustCompile(`(\(\d+( [\+\*] \d+)+\))`)
  groups := re.FindStringSubmatch(eq)

  if len(groups) > 0 {
    seq   := groups[0]
    mseq  := strings.Replace(seq, "(", "", 1)
    mseq   = strings.Replace(mseq, ")", "", 1)
    sr    := am.Compute(mseq)
    ssr   := strconv.Itoa(sr)
    eq     = strings.Replace(eq, seq, ssr, 1)

    return am.Solve(eq)
  } else {
    return am.Compute(eq)
  }
}

func (am AdvancedMath) Compute (eq string) int {
  re     := regexp.MustCompile(`(\d+ \+ \d+)`)
  groups := re.FindStringSubmatch(eq)

  if len(groups) > 0 {
    seq   := groups[0]
    sr    := am.Add(seq)
    ssr   := strconv.Itoa(sr)
    eq     = strings.Replace(eq, seq, ssr, 1)

    return am.Compute(eq)
  } else {
    return am.Multiply(eq)
  }
}


// ---------- UTILITIES -----------------------------------

func (am AdvancedMath) Add (eq string) int {
  operands := pie.Strings(strings.Split(eq, " + ")).Ints()

  return operands.Sum()
}

func (am AdvancedMath) Multiply (eq string) int {
  operands := pie.Strings(strings.Split(eq, " * ")).Ints()

  return operands.Product()
}
