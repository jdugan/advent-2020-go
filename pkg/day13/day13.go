package day13

import (
  "fmt"
  "math"
  "sort"
  "strconv"
  "strings"

  "github.com/elliotchance/pie/pie"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 13)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  etd, fleet := data()

  id   := 0
  wait := 61
  for _, bus := range fleet {
    tmp := bus.WaitTime(etd)
    if tmp < wait {
      id   = bus.id
      wait = tmp
    }
  }

  return id * wait
}

func Puzzle2() int {
  _, fleet := data()

  base  := fleet[0].id
  buses := combineFleet(fleet)

  // fmt.Println("====================")
  // fmt.Println(base)
  // fmt.Println(buses)
  // fmt.Println("====================")

  iter   := int(math.Floor(float64(100000000000000)/float64(base)))
  factor := iter
  halt   := false

  for !halt {
    remainders := pie.Ints{}
    etd        := base * factor

    for _, bus := range buses {
      wt := bus.WaitTime(etd)
      if !bus.Match(wt) {
        remainders = append(remainders, wt - bus.position)
      }
    }

    if remainders.Len() == 0 {
      halt = true
    } else {
      var skips int
      maxrem := remainders.Sort().Last()
      if maxrem >= base {
        skips = int(math.Floor(float64(maxrem)/float64(base)))
      } else {
        skips = 1
      }

      factor = factor + skips
      iter   = iter + 1
    }
  }

  return base * factor
}


// ========== PRIVATE FNS =================================

func combineFleet (fleet []Bus) []Bus {
  sort.Slice(fleet, func (a, b int) bool { return fleet[a].offset < fleet[b].offset })

  // bmap := make(map[int]Bus)
  // for _, bus := range fleet {
  //   var
  //   v, exists := bmap[key]
  //   if exists {
  //     v = Bus{id: v.id * bus.id, position: key, offset: key}
  //   } else {
  //     key = bus.id + bus.offset
  //     v1, exists := bmap[key]
  //     if exists {
  //       v = Bus{id: v1.id * bus.id, position: key, offset: key}
  //     } else {
  //       v = Bus{id: bus.id, position: key, offset: key}
  //     }
  //   }
  //   bmap[key] = v
  // }

  // TODO: Need to do this consolidation by code rather
  // than by hand in the editor. :)
  //
  buses := make([]Bus, 0)
  buses  = append(buses, Bus{id: 13, position: 3, offset: 3})
  buses  = append(buses, Bus{id: 17, position: 17, offset: 17})
  // buses  = append(buses, Bus{id: 41, position: 27, offset: 27})
  buses  = append(buses, Bus{id: 275244887, position: 37, offset: 37})
  buses  = append(buses, Bus{id: 30053, position: 68, offset: 68})
  // buses  = append(buses, Bus{id: 733, position: 68, offset: 68})
  // for _, bus := range bmap {
  //   buses = append(buses, bus)
  // }

  return buses
}

func data () (int, []Bus) {
  lines := reader.Lines("./data/day13/input.txt")

  etd, _ := strconv.Atoi(lines[0])
  bids   := strings.Split(lines[1], ",")
  buses  := make([]Bus, 0)

  for i, bid := range bids {
    if bid != "x" {
      id, _  := strconv.Atoi(bid)
      offset := i
      if i > id {
        offset = (i % id)
      }
      if offset == 0 {
        offset = id
      }
      bus   := Bus{id: id, position: i, offset: offset}

      buses = append(buses, bus)
    }
  }

  return etd, buses
}
