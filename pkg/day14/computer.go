package day14

import (
  "fmt"
  "strings"
  "strconv"

  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Computer struct {
  mask    pie.Strings
  memory  map[int]int
}


// ========== RECEIVERS ===================================

func (c Computer) Initialise (cmds pie.Strings) int {
  for _, cmd := range cmds {
    p1, p2 := c.SplitCommand(cmd)

    if p1 == "mask" {
      c.mask = c.SplitMask(p2)
    } else {
      k, _ := strconv.Atoi(p1)
      v, _ := strconv.Atoi(p2)
      barr := c.ConvertDecimalToBinaryArray(v)
      mv   := c.TransformBinaryArrayThroughMask(barr, c.mask)

      c.memory[k] = mv
    }
  }

  return c.Sum()
}

func (c Computer) Decode (cmds pie.Strings) int {
  for _, cmd := range cmds {
    p1, p2 := c.SplitCommand(cmd)

    if p1 == "mask" {
      c.mask = c.SplitMask(p2)
    } else {
      k, _  := strconv.Atoi(p1)
      v, _  := strconv.Atoi(p2)
      barr  := c.ConvertDecimalToBinaryArray(k)
      fmask := c.DecodeBinaryArrayThroughMask(barr, c.mask)
      addrs := c.ConvertFloatingMaskToAddresses(fmask)

      for _, addr := range addrs {
        c.memory[addr] = v
      }
    }
  }

  return c.Sum()
}


// ---------- MASK OPERATIONS -----------------------------

func (c Computer) DecodeBinaryArrayThroughMask (barr []string, mask pie.Strings) pie.Strings {
  carr := pie.Strings(make([]string, len(barr)))

  for i, mv := range mask {
    switch mv {
    case "0":
      carr[i] = barr[i]
    case "1":
      carr[i] = "1"
    case "X":
      carr[i] = "X"
    }
  }

  return carr
}

func (c Computer) TransformBinaryArrayThroughMask (barr []string, mask pie.Strings) int {
  carr := pie.Strings(make([]string, len(barr)))

  for i, mv := range mask {
    switch mv {
    case "0":
      carr[i] = "0"
    case "1":
      carr[i] = "1"
    case "X":
      carr[i] = barr[i]
    }
  }
  cval, _ := strconv.ParseInt(carr.Join(""), 2, 64)

  return int(cval)
}


// ---------- UTILITIES -----------------------------------

func (c Computer) ConvertDecimalToBinaryArray (val int) pie.Strings {
  bval := strconv.FormatInt(int64(val), 2)
  bval  = fmt.Sprintf("%036d", bval)
  bval  = strings.Replace(bval, "%!d(string=", "", -1)
  bval  = strings.Replace(bval, ")", "", -1)
  barr := strings.Split(bval, "")

  return pie.Strings(barr)
}

func (c Computer) ConvertFloatingMaskToAddresses (fmask pie.Strings) pie.Ints {
  addrs := pie.Ints{}

  xindices := pie.Ints{}
  for i, char := range fmask {
    if char == "X" {
      xindices = append(xindices, i)
    }
  }

  parrs := c.GenerateBinaryPermutations(len(xindices))
  for _, parr := range parrs {
    karr := pie.Strings(make([]string, len(fmask)))
    copy(karr, fmask)

    for i, bchar := range parr {
      xi      := xindices[i]
      karr[xi] = bchar
    }

    addr, _ := strconv.ParseInt(karr.Join(""), 2, 64)
    addrs    = append(addrs, int(addr))
  }

  return addrs
}

func (c Computer) GenerateBinaryPermutations (length int) []pie.Strings {
  perms := make([]pie.Strings, 0)
  for i := 0; i < length; i++ {
    if i == 0 {
      perms  = make([]pie.Strings, 0)
      perms  = append(perms, pie.Strings{"0"})
      perms  = append(perms, pie.Strings{"1"})
    } else {
      tmp := make([]pie.Strings, 0)
      for _, perm := range perms {
        size := len(perm)
        c0   := make(pie.Strings, size)
        c1   := make(pie.Strings, size)

        copy(c0, perm)
        copy(c1, perm)

        c0  = append(c0, "0")
        c1  = append(c1, "1")
        tmp = append(tmp, c0)
        tmp = append(tmp, c1)
      }
      perms = make([]pie.Strings, len(tmp))
      copy(perms, tmp)
    }
  }

  return perms
}

func (c Computer) SplitCommand (cmd string) (string, string) {
  cmd    = strings.Replace(cmd, "mem[", "", -1)
  cmd    = strings.Replace(cmd, "]", "", -1)
  parts := strings.Split(cmd, " = ")

  return parts[0], parts[1]
}

func (c Computer) SplitMask (mask string) pie.Strings {
  parts := strings.Split(mask, "")

  return pie.Strings(parts)
}

func (c Computer) Sum () int {
  sum := 0
  for _, v := range c.memory {
    sum = sum + v
  }
  return sum
}
