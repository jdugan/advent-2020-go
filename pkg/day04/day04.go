package day04

import (
  "fmt"
  "strings"

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
  count := 0
  for _, p := range passports() {
    if p.ValidateKeys() {
      count = count + 1
    }
  }
  return count
}

func Puzzle2() int {
  count := 0
  for _, p := range passports() {
    if p.ValidateFormats() {
      count = count + 1
    }
  }
  return count
}


// ========== PRIVATE FNS =================================

func parse (str string) map[string]string {
  str     = strings.TrimSpace(str)
  pairs  := strings.Split(str, " ")
  fields := make(map[string]string)

  for _, pair := range pairs {
    parts := strings.Split(pair, ":")
    k := parts[0]
    v := parts[1]
    fields[k] = v
  }

  return fields
}

func passports () []Passport {
  lines     := reader.Lines("./data/day04/input.txt")
  passports := make([]Passport, 0)
  fields    := ""

  for _, line := range lines {
    if line == "" {
      passports = append(passports, Passport{fields: parse(fields)})
      fields = ""
    } else {
      fields = fields + " " + line
    }
  }
  passports = append(passports, Passport{fields: parse(fields)})

  return passports
}
