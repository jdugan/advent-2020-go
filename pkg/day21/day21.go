package day21

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
  fmt.Println("DAY", 21)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  products, imap := groceries()
  inames         := imap.NonAllergenics()

  return products.CountOccurrences(inames)
}

func Puzzle2() string {
  _, imap     := groceries()
  allergenics := imap.Allergenics()
  names       := allergenics.SortedNames()

  return names.Join(",")
}


// ========== PRIVATE FNS =================================

func groceries () (Products, IngredientMap) {
  lines := reader.Lines("./data/day21/input.txt")
  re    := regexp.MustCompile(`^(.+) \(contains (.+)\)$`)

  products    := Products{}
  ingredients := IngredientMap{}
  allergens   := AllergenMap{}

  inames      := pie.Strings{}
  anames      := pie.Strings{}

  for _, line := range lines {
    groups := re.FindAllStringSubmatch(line, 1)
    if len(groups) > 0 {
      group := groups[0]
      inames = pie.Strings(strings.Split(group[1], " "))
      anames = pie.Strings(strings.Split(group[2], ", "))
    } else {
      inames = pie.Strings(strings.Split(line, " "))
      anames = pie.Strings{}
    }

    products    = append(products, Product{ingredients: inames})
    ingredients = ingredients.RegisterNamesIfMissing(inames)
    allergens   = allergens.RegisterEntriesAndReduce(anames, inames)
  }

  ingredients = ingredients.ApplyAllergenMap(allergens)

  return products, ingredients
}
