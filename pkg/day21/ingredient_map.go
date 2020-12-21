package day21

import (
  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION(S) ===============================

type IngredientMap map[string]Ingredient


// ========== RECEIVERS ===================================

func (im IngredientMap) Allergenics () Ingredients {
  matches := Ingredients{}
  for _, i := range im {
    if i.allergen != "" {
      matches = append(matches, i)
    }
  }
  return matches
}

func (im IngredientMap) NonAllergenics () pie.Strings {
  matches := pie.Strings{}
  for _, i := range im {
    if i.allergen == "" {
      matches = append(matches, i.name)
    }
  }
  return matches
}


// ---------- BUILD HELPERS -------------------------------

func (im IngredientMap) ApplyAllergenMap (am AllergenMap) IngredientMap {
  for aname, inames := range am {
    if inames.Len() == 1 {
      iname         := inames.First()
      ingredient, _ := im[iname]

      ingredient.allergen = aname
      im[iname] = ingredient
    }
  }
  return im
}


func (im IngredientMap) RegisterNamesIfMissing (names pie.Strings) IngredientMap {
  for _, name := range names {
    _, exists := im[name]
    if !exists {
      im[name] = Ingredient{name: name, allergen: ""}
    }
  }
  return im
}
