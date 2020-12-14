package day13

import (
  "fmt"
  "strconv"
  "strings"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 13)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  etd, fleet := data()

  return fleet.CheckSum(etd)
}

func Puzzle2() int {
  _, fleet := data()

  return fleet.PerfectTimestamp()
}


// ========== PRIVATE FNS =================================

func data () (int, Fleet) {
  lines  := reader.Lines("./data/day13/input.txt")

  etd, _ := strconv.Atoi(lines[0])
  bids   := strings.Split(lines[1], ",")
  fleet  := Fleet{}

  for i, bid := range bids {
    if bid != "x" {
      id, _  := strconv.Atoi(bid)
      offset := i

      if i > id {
        offset = (i % id)
      }
      if offset == 0 {
        offset = id
      }

      bus  := Bus{id: id, offset: offset}
      fleet = append(fleet, bus)
    }
  }

  return etd, fleet
}
