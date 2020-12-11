package day11

import (
  "fmt"

  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Seat struct {
  x         int
  y         int
  occupied  bool
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

func (s Seat) AdjacentSeatIds () []string {
  ids := make([]string, 0)
  ids  = append(ids, fmt.Sprintf("%d,%d", s.x - 1, s.y - 1))    // northwest
  ids  = append(ids, fmt.Sprintf("%d,%d", s.x, s.y - 1))        // north
  ids  = append(ids, fmt.Sprintf("%d,%d", s.x + 1, s.y - 1))    // northeast
  ids  = append(ids, fmt.Sprintf("%d,%d", s.x + 1, s.y))        // east
  ids  = append(ids, fmt.Sprintf("%d,%d", s.x + 1, s.y + 1))    // southeast
  ids  = append(ids, fmt.Sprintf("%d,%d", s.x, s.y + 1))        // south
  ids  = append(ids, fmt.Sprintf("%d,%d", s.x - 1, s.y + 1))    // southwest
  ids  = append(ids, fmt.Sprintf("%d,%d", s.x - 1, s.y))        // west
  return ids
}

func (s Seat) AdjacentSeats (sa SeatingArea) []Seat {
  as := make([]Seat,0)
  for _, id := range s.AdjacentSeatIds() {
    s, exists := sa.seats[id]
    if exists {
      as = append(as, s)
    }
  }
  return as
}

func (s Seat) OccupiedAdjacentSeats (sa SeatingArea) []Seat {
  as := s.AdjacentSeats(sa)
  os := make([]Seat, 0)
  for _, s := range as {
    if s.occupied {
      os = append(os, s)
    }
  }
  return os
}

func (s Seat) OccupiedAdjacentSeatCount (sa SeatingArea) int {
  return len(s.OccupiedAdjacentSeats(sa))
}


// ---------- VISIBLE HELPERS -----------------------------

func (s Seat) VisibleSlopes() []pie.Ints {
  slopes := make([]pie.Ints, 0)
  slopes  = append(slopes, pie.Ints{-1, -1})      // northwest
  slopes  = append(slopes, pie.Ints{0, -1})       // north
  slopes  = append(slopes, pie.Ints{1, -1})       // northeast
  slopes  = append(slopes, pie.Ints{1, 0})        // east
  slopes  = append(slopes, pie.Ints{1, 1})        // southeast
  slopes  = append(slopes, pie.Ints{0, 1})        // south
  slopes  = append(slopes, pie.Ints{-1, 1})       // southwest
  slopes  = append(slopes, pie.Ints{-1, 0})       // west
  return slopes
}

func (s Seat) VisibleSeats (sa SeatingArea) []Seat {
  vs := make([]Seat, 0)

  for _, slope := range s.VisibleSlopes() {
    tmp  := Seat{}
    halt := false
    iter := 1

    for !halt {
      tmp.x = s.x + (iter * slope[0])
      tmp.y = s.y + (iter * slope[1])

      if sa.IsSeatIdWithinArea(tmp) {
        id         := tmp.Id()
        s1, exists := sa.seats[id]
        if exists {
          vs   = append(vs, s1)     // found a seat
          halt = true
        } else {
          iter = iter + 1           // else, open floor so look further
        }
      } else {
        halt = true                 // off the board
      }
    }
  }
  
  return vs
}

func (s Seat) OccupiedVisibleSeats (sa SeatingArea) []Seat {
  vs := s.VisibleSeats(sa)
  os := make([]Seat, 0)
  for _, s := range vs {
    if s.occupied {
      os = append(os, s)
    }
  }
  return os
}

func (s Seat) OccupiedVisibleSeatCount (sa SeatingArea) int {
  return len(s.OccupiedVisibleSeats(sa))
}

// ---------- UTILITIES -----------------------------------

func (s Seat) Copy () Seat {
  return Seat{x: s.x, y: s.y, occupied: s.occupied}
}

func (s Seat) Id () string {
  return fmt.Sprintf("%d,%d", s.x, s.y)
}
