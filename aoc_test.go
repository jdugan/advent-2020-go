package aoc

import (
  "testing"

  . "github.com/franela/goblin"

  "./pkg/day01"
  "./pkg/day02"
  "./pkg/day03"
  "./pkg/day04"
  "./pkg/day05"
  "./pkg/day06"
  "./pkg/day07"
  "./pkg/day08"
  "./pkg/day09"
  "./pkg/day10"
  "./pkg/day11"
  "./pkg/day12"
  "./pkg/day13"
  "./pkg/day14"
  "./pkg/day15"
  "./pkg/day16"
  "./pkg/day17"
  "./pkg/day18"
  "./pkg/day19"
  "./pkg/day20"
  "./pkg/day21"
  "./pkg/day22"
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
    g.It("Should get correct answers for Day 04", func() {
      g.Assert(day04.Puzzle1()).Equal(202)
      g.Assert(day04.Puzzle2()).Equal(137)
    })
    g.It("Should get correct answers for Day 05", func() {
      g.Assert(day05.Puzzle1()).Equal(864)
      g.Assert(day05.Puzzle2()).Equal(739)
    })
    g.It("Should get correct answers for Day 06", func() {
      g.Assert(day06.Puzzle1()).Equal(6587)
      g.Assert(day06.Puzzle2()).Equal(3235)
    })
    g.It("Should get correct answers for Day 07", func() {
      g.Assert(day07.Puzzle1()).Equal(257)
      g.Assert(day07.Puzzle2()).Equal(1038)
    })
    g.It("Should get correct answers for Day 08", func() {
      g.Assert(day08.Puzzle1()).Equal(1446)
      g.Assert(day08.Puzzle2()).Equal(1403)
    })
    g.It("Should get correct answers for Day 09", func() {
      g.Assert(day09.Puzzle1()).Equal(10884537)
      g.Assert(day09.Puzzle2()).Equal(1261309)
    })
    g.It("Should get correct answers for Day 10", func() {
      g.Assert(day10.Puzzle1()).Equal(2380)
      g.Assert(day10.Puzzle2()).Equal(48358655787008)
    })
    g.It("Should get correct answers for Day 11", func() {
      g.Assert(day11.Puzzle1()).Equal(2324)
      // g.Assert(day11.Puzzle2()).Equal(2068)
    })
    g.It("Should get correct answers for Day 12", func() {
      g.Assert(day12.Puzzle1()).Equal(796)
      g.Assert(day12.Puzzle2()).Equal(39446)
    })
    g.It("Should get correct answers for Day 13", func() {
      g.Assert(day13.Puzzle1()).Equal(2935)
      g.Assert(day13.Puzzle2()).Equal(836024966345345)
    })
    g.It("Should get correct answers for Day 14", func() {
      g.Assert(day14.Puzzle1()).Equal(17934269678453)
      g.Assert(day14.Puzzle2()).Equal(3440662844064)
    })
    g.It("Should get correct answers for Day 15", func() {
      g.Assert(day15.Puzzle1()).Equal(475)
      // g.Assert(day15.Puzzle2()).Equal(11261)
    })
    g.It("Should get correct answers for Day 16", func() {
      g.Assert(day16.Puzzle1()).Equal(25972)
      g.Assert(day16.Puzzle2()).Equal(622670335901)
    })
    g.It("Should get correct answers for Day 17", func() {
      g.Assert(day17.Puzzle1()).Equal(391)
      // g.Assert(day17.Puzzle2()).Equal(2264)
    })
    g.It("Should get correct answers for Day 18", func() {
      g.Assert(day18.Puzzle1()).Equal(6811433855019)
      g.Assert(day18.Puzzle2()).Equal(129770152447927)
    })
    g.It("Should get correct answers for Day 19", func() {
      g.Assert(day19.Puzzle1()).Equal(126)
      // g.Assert(day19.Puzzle2()).Equal(282)
    })
    g.It("Should get correct answers for Day 20", func() {
      g.Assert(day20.Puzzle1()).Equal(32287787075651)
      g.Assert(day20.Puzzle2()).Equal(1939)
    })
    g.It("Should get correct answers for Day 21", func() {
      g.Assert(day21.Puzzle1()).Equal(2374)
      g.Assert(day21.Puzzle2()).Equal("fbtqkzc,jbbsjh,cpttmnv,ccrbr,tdmqcl,vnjxjg,nlph,mzqjxq")
    })
    g.It("Should get correct answers for Day 22", func() {
      g.Assert(day22.Puzzle1()).Equal(-1)
      g.Assert(day22.Puzzle2()).Equal(-2)
    })
  })
}
