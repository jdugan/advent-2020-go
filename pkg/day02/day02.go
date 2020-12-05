package day02

import (
  "fmt"
  "regexp"
  "strconv"

  "github.com/elliotchance/pie/pie"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 2)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  count := 0
  for _, line := range data() {
    policy, password := parse(line)
    if policy.OldMatch(password) {
      count = count + 1
    }
  }
  return count
}

func Puzzle2() int {
  count := 0
  for _, line := range data() {
    policy, password := parse(line)
    if policy.CurrentMatch(password) {
      count = count + 1
    }
  }
  return count
}


// ========== PRIVATE FNS =================================

func data () pie.Strings {
  return reader.Lines("./data/day02/input.txt")
}

func parse (line string) (policy Policy, password string) {
  pattern, _ := regexp.Compile(`^(\d+)-(\d+) ([a-z]): (\w+)$`)
  elements   := pattern.FindStringSubmatch(line)[1:]

  min, _   := strconv.Atoi(elements[0])
  max, _   := strconv.Atoi(elements[1])
  policy    = Policy{min: min, max: max, letter: elements[2]}
  password  = elements[3]

  return policy, password
}
