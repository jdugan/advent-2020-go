package slice

import (
  "fmt"
  "strconv"
  "strings"
)


// ========== PUBLIC FUNCTIONS ============================

func CastListToInts (str string) []int {
  strs := strings.Split(str, ",")
  return CastStringsToInts(strs)
}

func CastListToStrings (str string) []string {
  strs := strings.Split(str, ",")
  for index, str := range strs {
    strs[index] = strings.TrimSpace(str)
  }
  return strs
}

func CastStringsToInts (strs []string) []int {
  ints := make([]int, len(strs))
  for index, str := range strs {
    ints[index] = castStringToInt(str)
  }
  return ints
}


// ========== PRIVATE FUNCTIONS ===========================

// Atoi trims whitespace by default :)
//
func castStringToInt (str string) int {
  i, err := strconv.Atoi(str)
  if err != nil {
    fmt.Println(err)
  }
  return i
}
