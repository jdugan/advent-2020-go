package day19

import (
  "fmt"
  "strconv"
  "strings"

  "github.com/elliotchance/pie/pie"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 19)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  rules, msgs := data()
  rules        = rules.Expand()
  rstr        := pie.Strings{"^", rules[0], "$"}.Join("")
  analyser    := Analyser{messages: msgs}

  return analyser.CountMatches(rstr)
}

func Puzzle2() int {
  rules, msgs := data()
  analyser    := Analyser{messages: msgs}
  rules        = analyser.CorrectedRules(rules)

  return analyser.CountCorrectedMatches(rules)
}


// ========== PRIVATE FNS =================================

func data () (RuleSet, pie.Strings) {
  lines := reader.Lines("./data/day19/input.txt")

  rules := RuleSet{}
  msgs  := pie.Strings{}
  found := false

  for _, line := range lines {
    if found {
      msgs = append(msgs, line)
    } else {
      if line == "" {
        found = true
      } else {
        parts  := strings.Split(line, ": ")
        key, _ := strconv.Atoi(parts[0])
        val    := parts[1]
        val     = strings.Replace(val, "\"", "", -1)

        rules[key] = val
      }
    }
  }

  return rules, msgs
}
