package day16

import (
  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type FieldMap map[string]pie.Ints


// ========== RECEIVERS ===================================

func (fm FieldMap) Consolidate () FieldMap {
  if fm.NeedsConsolidation() {
    kpairs := fm.KnownPairs()
    for k, vs := range kpairs {
      fm = fm.RestrictValueToKey(vs.First(), k)
    }
    return fm.Consolidate()
  } else {
    return fm
  }
}

func (fm FieldMap) RestrictValueToKey (value int, key string) FieldMap {
  for k, vs := range fm {
    if k != key {
      fm[k] = vs.Filter(func (v int) bool { return v != value })
    }
  }
  return fm
}


// ---------- UTILITIES -----------------------------------

func (fm FieldMap) KnownPairs () FieldMap {
  kps := FieldMap{}
  for name, indices := range fm {
    if indices.Len() == 1 {
      kps[name] = indices
    }
  }
  return kps
}

func (fm FieldMap) NeedsConsolidation () bool {
  result := false
  for _, v := range fm {
    if v.Len() > 1 {
      result = true
      break
    }
  }
  return result
}
