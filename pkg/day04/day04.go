package day04

import (
  "fmt"
  "sort"
  "strconv"
  "strings"

  "github.com/elliotchance/pie/pie"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 4)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  limits := data()
  low    := limits.First()
  high   := limits.Last()
  count  := 0

  for i := low; i <= high; i++ {
    if hasSortedDigits(i) && hasRepeatedDigits(i) {
      count = count + 1
    }
  }

  return count
}

func Puzzle2() int {
  limits := data()
  low    := limits.First()
  high   := limits.Last()
  count  := 0

  for i := low; i <= high; i++ {
    if hasSortedDigits(i) && hasDoubledDigits(i) {
      count = count + 1
    }
  }

  return count
}


// ========== PRIVATE FNS =================================

func data () pie.Ints {
  lines := reader.Lines("./data/day04/input.txt")
  strs  := strings.Split(lines[0], ",")
  return pie.Strings(strs).Ints()
}

func hasDoubledDigits (n int) bool {
  str    := strconv.Itoa(n)
  digits := strings.Split(str, "")
  dmap   := make(map[string]int)

  for _, d := range digits {
    c := dmap[d] + 1
    dmap[d] = c
  }

  found := false
  for _, v := range dmap {
    if v == 2 {
      found = true
      break
    }
  }
  return found
}

func hasRepeatedDigits (n int) bool {
  str    := strconv.Itoa(n)
  digits := strings.Split(str, "")
  dmap   := make(map[string]int)

  found  := false
  for _, d := range digits {
    c := dmap[d] + 1
    dmap[d] = c
    if c > 1 {
      found = true
      break
    }
  }

  return found
}

func hasSortedDigits (n int) bool {
  str    := strconv.Itoa(n)
  digits := strings.Split(str, "")
  return sort.StringsAreSorted(digits)
}
