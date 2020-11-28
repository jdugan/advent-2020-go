package wire

import (
  "fmt"
  "strconv"
)


// ========== PUBLIC FNS ==================================

func BuildPath (instructions []string) map[string]int {
  steps := 0
  pos   := map[string]int{"x": 0, "y": 0}
  path  := make(map[string]int)

  for _, instr := range instructions {
    moves   := 0
    dir     := instr[0:1]
    dist, _ := strconv.Atoi(instr[1:len(instr)])
    for moves < dist {
      moves  = moves + 1
      steps  = steps + 1
      pos    = move(pos, dir)
      coord := fmt.Sprintf("%d,%d", pos["x"], pos["y"])
      path[coord] = steps
    }
  }

  return path
}

func FindIntersections (w1 map[string]int, w2 map[string]int) []string {
  coords := make([]string, 0)
  for k1, _ := range w1 {
    if w2[k1] != 0 {
      coords = append(coords, k1)
    }
  }
  return coords
}


// ========== PRIVATE FNS =================================

func move (pos map[string]int, dir string) map[string]int {
  x := pos["x"]
  y := pos["y"]

  switch dir {
  case "D":
    y = y - 1
  case "L":
    x = x - 1
  case "R":
    x = x + 1
  case "U":
    y = y + 1
  }

  return map[string]int{"x": x, "y": y}
}
