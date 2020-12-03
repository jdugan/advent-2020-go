package aoc

import (
  "testing"

  . "github.com/franela/goblin"

  "./pkg/day01"
  "./pkg/day02"
  "./pkg/day03"
)

func Test(t *testing.T) {
  g := Goblin(t)

  g.Describe("AOC", func() {
    g.It("Should get correct answers for Day 01", func() {
      g.Assert(day01.Puzzle1()).Equal(842016)
      g.Assert(day01.Puzzle2()).Equal(9199664)
    })
    g.It("Should get correct answers for Day 02", func() {
      g.Assert(day02.Puzzle1()).Equal(454)
      g.Assert(day02.Puzzle2()).Equal(649)
    })
    g.It("Should get correct answers for Day 03", func() {
      g.Assert(day03.Puzzle1()).Equal(200)
      g.Assert(day03.Puzzle2()).Equal(3737923200)
    })
  })
}
