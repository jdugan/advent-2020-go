package day16

import (
  "fmt"
  "regexp"
  "strings"

  "github.com/elliotchance/pie/pie"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 16)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  _, _, _, its := data()

  codes := pie.Ints{}
  for _, t := range its {
    codes = append(codes, t.errorcode)
  }

  return codes.Sum()
}

func Puzzle2() int {
  rs, mt, vts, _ := data()

  // calculate possible indices for
  // each ticket
  fmaps := make([]FieldMap, 0)
  for _, t := range vts {
    fmap := rs.AnalyseTicket(t)
    fmaps = append(fmaps, fmap)
  }

  // intersect those disconnected possibilities
  // into a single map
  is    := mt.DefaultIndices()
  cfmap := rs.IntersectableFieldMap(is)
  for _, fm := range fmaps {
    for n, vs := range fm {
      cfmap[n] = cfmap[n].Intersect(vs).Sort()
    }
  }

  // consolidate the consolidated map further
  // until we have a distinct value for each
  // field name
  ccfmap := cfmap.Consolidate()

  // pluck out values for keys that refer
  // to departure info
  re  := regexp.MustCompile(`^departure `)
  dis := pie.Ints{}
  for name, indices := range ccfmap {
    if re.MatchString(name) {
      dis = append(dis, indices.First())
    }
  }

  // convert departure indices into
  // departure values
  dvs := pie.Ints{}
  for i, v := range mt.values {
    if dis.Contains(i) {
      dvs = append(dvs, v)
    }
  }

  return dvs.Product()
}


// ========== PRIVATE FNS =================================

func data () (RuleSet, Ticket, []Ticket, []Ticket) {
  lines := reader.Lines("./data/day16/input.txt")

  // rules
  rs  := RuleSet{}
  rls := lines[0:19]
  for _, rl := range rls {
    name, rmap := parseRuleLine(rl)
    rs[name] = rmap
  }

  // my ticket
  mtl  := lines[22]
  mt   := parseTicketLine(mtl)

  // other tickets
  its := make([]Ticket, 0)
  vts := make([]Ticket, 0)
  otls := lines[25:]
  for _, tl := range otls {
    t          := parseTicketLine(tl)
    t.errorcode = rs.ValidateTicket(t)
    if t.errorcode == 0 {
      vts = append(vts, t)
    } else {
      its = append(its, t)
    }
  }

  return rs, mt, vts, its
}

func parseRuleLine (rl string) (string, RangeMap) {
  re    := regexp.MustCompile(`^(.+): (\d+)-(\d+) or (\d+)-(\d+)$`)
  parts := re.FindStringSubmatch(rl)

  name  := parts[1]
  rvals := pie.Strings(parts[2:]).Ints()
  rmap  := RangeMap{}
  for i := rvals[0]; i <= rvals[1]; i++ {
    rmap[i] = true
  }
  for i := rvals[2]; i <= rvals[3]; i++ {
    rmap[i] = true
  }

  return name, rmap
}

func parseTicketLine (tl string) Ticket {
  tvs := strings.Split(tl, ",")
  tvi := pie.Strings(tvs).Ints()

  return Ticket{values: tvi}
}
