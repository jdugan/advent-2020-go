package day16

import (
  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITIONS =================================

type RangeMap map[int]bool

type RuleSet map[string]RangeMap


// ========== RECEIVERS ===================================

func (rs RuleSet) AnalyseTicket (t Ticket) FieldMap {
  fmap := rs.DefaultFieldMap()
  for i, v := range t.values {
    fields := rs.Analyse(v)
    for _, name := range fields {
      fmap[name] = append(fmap[name], i)
    }
  }
  return fmap
}

func (rs RuleSet) ValidateTicket (t Ticket) int {
  values := pie.Ints{}
  for _, v := range t.values {
    if !rs.Validate(v) {
      values = append(values, v)
    }
  }
  return values.Sum()
}


// ---------- PROCESSES -----------------------------------

func (rs RuleSet) Analyse (v int) pie.Strings {
  names := pie.Strings{}
  for k, rmap := range rs {
    _, exists := rmap[v]
    if exists {
      names = append(names, k)
    }
  }
  return names
}

func (rs RuleSet) Validate (v int) bool {
  result := false
  for _, rmap := range rs {
    _, exists := rmap[v]
    if exists {
      result = true
      break
    }
  }
  return result
}


// ---------- UTILITIES -----------------------------------

func (rs RuleSet) DefaultFieldMap () FieldMap {
  fmap := FieldMap{}
  for k, _ := range rs {
    fmap[k] = pie.Ints{}
  }
  return fmap
}

func (rs RuleSet) IntersectableFieldMap (values pie.Ints) FieldMap {
  fmap := FieldMap{}
  for k, _ := range rs {
    fmap[k] = pie.Ints(values)
  }
  return fmap
}
