package day23

import (
  "fmt"
  "strings"

  "github.com/elliotchance/pie/pie"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 23)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  id, cups := basicSetup() // change this
  game     := Game{cups: cups}
  game      = game.Shuffle(id, 100)

  return game.CheckSum()
}

func Puzzle2() int {
  id, cups := complexSetup() // change this
  game     := Game{cups: cups}
  game      = game.Shuffle(id, 10000000)

  return game.FinalScore()
}


// ========== PRIVATE FNS =================================

func basicSetup () (int, CupLinkedList) {
  ids       := seeds()
  id, clist := setup(ids)

  return id, clist
}

func complexSetup () (int, CupLinkedList) {
  ids       := seeds()
  ids        = extend(ids, 1000000)
  id, clist := setup(ids)

  return id, clist
}


// ---------- HELPERS -------------------------------------

func extend (ids pie.Ints, max int) pie.Ints {
  for i := len(ids); i < max; i++ {
    id  := i + 1
    ids  = append(ids, id)
  }

  return ids
}

func seeds () pie.Ints {
  lines := reader.Lines("./data/day23/input.txt")
  ids   := pie.Strings(strings.Split(lines[0], "")).Ints()

  return ids
}

func setup (ids pie.Ints) (int, CupLinkedList) {
  clist := CupLinkedList{}
  for _, id := range ids {
    clist[id] = Cup{id: id}
  }

  prevId := ids.Last()
  for _, id := range ids {
    prev            := clist[prevId]
    prev.clockwiseId = id
    clist[prevId]    = prev

    prevId = id
  }

  return ids.First(), clist
}
