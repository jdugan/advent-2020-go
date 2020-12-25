package day25

import (
  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Device struct {
  loops     int
  pubkey    int
}

const DIVISOR = 20201227


// ========== RECEIVERS ===================================

func (d Device) Decrypt (sub int) int {
  powers := d.FixedPowers(sub)
  loops  := 0

  value  := d.pubkey
  index  := powers.FindFirstUsing(func (p int) bool { return value == p })

  for index == -1 {
    loops = loops + 1
    value = value + DIVISOR
    for value % sub != 0 {
      value = value + DIVISOR
    }
    value = value / sub
    index = powers.FindFirstUsing(func (p int) bool { return value == p })
  }

  return loops + index
}

func (d Device) Encrypt (sub int) int {
  val := 1
  for i := 0; i < d.loops; i++ {
    val = (val * sub) % DIVISOR
  }
  return val
}


// ---------- HELPERS -------------------------------------

func (d Device) FixedPowers (sub int) pie.Ints {
  powers := pie.Ints{}

  val := 1
  for val < DIVISOR {
    powers = append(powers, val)
    val    = val * sub
  }

  return powers
}
