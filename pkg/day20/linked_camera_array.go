package day20

import (
  "math"
)


// ========== DEFINITION ==================================

type LinkedCameraArray struct {
  grid   [][]Tile
}


// ========== RECEIVERS ===================================

func (lca LinkedCameraArray) ConsolidateTiles () Tile {
  tile := Tile{}
  tile.Initialise(1001)

  for y, row := range lca.grid {
    for x, t0 := range row {
      t1 := t0.Borderless()
      for _, c := range t1.coords {
        x1 := (x * t1.width) + c.x
        y1 := (y * t1.height) + c.y
        c1 := Coord{x: x1, y: y1, display: c.display}

        tile.RegisterCoord(c1)
      }
    }
  }

  base := float64(len(tile.coords))
  size := int(math.Sqrt(base))

  tile.width  = size
  tile.height = size
  tile.SetBorders()
  
  return tile
}


// ---------- BUILD HELPERS -------------------------------

func (lca *LinkedCameraArray) Initialise (size int) {
  g := make([][]Tile, size)
  for i, _ := range g {
    g[i] = make([]Tile, size)
  }
  lca.grid = g
}

func (lca *LinkedCameraArray) RegisterTile (x int, y int, t Tile) {
  lca.grid[y][x] = t
}
