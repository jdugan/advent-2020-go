package day19

import (
  "fmt"
  "regexp"
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

  return matchCount(msgs, rstr)
}

func Puzzle2() int {
  rules, msgs := data()
  rules[8]     = "42 | 42 8"
  rules[11]    = "42 31 | 42 11 31"
  rules        = rules.Expand()
  rules[8]     = pie.Strings{"(", rules[42], "+)"}.Join("")

  sum := 0
  for i := 1; i < 20; i++ {
    s := strconv.Itoa(i)

    rules[11]    = pie.Strings{"((", rules[42], "){", s, "}(", rules[31], "){", s, "})"}.Join("")
    rules[0]     = pie.Strings{"(", rules[8], rules[11], ")"}.Join("")

    rstr  := pie.Strings{"^", rules[0], "$"}.Join("")
    count := matchCount(msgs, rstr)
    sum    = sum + count
  }

  return sum
}


// ========== PRIVATE FNS =================================

func matchCount (msgs []string, rstr string) int {
  count := 0

  for _, msg := range msgs {
    matched, _ := regexp.MatchString(rstr, msg)
    if matched {
      count = count + 1
    }
  }

  return count
}

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
