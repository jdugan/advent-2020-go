package day11

import (
  "fmt"

  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Seat struct {
  x             int
  y             int
  occupied      bool

  adjacentIds   pie.Strings
  visibleSlopes []pie.Ints
}


// ========== RECEIVERS ===================================

func (s Seat) ReassignByAdjacentSeats (sa SeatingArea) Seat {
  s1   := s.Copy()
  oasc := s.OccupiedAdjacentSeatCount(sa)
  if s.occupied {
    if oasc > 3 {
      s1.occupied = false
    }
  } else {
    if oasc == 0 {
      s1.occupied = true
    }
  }
  return s1
}

func (s Seat) ReassignByVisibleSeats (sa SeatingArea) Seat {
  s1   := s.Copy()
  oasc := s.OccupiedVisibleSeatCount(sa)
  if s.occupied {
    if oasc > 4 {
      s1.occupied = false
    }
  } else {
    if oasc == 0 {
      s1.occupied = true
    }
  }
  return s1
}


// ---------- ADJACENT HELPERS ----------------------------

func (s Seat) AdjacentSeatIds (sa SeatingArea) pie.Strings {
  ids := make([]string, 0)
  ids  = append(ids, fmt.Sprintf("%d,%d", s.x - 1, s.y - 1))    // northwest
  ids  = append(ids, fmt.Sprintf("%d,%d", s.x, s.y - 1))        // north
  ids  = append(ids, fmt.Sprintf("%d,%d", s.x + 1, s.y - 1))    // northeast
  ids  = append(ids, fmt.Sprintf("%d,%d", s.x + 1, s.y))        // east
  ids  = append(ids, fmt.Sprintf("%d,%d", s.x + 1, s.y + 1))    // southeast
  ids  = append(ids, fmt.Sprintf("%d,%d", s.x, s.y + 1))        // south
  ids  = append(ids, fmt.Sprintf("%d,%d", s.x - 1, s.y + 1))    // southwest
  ids  = append(ids, fmt.Sprintf("%d,%d", s.x - 1, s.y))        // west

  asids := pie.Strings{}
  for _, id := range ids {
    _, exists := sa.seats[id]
    if exists {
      asids = append(asids, id)
    }
  }

  return asids
}

func (s Seat) OccupiedAdjacentSeatCount (sa SeatingArea) int {
  count := 0
  for _, id := range s.adjacentIds {
    s, exists := sa.seats[id]
    if exists && s.occupied {
      count = count + 1
    }
  }
  return count
}


// ---------- VISIBLE HELPERS -----------------------------

func (s Seat) VisibleSlopes (sa SeatingArea) []pie.Ints {
  slopes := make([]pie.Ints, 0)
  slopes  = append(slopes, pie.Ints{-1, -1})      // northwest
  slopes  = append(slopes, pie.Ints{0, -1})       // north
  slopes  = append(slopes, pie.Ints{1, -1})       // northeast
  slopes  = append(slopes, pie.Ints{1, 0})        // east
  slopes  = append(slopes, pie.Ints{1, 1})        // southeast
  slopes  = append(slopes, pie.Ints{0, 1})        // south
  slopes  = append(slopes, pie.Ints{-1, 1})       // southwest
  slopes  = append(slopes, pie.Ints{-1, 0})       // west

  vslopes := make([]pie.Ints, 0)

  for _, slope := range slopes {
    tmp  := Seat{}
    halt := false
    iter := 1

    for !halt {
      tmp.x = s.x + (iter * slope[0])
      tmp.y = s.y + (iter * slope[1])

      if sa.IsSeatIdWithinArea(tmp) {
        id        := tmp.Id()
        _, exists := sa.seats[id]
        if exists {                         // visible seat; save and stop
          vslopes = append(vslopes, slope)
          halt    = true
        } else {                            // open floor; continue
          iter = iter + 1
        }
      } else {                              // off the board; stop
        halt = true
      }
    }
  }

  return vslopes
}

func (s Seat) OccupiedVisibleSeatCount (sa SeatingArea) int {
  count := 0

  for _, slope := range s.visibleSlopes {
    tmp  := Seat{}
    halt := false
    iter := 1

    for !halt {
      tmp.x = s.x + (iter * slope[0])
      tmp.y = s.y + (iter * slope[1])

      if sa.IsSeatIdWithinArea(tmp) {
        id         := tmp.Id()
        s1, exists := sa.seats[id]
        if exists {                 // visible seat; check and stop
          if s1.occupied {
            count = count + 1       // occupied visible seat; add
          }
          halt = true
        } else {                    // open floor; continue
          iter = iter + 1
        }
      } else {                      // off the board; stop
        halt = true
      }
    }
  }

  return count
}


// ---------- UTILITIES -----------------------------------

func (s Seat) Copy () Seat {
  return Seat{x:              s.x,
              y:              s.y,
              occupied:       s.occupied,
              adjacentIds:    s.adjacentIds,
              visibleSlopes:  s.visibleSlopes}
}

func (s Seat) Id () string {
  return fmt.Sprintf("%d,%d", s.x, s.y)
}
