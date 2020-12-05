package day05

import (
  "fmt"

  "github.com/elliotchance/pie/pie"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 5)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  return sortedTicketIds().Last()
}

func Puzzle2() int {
  prev, ids := sortedTicketIds().Shift()

  for _, id := range ids {
    if id - prev == 2 {
      break
    } else {
      prev = id
    }
  }

  return prev + 1
}


// ========== PRIVATE FNS =================================

func tickets () []Ticket {
  codes   := reader.Lines("./data/day05/input.txt")
  tickets := make([]Ticket, len(codes))
  for i, code := range codes {
    tickets[i] = Ticket{code: code}
  }
  return tickets
}

func sortedTicketIds () pie.Ints {
  ts  := tickets()
  ids := pie.Ints(make([]int, len(ts)))
  for i, t := range ts {
    ids[i] = t.Id()
  }
  return ids.Sort()
}
