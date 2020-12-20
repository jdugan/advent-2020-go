package day20


// ========== DEFINITION ==================================

type Analyser struct {
  tiles   map[string]Tile
}


// ========== RECEIVERS ===================================

func (a Analyser) SeaRoughness () int {
  // make a monster for comparison
  // purposes
  monster := Monster{}
  monster.Initialise()

  // loop tiles until we find at least
  // one monster; replace the ones we
  // do find; then quit so we can count
  tile   := Tile{}
  mcount := 0
  for _, t := range a.tiles {
    xmax := t.width - monster.width - 1

    for y := 0; y < t.height; y++ {
      for x := 0; x < xmax; x++ {
        found := true
        for _, os := range monster.offsets {
          cid  := Coord{x: x + os[0], y: y + os[1]}.Id()
          c, exists := t.coords[cid]
          if !exists || c.display != "#" {
            found = false
            break
          }
        }
        if found {
          mcount = mcount + 1
          for _, os := range monster.offsets {
            cid  := Coord{x: x + os[0], y: y + os[1]}.Id()
            c, _ := t.coords[cid]
            c.display = "O"
            t.coords[cid] = c
          }
        }
      }
    }

    if mcount > 0 {
      tile = t
      break
    }
  }

  // count the rough seas
  rcount := 0
  for _, c := range tile.coords {
    if c.display == "#" {
      rcount = rcount + 1
    }
  }

  return rcount
}
