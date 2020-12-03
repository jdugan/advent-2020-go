package day03


// ========== DEFINITION ==================================

type Slope struct {
  trees   map[string]Coord
  length  int
  period  int
}


// ========== RECEIVERS ===================================

func (s Slope) Count (dx int, dy int) int {
  coord := Coord{x: 0, y: 0}
  count := 0

  for coord.y < s.length {
    base      := coord.Base(s.period)
    _, isTree := s.trees[base.Key()]
    if isTree {
      count = count + 1
    }
    coord = coord.Move(dx, dy)
  }

  return count
}
