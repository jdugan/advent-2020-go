package day19

import (
  "regexp"
  "strconv"

  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITIONS =================================

type Analyser struct {
  messages  pie.Strings
}


// ========== RECEIVERS ===================================

func (a Analyser) CountMatches (rstr string) int {
  count := 0
  for _, msg := range a.messages {
    matched, _ := regexp.MatchString(rstr, msg)
    if matched {
      count = count + 1
    }
  }
  return count
}

func (a Analyser) CountCorrectedMatches (rs RuleSet) int {
  sum  := 0
  iter := 1
  halt := false

  for !halt {
    size  := strconv.Itoa(iter)
    rs[11] = pie.Strings{"((", rs[42], "){",size, "}(", rs[31], "){", size, "})"}.Join("")
    rs[0]  = pie.Strings{"(", rs[8], rs[11], ")"}.Join("")

    rstr  := pie.Strings{"^", rs[0], "$"}.Join("")
    count := a.CountMatches(rstr)

    if count == 0 {     // this is an assumption, but it
      halt = true       // is a correct one. :)
    } else {
      sum  = sum + count
      iter = iter + 1
    }
  }

  return sum
}


// ---------- OVERRIDE HELPERS ----------------------------

func (a Analyser) CorrectedRules (rs RuleSet) RuleSet {
  rs[8]     = "42 | 42 8"
  rs[11]    = "42 31 | 42 11 31"
  rs        = rs.Expand()
  rs[8]     = pie.Strings{"(", rs[42], "+)"}.Join("")

  return rs
}
