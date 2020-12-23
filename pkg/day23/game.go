package day23

import (
  "strconv"

  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Game struct {
  cups    CupLinkedList
}


// ========== RECEIVERS ===================================

// This function shuffles the order of the cups by
// plucking three cups out of the rotation and
// inserting them into a different part of the
// larger chain.
//
func (g Game) Shuffle (anchorId int, moves int) Game {
  clist  := g.cups
  currId := anchorId

  for i := 0; i < moves; i++ {
    var ids pie.Ints

    clist, ids  = clist.Pluck(currId, 3)
    destId     := clist.FindDestinationId(currId, ids)
    clist       = clist.InsertChain(destId, ids)

    currId      = clist[currId].clockwiseId
  }

  return Game{cups: clist}
}


// ---------- SCORING HELPERS -----------------------------

// This function produces a checksum by locating
// the cup with id=1 and collecting the remaining ids
// in a clockwise arc. These ids are then concatenated
// and converted to an int.
//
func (g Game) CheckSum () int {
  digits  := g.cups.DigitsAround(1)
  str     := digits.Strings().Join("")
  csum, _ := strconv.Atoi(str)

  return csum
}

func (g Game) FinalScore () int {
  _, ids := g.cups.Pluck(1, 2)

  return ids.Product()
}
