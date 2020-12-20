package day20

import (
  "fmt"

  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Tile struct {
  id            int
  width         int
  height        int
  orientation   string

  borders       map[string]string
  coords        map[string]Coord
}


// ========== RECEIVERS ===================================

func (t *Tile) SetBorders () {
  bmap     := make(map[string]string)
  bmap["n"] = t.BorderNorth()
  bmap["e"] = t.BorderEast()
  bmap["s"] = t.BorderSouth()
  bmap["w"] = t.BorderWest()

  t.borders = bmap
}

func (t Tile) StateId () string {
  return fmt.Sprintf("%d%s", t.id, t.orientation)
}


// ---------- ACTIONS (COPY) ------------------------------

func (t Tile) Flip (orientation string) Tile {
  t1 := t.ShallowCopy(orientation)

  for y := 0; y < t.height; y++ {
    for x := 0; x < t.width; x++ {
      x1 := t1.width - x - 1
      y1 := y

      c         := Coord{x: x, y: y}
      c1        := Coord{x: x1, y: y1}
      c1.display = t.coords[c.Id()].display

      t1.RegisterCoord(c1)
    }
  }
  t1.SetBorders()

  return t1
}

func (t Tile) Rotate (orientation string) Tile {
  t1 := t.ShallowCopy(orientation)

  for x := 0; x < t.width; x++ {
    for y := 0; y < t.height; y++ {
      x1 := t1.height - y - 1
      y1 := x

      c          := Coord{x: x, y: y}
      c1         := Coord{x: x1, y: y1}
      c1.display  = t.coords[c.Id()].display

      t1.RegisterCoord(c1)
    }
  }
  t1.SetBorders()

  return t1
}


// ---------- ACTIONS (OTHER) -----------------------------

func (t *Tile) AddRow (displays pie.Strings) {
  if displays.Len() > t.width {
    t.width = displays.Len()
  }
  t.height = t.height + 1

  for x, display := range displays {
    coord := Coord{x: x, y: t.height, display: display}
    t.coords[coord.Id()] = coord
  }
}

func (t *Tile) RegisterCoord (c Coord) {
  t.coords[c.Id()] = c
}


// ---------- BORDERS -------------------------------------

func (t Tile) BorderEast () string {
  a := pie.Strings{}

  x := t.width - 1
  for y := 0; y < t.height; y++ {
    id   := Coord{}.GenerateId(x, y)
    c, _ := t.coords[id]
    a = append(a, c.display)
  }

  return a.Join("")
}

func (t Tile) BorderNorth () string {
  a := pie.Strings{}

  y := 0
  for x := 0; x < t.width; x++ {
    id   := Coord{}.GenerateId(x, y)
    c, _ := t.coords[id]
    a = append(a, c.display)
  }

  return a.Join("")
}

func (t Tile) BorderSouth () string {
  a := pie.Strings{}

  y := t.height - 1
  for x := 0; x < t.width; x++ {
    id   := Coord{}.GenerateId(x, y)
    c, _ := t.coords[id]
    a = append(a, c.display)
  }

  return a.Join("")
}

func (t Tile) BorderWest () string {
  a := pie.Strings{}

  x := 0
  for y := 0; y < t.height; y++ {
    id   := Coord{}.GenerateId(x, y)
    c, _ := t.coords[id]
    a = append(a, c.display)
  }

  return a.Join("")
}

// ---------- PRINTING ------------------------------------

func (t Tile) PrintBorders () {
  bm := t.borders
  fmt.Println("  n:", bm["n"])
  fmt.Println("  e:", bm["e"])
  fmt.Println("  s:", bm["s"])
  fmt.Println("  w:", bm["w"])
}

func (t Tile) PrintGrid () {
  for y := 0; y < t.height; y++ {
    row := pie.Strings{}
    for x := 0; x < t.width; x++ {
      id   := Coord{x: x, y: y}.Id()
      c, _ := t.coords[id]
      row   = append(row, c.display)
    }
    fmt.Println(row.Join(""))
  }
}

func (t Tile) PrintReport () {
  t.PrintTitle()
  t.PrintGrid()
  t.PrintBorders()
  fmt.Println("")
}

func (t Tile) PrintTitle () {
  title := fmt.Sprintf("%d: %s", t.id, t.orientation)
  fmt.Println(title)
}


// ---------- UTILITIES -----------------------------------

func (t Tile) Borderless () Tile {
  t1 := t.ShallowCopy("B0")
  for _, c := range t.coords {
    if c.x != 0 && c.x != t.width - 1 && c.y != 0 && c.y != t.height - 1 {
      c1 := Coord{x: c.x - 1, y: c.y - 1, display: c.display}
      t1.RegisterCoord(c1)
    }
  }
  t1.width  = t1.width - 2
  t1.height = t1.height - 2
  return t1
}

func (t *Tile) Initialise (id int) {
  t.id          = id
  t.width       = -1
  t.height      = -1
  t.orientation = "B0"
  t.borders     = make(map[string]string)
  t.coords      = make(map[string]Coord)
}

func (t Tile) ShallowCopy (orientation string) Tile {
  t1            := Tile{}
  t1.id          = t.id
  t1.width       = t.width
  t1.height      = t.height
  t1.orientation = orientation
  t1.borders     = make(map[string]string)
  t1.coords      = make(map[string]Coord)

  return t1
}
