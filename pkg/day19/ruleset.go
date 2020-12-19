package day19

import (
  "fmt"
  "regexp"
  "strconv"
  "strings"

  "github.com/elliotchance/pie/pie"
)


// ========== DEFINITIONS =================================

type RuleSet map[int]string


// ========== RECEIVERS ===================================

// Here's what we want to get.
//
// 0: 4 1 5
// 1: 2 3 | 3 2
// 2: 4 4 | 5 5
// 3: 4 5 | 5 4
// 4: "a"
// 5: "b"
//
// 0: "(a((aa|bb)(ab|ba)|(ab|ba)(aa|bb))b)"
// 1: "((aa|bb)(ab|ba)|(ab|ba)(aa|bb))"
// 2: "(aa|bb)"
// 3: "(ab|ba)"
// 4: "a"
// 5: "b"
//
func (rs RuleSet) Expand () RuleSet {
  ers := RuleSet{}
  frs := RuleSet{}
  wrs := RuleSet{}

  // first, separate rules that are fully
  // expanded from those that are not
  for k, v := range rs {
    pending, _ := regexp.MatchString(`\d`, v)
    if pending {
      wrs[k] = v
    } else {
      ers[k] = v
      frs[k] = v
    }
  }

  // then loop the unexpanded items and
  // replace as much as we can
  re0   := regexp.MustCompile(`(\d+)`)
  pkeys := pie.Ints{}
  done  := true
  for k, v := range wrs {
    groups := re0.FindAllStringSubmatch(v, -1)
    ikeys  := pie.Ints{}

    for _, garr := range groups {
      skey      := garr[0]
      ikey, _   := strconv.Atoi(skey)
      _, exists := ers[ikey]
      if exists {
        ikeys = append(ikeys, ikey)
      } else {
        break
      }
    }

    if ikeys.Len() > 0 && ikeys.Len() == len(groups) {
      fopts := pie.Strings{}
      wopts := strings.Split(v, " | ")
      for _, opt := range wopts {
        ogroups := re0.FindAllStringSubmatch(opt, -1)
        for _, ogarr := range ogroups {
          skey    := ogarr[0]
          ikey, _ := strconv.Atoi(skey)
          rule, _ := frs[ikey]

          opt  = pie.Strings{" ", opt, " "}.Join("")
          skey = pie.Strings{" ", skey, " "}.Join("")
          rule = pie.Strings{" ", rule, " "}.Join("")
          opt  = strings.Replace(opt, skey, rule, -1)
        }
        opt   = strings.Replace(opt, " ", "", -1)
        smatch, _ := regexp.MatchString(`^\(`, opt)
        ematch, _ := regexp.MatchString(`\)$`, opt)
        if (smatch && !ematch) || (ematch && !smatch) {
          opt = pie.Strings{"(", opt, ")"}.Join("")
        }
        fopts = append(fopts, opt)
      }
      v = fopts.Join("|")
      v = pie.Strings{"(", v, ")"}.Join("")
    } else {
      done  = false
      pkeys = append(pkeys, k)
    }

    frs[k] = v
  }

  // if all unprocessed keys are unprocessable,
  // then the remaining values involve recursion.
  // break out for now.
  if pkeys.Len() == wrs.OrderedKeys().Len() {
    done = true
  }

  // if everything expanded, quit; else,
  // recurse until that's true.
  if done {
    return frs
  } else {
    return frs.Expand()
  }
}


// ---------- UTILITIES -----------------------------------

func (rs RuleSet) OrderedKeys () pie.Ints {
  keys := pie.Ints{}
  for k, _ := range rs {
    keys = append(keys, k)
  }

  return keys.Sort()
}

func (rs RuleSet) Print () {
  fmt.Println("==========================")
  for _, key := range rs.OrderedKeys() {
    fmt.Println(key, "=>", rs[key])
  }
  fmt.Println("==========================")
}
