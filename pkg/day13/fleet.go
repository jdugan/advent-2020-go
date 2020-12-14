package day13

import (
  "math"

  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Fleet []Bus


// ========== RECEIVERS ===================================

func (f Fleet) CheckSum (ts int) int {
  id   := 0
  wait := 61

  for _, b := range f {
    tmp := b.WaitTime(ts)
    if tmp < wait {
      id   = b.id
      wait = tmp
    }
  }

  return id * wait
}

func (f Fleet) PerfectTimestamp () int {
  base  := f[0].id
  fleet := f.Consolidate()

  iter  := int(math.Floor(float64(8*100000000000000)/float64(base)))
  halt  := false

  for !halt {
    remainders := pie.Ints{}
    etd        := base * iter

    for _, bus := range fleet {
      wt := bus.WaitTime(etd)
      if !bus.Match(wt) {
        remainders = append(remainders, wt - bus.offset)
      }
    }

    if remainders.Len() == 0 {
      halt = true
    } else {
      maxrem := remainders.Sort().Last()
      skips  := 1
      if maxrem >= base {
        skips = int(math.Floor(float64(maxrem)/float64(base)))
      }

      iter = iter + skips
    }
  }

  return base * iter
}


// ---------- UTILITIES -----------------------------------

func (f Fleet) Consolidate () Fleet {
  bmap := make(map[int]Bus)
  keys := pie.Ints{}

  // convert fleet to map, consolidating
  // where buses have the same offset
  for _, bus := range f {
    k := bus.offset
    v := bus.id

    ebus, exists := bmap[bus.offset]
    if exists {
      bmap[k] = Bus{id: v * ebus.id, offset: k}
    } else {
      bmap[k] = bus
    }

    keys = append(keys, k)
  }

  // consolidate offsets forward when possible to
  // produce the smallest number of factors that
  // are as large themsevles as possible
  for _, k0 := range keys {
    bus0, _      := bmap[k0]
    k1           := bus0.id + bus0.offset
    bus1, exists := bmap[k1]

    if exists {
      bmap[k0] = Bus{id: 0, offset: 0}
      bmap[k1] = Bus{id: bus0.id * bus1.id, offset: k1}
    }
  }

  // extract significant buses
  fleet := Fleet{}
  for _, bus := range bmap {
    if bus.id > 0 {
      fleet = append(fleet, bus)
    }
  }

  return fleet
}
