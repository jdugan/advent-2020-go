package day99


// ========== DEFINITION ==================================

type Intersection struct {
  c0  Coord
  c1  Coord
}


// ========== RECEIVERS ===================================

func (is Intersection) ManhattanDistance () int {
  return is.c0.Manhattan()
}

func (is Intersection) StepDistance () int {
  return is.c0.steps + is.c1.steps
}
