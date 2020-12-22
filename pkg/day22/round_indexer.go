package day22


// ========== DEFINITION ==================================

type RoundIndexer map[string]bool


// ========== RECEIVERS ===================================

func (ri RoundIndexer) Register (r Round) (RoundIndexer, bool) {
  success   := true
  k         := r.Id()
  _, exists := ri[k]

  if exists {
    success = false     // duplicate found
  } else {
    ri[k] = true
  }

  return ri, success
}
