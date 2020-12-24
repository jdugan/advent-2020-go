package day24

import (
  "fmt"
  "strings"

  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Tile struct {
  x       int
  y       int
  color   int       // 0 = white; 1 = black
}


// ========== RECEIVERS ===================================

func (t Tile) Flip () Tile {
  t.color = (t.color + 1) % 2

  return t
}

func (t Tile) FromId (id string) Tile {
  coords := pie.Strings(strings.Split(id, ",")).Ints()

  return Tile{x: coords[0], y: coords[1]}
}

func (t Tile) Id () string {
  return fmt.Sprintf("%d,%d", t.x, t.y)
}

func (t Tile) IsBlack () bool {
  return t.color == 1
}


// ---------- NEIGHBORS -----------------------------------

func (t Tile) BuildNeighbor (dir string) Tile {
  x, y := t.GetNeighborCoords(dir)

  return Tile{x: x, y: y}
}

func (t Tile) GetNeighborCoords (dir string) (int, int) {
  x := t.x
  y := t.y

  switch dir {
  case "e":
    x = x + 2
  case "w":
    x = x - 2
  case "ne":
    x = x + 1
    y = y + 1
  case "nw":
    x = x - 1
    y = y + 1
  case "se":
    x = x + 1
    y = y - 1
  case "sw":
    x = x - 1
    y = y - 1
  }

  return x, y
}

func (t Tile) GetNeighborIds () pie.Strings {
  nids := pie.Strings{}

  nids  = append(nids, t.BuildNeighbor("e").Id())
  nids  = append(nids, t.BuildNeighbor("w").Id())
  nids  = append(nids, t.BuildNeighbor("ne").Id())
  nids  = append(nids, t.BuildNeighbor("nw").Id())
  nids  = append(nids, t.BuildNeighbor("se").Id())
  nids  = append(nids, t.BuildNeighbor("sw").Id())

  return nids
}
