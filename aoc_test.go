package aoc

import (
  "testing"

  . "github.com/franela/goblin"

  "./pkg/day01"
)

func Test(t *testing.T) {
  g := Goblin(t)

  g.Describe("AOC", func() {
    g.It("Should get correct answers for Day 01", func() {
      g.Assert(day01.Puzzle1()).Equal(-1)
      g.Assert(day01.Puzzle2()).Equal(-2)
    })
  })
}
