package day11

import (
  "fmt"

  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type SeatingArea struct {
  seats   map[string]Seat
  width   int
  height  int
}


// ========== RECEIVERS ===================================

func (sa SeatingArea) IterateByAdjacency () SeatingArea {
  wsa := sa.Copy()
  for id, s := range sa.seats {
    s1 := sa.ReassignByAdjacentSeats(s)
    wsa.seats[id] = s1
  }
  return wsa
}

func (sa SeatingArea) IterateByVisibility () SeatingArea {
  wsa := sa.Copy()
  for id, s := range sa.seats {
    s1 := sa.ReassignByVisibleSeats(s)
    wsa.seats[id] = s1
  }
  return wsa
}


// ---------- PRE-PROCESSORS ------------------------------

func (sa SeatingArea) SetAdjacentSeatIds () SeatingArea {
  for id, s := range sa.seats {
    s.adjacentIds = sa.FindAdjacentSeatIds(s)
    sa.seats[id]  = s
  }
  return sa
}

func (sa SeatingArea) SetVisibleIds () SeatingArea {
  for id, s := range sa.seats {
    s.visibleIds = sa.FindVisibleIds(s)
    sa.seats[id] = s
  }
  return sa
}


// ---------- ADJACENT HELPERS ----------------------------

func (sa SeatingArea) FindAdjacentSeatIds (s Seat) pie.Strings {
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

func (sa SeatingArea) ReassignByAdjacentSeats (s Seat) Seat {
  s1   := s.Copy()
  oasc := sa.OccupiedAdjacentSeatCount(s)
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

func (sa SeatingArea) OccupiedAdjacentSeatCount (s Seat) int {
  count := 0
  for _, id := range s.adjacentIds {
    s, exists := sa.seats[id]
    if exists && s.occupied {
      count = count + 1
      if count > 3 {
        break
      }
    }
  }
  return count
}


// ---------- VISIBLE HELPERS -----------------------------

func (sa SeatingArea) FindVisibleIds (s Seat) pie.Strings {
  slopes := make([]pie.Ints, 0)
  slopes  = append(slopes, pie.Ints{-1, -1})      // northwest
  slopes  = append(slopes, pie.Ints{0, -1})       // north
  slopes  = append(slopes, pie.Ints{1, -1})       // northeast
  slopes  = append(slopes, pie.Ints{1, 0})        // east
  slopes  = append(slopes, pie.Ints{1, 1})        // southeast
  slopes  = append(slopes, pie.Ints{0, 1})        // south
  slopes  = append(slopes, pie.Ints{-1, 1})       // southwest
  slopes  = append(slopes, pie.Ints{-1, 0})       // west

  vsids := pie.Strings{}
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
          vsids = append(vsids, id)
          halt  = true
        } else {                            // open floor; continue
          iter = iter + 1
        }
      } else {                              // off the board; stop
        halt = true
      }
    }
  }

  return vsids
}

func (sa SeatingArea) OccupiedVisibleSeatCount (s Seat, max int) int {
  count := 0
  for _, id := range s.visibleIds {
    s, exists := sa.seats[id]
    if exists && s.occupied {
      count = count + 1
      if count > max {
        break
      }
    }
  }
  return count
}

func (sa SeatingArea) ReassignByVisibleSeats (s Seat) Seat {
  s1 := s.Copy()
  if s.occupied {
    oasc := sa.OccupiedVisibleSeatCount(s, 4)
    if oasc > 4 {
      s1.occupied = false
    }
  } else {
    oasc := sa.OccupiedVisibleSeatCount(s, 0)
    if oasc == 0 {
      s1.occupied = true
    }
  }
  return s1
}


// ---------- STATE HELPERS -------------------------------

func (sa SeatingArea) IsSeatIdWithinArea (s Seat) bool {
  return s.x > 0 && s.x <= sa.width &&
         s.y > 0 && s.y <= sa.height
}

func (sa SeatingArea) TooManyOccupiedSeats (seatIds pie.Strings, maxOccupied int) bool {
  result := false
  count  := 0
  for _, id := range seatIds {
    s, exists := sa.seats[id]
    if exists && s.occupied {
      count = count + 1
      if count > maxOccupied {
        result = true
        break
      }
    }
  }
  return result
}

func (sa SeatingArea) TotalOccupiedSeatCount () int {
  count := 0
  for _, v := range sa.seats {
    if v.occupied {
      count = count + 1
    }
  }
  return count
}


// ---------- UTILITIES -----------------------------------

func (sa SeatingArea) Copy () SeatingArea {
  sa1 := SeatingArea{seats: make(map[string]Seat), width: sa.width, height: sa.height}
  for k, v := range sa.seats {
    sa1.seats[k] = v.Copy()
  }
  return sa1
}

func (sa SeatingArea) Print (step int) int {
  fmt.Println("")
  for y := 1; y <= sa.height; y++ {
    row := make(pie.Strings, 0)
    for x := 1; x <= sa.width; x++ {
      tmp       := Seat{x: x, y: y}
      s, exists := sa.seats[tmp.Id()]
      if exists {
        if s.occupied {
          row = append(row, "#")
        } else {
          row = append(row, "L")
        }
      } else {
        row = append(row, ".")
      }
    }
    fmt.Println(row.Join(""))
  }
  fmt.Println("")
  fmt.Println("Occupied: ", sa.TotalOccupiedSeatCount())
  fmt.Println("Step:     ", step)
  fmt.Println("")

  return sa.TotalOccupiedSeatCount()
}
