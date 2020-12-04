package day04

import (
  "regexp"

  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITION ==================================

type Passport struct {
  fields  map[string]string
}


// ========== RECEIVERS ===================================

func (p Passport) Keys () pie.Strings {
  keys := make([]string, 0)
  for k, _ := range p.fields {
    if k != "cid" {
      keys = append(keys, k)
    }
  }
  return pie.Strings(keys).Sort()
}

func (p Passport) ValidateFormats (rkeys pie.Strings) bool {
  result := p.ValidateKeys(rkeys)

  rxmap       := make(map[string]*regexp.Regexp)
  rxmap["byr"] = regexp.MustCompile(`^(19[2-9][0-9]|200[0-2])$`)
  rxmap["iyr"] = regexp.MustCompile(`^(201[0-9]|2020)$`)
  rxmap["eyr"] = regexp.MustCompile(`^(202[0-9]|2030)$`)
  rxmap["hgt"] = regexp.MustCompile(`^((1([5-8][0-9]|9[0-3])cm)|((59|6[0-9]|7[0-6])in))$`)
  rxmap["hcl"] = regexp.MustCompile(`^#[0-9a-f]{6}$`)
  rxmap["ecl"] = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
  rxmap["pid"] = regexp.MustCompile(`^\d{9}$`)

  if result {
    for _, k := range rkeys {
      rxp := rxmap[k]
      val := p.fields[k]

      result = rxp.MatchString(val)
      if !result {
        break
      }
    }
  }

  return result
}

func (p Passport) ValidateKeys (rkeys pie.Strings) bool {
  ikeys := p.Keys().Intersect(rkeys).Sort()
  return rkeys.Equals(ikeys)
}
