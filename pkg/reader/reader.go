package reader

import (
  "bufio"
  "fmt"
  "os"
  "strings"

  "github.com/elliotchance/pie/pie"
)


// ========== IO/OS HELPERS ===============================

func Lines (path string) pie.Strings {
  lines := make([]string, 0)

  file, err := os.Open(path)
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    str  := strings.TrimSpace(scanner.Text())
    lines = append(lines, str)
  }

  if err := scanner.Err(); err != nil {
    fmt.Println(err)
  }

  return pie.Strings(lines)
}
