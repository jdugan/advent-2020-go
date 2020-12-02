package day99

import (
  "fmt"
  "strings"

  "github.com/elliotchance/pie/pie"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 3)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  intersections := Intersections()
  shortest      := intersections[0].ManhattanDistance()
  for _, intersection := range intersections[1:] {
    d := intersection.ManhattanDistance()
    if d < shortest {
      shortest = d
    }
  }
  return shortest
}

func Puzzle2() int {
  intersections := Intersections()
  shortest      := intersections[0].StepDistance()
  for _, intersection := range intersections[1:] {
    d := intersection.StepDistance()
    if d < shortest {
      shortest = d
    }
  }
  return shortest
}


// ========== PRIVATE FNS =================================

func Data () pie.Strings {
  lines := reader.Lines("./data/day03/input.txt")
  return lines
}

func Intersections () []Intersection {
  wires := Wires()
  w0    := wires[0]
  w1    := wires[1]
  p0    := w0.path
  p1    := w1.path
  keys  := w0.Keys().Intersect(w1.Keys())

  intersections := make([]Intersection, len(keys))
  for i, k := range keys {
    c0 := p0[k]
    c1 := p1[k]
    intersections[i] = Intersection{c0: c0, c1: c1}
  }

  return intersections
}

func Wires () []Wire {
  lines := Data()
  wires  := make([]Wire, len(lines))
  for index, line := range lines {
    cmds := pie.Strings(strings.Split(line, ","))
    wires[index] = Wire{}.Build(cmds)
  }
  return wires
}
