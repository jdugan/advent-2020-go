package day99

import (
  "strconv"

  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Wire struct {
  path map[string]Coord
}


// ========== RECEIVERS ===================================

func (w Wire) Build (cmds pie.Strings) Wire {
  path  := make(map[string]Coord)
  coord := Coord{x: 0, y: 0, steps: 0}

  for _, cmd := range cmds {
    runes := []rune(cmd)
    dir     := string(runes[0])
    dist, _ := strconv.Atoi(string(runes[1:]))

    for i := 0; i < dist; i++ {
      coord  = coord.Move(dir)
      key   := coord.Key()

      _, found := path[key]
      if !found {
        path[key] = coord
      }
    }
  }

  return Wire{path: path}
}

func (w Wire) Keys () pie.Strings {
  keys := make([]string, 0)
  for k, _ := range w.path {
    keys = append(keys, k)
  }
  return pie.Strings(keys).Sort()
}
