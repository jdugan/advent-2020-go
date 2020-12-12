package day12


// ========== DEFINITION ==================================

type Command struct {
  dir     string
  dist    int
}


// ========== RECEIVERS ===================================

func (c Command) CalculateHeading (heading string) string {
  h := heading
  switch heading {
  case "L":
    h = TurnLeft(heading)
  case "R":
    h = TurnRight(heading)
  }
  return h
}
