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
    s1 := s.ReassignByAdjacentSeats(sa)
    wsa.seats[id] = s1
  }
  return wsa
}

func (sa SeatingArea) IterateByVisibility () SeatingArea {
  wsa := sa.Copy()
  for id, s := range sa.seats {
    s1 := s.ReassignByVisibleSeats(sa)
    wsa.seats[id] = s1
  }
  return wsa
}


// ---------- PRE-PROCESSORS ------------------------------

func (sa SeatingArea) SetAdjacentSeatIds () SeatingArea {
  for id, s := range sa.seats {
    s.adjacentIds = s.AdjacentSeatIds(sa)
    sa.seats[id]  = s
  }
  return sa
}

func (sa SeatingArea) SetVisibleSlopes () SeatingArea {
  for id, s := range sa.seats {
    s.visibleSlopes = s.VisibleSlopes(sa)
    sa.seats[id]    = s
  }
  return sa
}


// ---------- STATE HELPERS -------------------------------

func (sa SeatingArea) IsSeatIdWithinArea (s Seat) bool {
  return s.x > 0 && s.x <= sa.width &&
         s.y > 0 && s.y <= sa.height
}

func (sa SeatingArea) OccupiedSeatCount () int {
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
  fmt.Println("Occupied: ", sa.OccupiedSeatCount())
  fmt.Println("Step:     ", step)
  fmt.Println("")

  return sa.OccupiedSeatCount()
}
