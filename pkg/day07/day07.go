package day07

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
  fmt.Println("DAY", 7)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  lh      := LuggageHandler{ruleset: ruleset()}
  matches := lh.CountPossibleContainers(pie.Strings{}, pie.Strings{"shiny_gold"})

  return len(matches)
}

func Puzzle2() int {
  lh   := LuggageHandler{ruleset: ruleset()}
  root := BagGroup{quantity: 1, name: "shiny_gold"}
  sum  := lh.CountEnclosedBags(root)

  return sum - 1    // don't include the root bag!
}


// ========== PRIVATE FNS =================================

// Builds a map where the keys are underscored bag type
// names (e.g., shiny_gold) and the values are an array
// of bag groups (e.g., 7 faded_green, 3 bright_blue).
//
func ruleset () map[string]BagGroups {
  lines := reader.Lines("./data/day07/input.txt")
  rules := make(map[string]BagGroups)

  for _, line := range lines {
    re0 := regexp.MustCompile(` (bags|bag)`)
    line = re0.ReplaceAllString(line, "")

    re1      := regexp.MustCompile(`^(.+) contain (.+)\.$`)
    elements := re1.FindStringSubmatch(line)[1:]

    bagName := strings.Replace(elements[0], " ", "_", -1)
    list    := strings.Replace(elements[1], "no other", "", -1)
    items   := make([]string, 0)
    if list != "" {
      items = strings.Split(list, ", ")
    }

    bagGroups := make(BagGroups, 0)
    re2       := regexp.MustCompile(`(\d+) (.+)`)
    for _, item := range items {
      parts := re2.FindStringSubmatch(item)[1:]
      qstr        := parts[0]
      quantity, _ := strconv.Atoi(qstr)
      name        := strings.Replace(parts[1], " ", "_", -1)
      bagGroups    = append(bagGroups, BagGroup{quantity: quantity, name: name})
    }

    rules[bagName] = bagGroups
  }

  return rules
}
