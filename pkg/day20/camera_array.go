package day20

import (
  "math"

  "github.com/elliotchance/pie/pie"
)


// ========== SUB-DEFINITIONS =============================

type DirectionMatchMap struct {
  e   string
  n   string
  s   string
  w   string
}

type TileMatchMap map[string]DirectionMatchMap


// ========== DEFINITION ==================================

type CameraArray struct {
  tiles   map[string]Tile
}


// ========== RECEIVERS ===================================

func (ca CameraArray) LinkTiles () LinkedCameraArray {
  base := float64(len(ca.tiles)/8)
  sqrt := int(math.Sqrt(base))
  tmm  := ca.CalculateBorderMatches()

  anchors   := make([]Tile, sqrt)
  anchors[0] = ca.OriginTiles()[0]
  for i := 1; i < sqrt; i++ {
    dmm, _    := tmm[anchors[i-1].StateId()]
    aid       := dmm.s
    anchor, _ := ca.tiles[aid]

    anchors[i] = anchor
  }

  lca  := LinkedCameraArray{}
  lca.Initialise(sqrt)
  for y := 0; y < sqrt; y++ {
    tile := anchors[y]
    for x := 0; x < sqrt; x++ {
      lca.RegisterTile(x, y, tile)

      dmm, _ := tmm[tile.StateId()]
      aid    := dmm.e
      tile, _ = ca.tiles[aid]
    }
  }

  return lca
}

func (ca CameraArray) CornerTileIds () pie.Ints {
  tiles := ca.OriginTiles()

  ids := pie.Ints{}
  for _, t := range tiles {
    ids = append(ids, t.id)
  }

  return ids.Sort().Unique()
}


// ---------- ANALYSIS HELPERS ----------------------------

func (ca CameraArray) CalculateBorderMatches () TileMatchMap {
  tmap := TileMatchMap{}
  for id0, t0 := range ca.tiles {
    dmm := DirectionMatchMap{}

    for id1, t1 := range ca.tiles {
      if t0.id != t1.id {
        if t0.borders["e"] == t1.borders["w"] {
          dmm.e = id1
        }
        if t0.borders["n"] == t1.borders["s"] {
          dmm.n = id1
        }
        if t0.borders["s"] == t1.borders["n"] {
          dmm.s = id1
        }
        if t0.borders["w"] == t1.borders["e"] {
          dmm.w = id1
        }
      }
    }
    tmap[id0] = dmm
  }
  return tmap
}

func (ca CameraArray) OriginTiles () []Tile {
  tmm := ca.CalculateBorderMatches()

  tiles := make([]Tile, 0)
  for k, dmm := range tmm {
    if dmm.n == "" && dmm.w == "" {
      tiles = append(tiles, ca.tiles[k])
    }
  }

  return tiles
}


// ---------- BUILD HELPERS -------------------------------

func (ca *CameraArray) Initialise () {
  ca.tiles = make(map[string]Tile)
}

func (ca *CameraArray) RegisterTile (t Tile) {
  ca.tiles[t.StateId()] = t
}

func (ca *CameraArray) RegisterTilePossibilities (tb0 Tile) {
  tb0.SetBorders()
  tb1 := tb0.Rotate("B1")
  tb2 := tb1.Rotate("B2")
  tb3 := tb2.Rotate("B3")
  tf0 := tb0.Flip("F0")
  tf1 := tb1.Flip("F1")
  tf2 := tb2.Flip("F2")
  tf3 := tb3.Flip("F3")

  ca.RegisterTile(tb0)
  ca.RegisterTile(tb1)
  ca.RegisterTile(tb2)
  ca.RegisterTile(tb3)
  ca.RegisterTile(tf0)
  ca.RegisterTile(tf1)
  ca.RegisterTile(tf2)
  ca.RegisterTile(tf3)
}
