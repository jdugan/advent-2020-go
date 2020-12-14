package day13


// ========== DEFINITION ==================================

type Bus struct {
  id        int
  offset    int
}


// ========== RECEIVERS ===================================

func (b Bus) Match (wt int) bool {
  if wt > b.offset {
    return b.offset % wt == 0
  } else {
    return b.offset == wt
  }
}

func (b Bus) WaitTime (etd int) int {
  wt := b.id - (etd % b.id)

  return wt
}
