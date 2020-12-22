package day22


// ========== DEFINITION ==================================

type Result struct {
  winner  int
  score   int
}


// ========== RECEIVERS ===================================

func (r Result) Forfeit () Result {
  return Result{winner: 1, score: -1}
}
