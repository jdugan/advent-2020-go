package day17

import (
  // "fmt"

  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type PowerSource struct {
  cubes  map[string]Cube
}


// ========== RECEIVERS ===================================

func (ps PowerSource) Run (cycles int, dims int) int {
  for i := 0; i < cycles; i++ {
    ps = ps.Cycle(dims)
  }

  return ps.CubeCount()
}


// ---------- PROCESSES -----------------------------------

func (ps PowerSource) Cycle (dims int) PowerSource {
  wps := PowerSource{cubes: make(map[string]Cube)}

  // find all ids that might become active
  // during this cycle
  ids := ps.PossibleCoordIds(dims)

  // for each possible id, determine its new state,
  // and add to new map if new state is on
  for _, id := range ids {
    coord        := Coord{}.FromId(id)
    cube, exists := ps.cubes[id]
    if !exists {
      cube = Cube{coord: coord, state: false}
    }

    count     := ps.ActiveAdjacentCubeCount(coord, dims)
    cube.state = cube.NewState(count)

    if !!cube.state {
      wps.cubes[cube.Id()] = cube
    }
  }

  return wps
}


// ---------- UTILITIES -----------------------------------

func (ps PowerSource) ActiveAdjacentCubeCount (coord Coord, dims int) int {
  count := 0
  for _, aid := range coord.AdjacentIds(dims) {
    _, exists := ps.cubes[aid]
    if exists {
      count = count + 1
    }
  }
  return count
}

func (ps PowerSource) CubeCount () int {
  return len(ps.cubes)
}

func (ps PowerSource) PossibleCoordIds (dims int) pie.Strings {
  ids := pie.Strings{}

  for id, cube := range ps.cubes {
    ids = append(ids, id)
    ids = append(ids, cube.coord.AdjacentIds(dims)...)
  }
  ids = ids.Sort().Unique()

  return ids
}
