package day20

import (
  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Monster struct {
  width     int
  offsets   []pie.Ints
}


// ========== RECEIVERS ===================================

func (m *Monster) Initialise () {
  m.offsets     = make([]pie.Ints, 15)
  m.offsets[0]  = pie.Ints{0, 0}
  m.offsets[1]  = pie.Ints{1, 1}
  m.offsets[2]  = pie.Ints{4, 1}
  m.offsets[3]  = pie.Ints{5, 0}
  m.offsets[4]  = pie.Ints{6, 0}
  m.offsets[5]  = pie.Ints{7, 1}
  m.offsets[6]  = pie.Ints{10, 1}
  m.offsets[7]  = pie.Ints{11, 0}
  m.offsets[8]  = pie.Ints{12, 0}
  m.offsets[9]  = pie.Ints{13, 1}
  m.offsets[10]  = pie.Ints{16, 1}
  m.offsets[11] = pie.Ints{17, 0}
  m.offsets[12] = pie.Ints{18, 0}
  m.offsets[13] = pie.Ints{18, -1}
  m.offsets[14] = pie.Ints{19,0}

  m.width  = 20
}
