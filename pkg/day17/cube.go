package day17


// ========== DEFINITION ==================================

type Cube struct {
  coord         Coord
  state         bool
}


// ========== RECEIVERS ===================================

func (c Cube) Id () string {
  return c.coord.Id()
}

func (c Cube) NewState (count int) bool {
  state := c.state

  if !!state {
    if count < 2 || count > 3 {
      state = false
    }
  } else {
    if count == 3 {
      state = true
    }
  }

  return state
}
