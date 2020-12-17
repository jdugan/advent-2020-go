package day17

import (
  "fmt"
  "strings"

  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Coord struct {
  x   int
  y   int
  z   int
  w   int
}


// ========== RECEIVERS ===================================

func (c Coord) AdjacentIds (dims int) pie.Strings {
  ids := pie.Strings{}

  if dims == 3 {
    ids = c.AdjacentIdsInThreeDims()
  } else {
    ids = c.AdjacentIdsInFourDims()
  }

  return ids
}

func (c Coord) FromId (id string) Coord {
  vs := pie.Strings(strings.Split(id, ",")).Ints()

  return Coord{x: vs[0], y: vs[1], z: vs[2], w: vs[3]}
}

func (c Coord) Id () string {
  return fmt.Sprintf("%d,%d,%d,%d", c.x, c.y, c.z, c.w)
}


// ---------- DIMENSIONS ----------------------------------

func (c Coord) AdjacentIdsInThreeDims () pie.Strings {
  ids := pie.Strings{}

  for x:= -1; x < 2; x++ {
    for y:= -1; y < 2; y++ {
      for z:= -1; z < 2; z++ {
        tmp := Coord{x: c.x + x, y: c.y + y, z: c.z + z}
        id  := tmp.Id()

        if id != c.Id() {
          ids = append(ids, id)
        }
      }
    }
  }

  return ids
}

func (c Coord) AdjacentIdsInFourDims () pie.Strings {
  ids := pie.Strings{}

  for x:= -1; x < 2; x++ {
    for y:= -1; y < 2; y++ {
      for z:= -1; z < 2; z++ {
        for w:= -1; w < 2; w++ {
          tmp := Coord{x: c.x + x, y: c.y + y, z: c.z + z, w: c.w + w}
          id  := tmp.Id()

          if id != c.Id() {
            ids = append(ids, id)
          }
        }
      }
    }
  }

  return ids
}
