package day20

import (
  "fmt"
  "regexp"
  "strconv"
  "strings"

  "github.com/elliotchance/pie/pie"

  "../../pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 20)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  ca    := cameraArray()
  ctids := ca.CornerTileIds()

  return ctids.Product()
}

func Puzzle2() int {
  bca := cameraArray()
  lca := bca.LinkTiles()
  tb0 := lca.ConsolidateTiles()

  ca := CameraArray{}
  ca.Initialise()
  ca.RegisterTilePossibilities(tb0)

  analyser := Analyser{tiles: ca.tiles}
  count    := analyser.SeaRoughness()

  return count
}


// ========== PRIVATE FNS =================================

func cameraArray () CameraArray {
  lines := reader.Lines("./data/day20/input.txt")

  re0  := regexp.MustCompile(`^Tile (\d+):$`)
  ca   := CameraArray{tiles: make(map[string]Tile)}
  t    := Tile{}

  for _, line := range lines {
    if line == "" {
      t.height = t.height + 1
      ca.RegisterTilePossibilities(t)
    } else {
      groups := re0.FindAllStringSubmatch(line, -1)
      if len(groups) > 0 {
        id, _ := strconv.Atoi(groups[0][1])
        t = Tile{}
        t.Initialise(id)
      } else {
        cvs := pie.Strings(strings.Split(line, ""))
        t.AddRow(cvs)
      }
    }
  }
  t.height = t.height + 1
  ca.RegisterTilePossibilities(t)

  return ca
}
