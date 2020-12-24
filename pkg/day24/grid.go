package day24

import (
  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION(S) ===============================

type Grid map[string]Tile


// ========== RECEIVERS ===================================

func (g Grid) Count () int {
  return len(g)
}

func (g Grid) Install (instrs []pie.Strings) Grid {
  for _, path := range instrs {
    dest := g.MapPath(path)
    did  := dest.Id()

    tile, exists := g[did]
    if exists {
      g[did] = tile.Flip()
    } else {
      g[did] = dest
    }
  }

  return g.Purge()
}

func (g Grid) Purge () Grid {
  g1 := Grid{}
  for id, tile := range g {
    if tile.IsBlack() {
      g1[id] = tile
    }
  }
  return g1
}

func (g Grid) Rotate (days int) Grid {
  for i := 0; i < days; i++ {
    g = g.Shuffle()
  }
  return g
}

func (g Grid) Shuffle () Grid {
  g1 := Grid{}

  sids := g.SignificantIds()
  for _, id := range sids {
    t, exists := g[id]
    if !exists {
      t = Tile{}.FromId(id)
    }

    count := 0
    nids  := t.GetNeighborIds()
    for _, nid := range nids {
      nt, exists := g[nid]
      if exists && nt.IsBlack() {
        count= count + 1
        if count > 2 {
          break
        }
      }
    }

    if t.IsBlack() && (count == 0 || count > 2) {
      t.color = 0
    } else {
      if !t.IsBlack() && count == 2 {
        t.color = 1
      }
    }

    if t.IsBlack() {
      g1[t.Id()] = t
    }
  }

  return g1
}


// ---------- HELPERS -------------------------------------

func (g Grid) MapPath (path pie.Strings) Tile {
  dest := g.Origin()

  for _, dir := range path {
    dest = dest.BuildNeighbor(dir).Flip()
  }

  return dest
}

func (g Grid) Origin () Tile {
  return Tile{}
}

func (g Grid) SignificantIds () pie.Strings {
  sids := pie.Strings{}

  for _, t := range g {
    sids = append(sids, t.Id())
    sids = append(sids, t.GetNeighborIds()...)
  }

  return sids.Sort().Unique()
}
