package day25

import (
  "fmt"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 25)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  card, door := data()
  enckey     := card.Encrypt(door.pubkey)

  return enckey
}

func Puzzle2() int {
  return 50   // no puzzle to solve; woohoo!
}


// ========== PRIVATE FNS =================================

func data () (Device, Device) {
  lines := reader.Lines("./data/day25/input.txt").Ints()

  card      := Device{pubkey: lines[0]}
  cls       := card.Decrypt(7)
  card.loops = cls

  door      := Device{pubkey: lines[1]}
  dls       := door.Decrypt(7)
  door.loops = dls

  return card, door
}
