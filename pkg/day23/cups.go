package day23

import (
  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION(S) ===============================

type Cup struct {
  id            int
  clockwiseId   int
}

type CupLinkedList map[int]Cup


// ========== RECEIVERS ===================================

// This function starts with a specified cup id and
// creates an ordered list of the surrounding ids by
// navigating clockwise until returning back to the
// specified cup.
//
func (clist CupLinkedList) DigitsAround (anchorId int) pie.Ints {
  digits := pie.Ints{}
  anchor := clist[anchorId]
  next   := clist[anchor.clockwiseId]

  for next != anchor {
    digits = append(digits, next.id)

    if next.clockwiseId == 0 {
      break
    } else {
      next = clist[next.clockwiseId]
    }
  }

  return digits
}

// This function finds the id of the cup that will
// serve as the insertion point for cups plucked
// during a shuffle.
//
func (clist CupLinkedList) FindDestinationId (anchorId int, pluckedIds pie.Ints) int {
  destId := anchorId
  maxId  := len(clist)

  found  := false
  for !found {
    destId = destId - 1
    if destId < 1 {
      destId = maxId
    }
    if !pluckedIds.Contains(destId) {
      found = true
    }
  }

  return destId
}

// This function inserts a chain of cup ids into
// the larger chain immediately clockwise of a
// specified cup id.
//
func (clist CupLinkedList) InsertChain (anchorId int, chainIds pie.Ints) CupLinkedList {
  lastId := chainIds.Last()

  anchor := clist[anchorId]
  last   := clist[lastId]

  last.clockwiseId   = anchor.clockwiseId
  anchor.clockwiseId = chainIds.First()

  clist[anchorId] = anchor
  clist[lastId]   = last

  return clist
}


// This function starts with a given cup and plucks out
// the next N cups. It does so by collecting the ids
// of cups placed at the specified steps around the ring and
// then linking the given cup to the cup one step clockwise
// from the last collected cup.
//
// The plucked cups remain in the map and remain linked
// in there original order.  The are simply no longer in
// the larger chain.
//
func (clist CupLinkedList) Pluck (anchorId int, size int) (CupLinkedList, pie.Ints) {
  ids    := pie.Ints{}
  currId := anchorId
  for i := 0; i < size; i++ {
    curr  := clist[currId]
    ids    = append(ids, curr.clockwiseId)
    currId = curr.clockwiseId
  }
  lastId := ids.Last()

  anchor            := clist[anchorId]
  last              := clist[lastId]

  anchor.clockwiseId = last.clockwiseId
  last.clockwiseId   = 0

  clist[anchorId] = anchor
  clist[lastId]   = last

  return clist, ids
}
