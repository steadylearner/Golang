package main

import (
  "fmt"
  "os"
  // "strings"
)

func main() {
  os.Setenv("WEBSITE", "www.steadylearner")
  fmt.Println("WEBSITE:", os.Getenv("WEBSITE"))

  // fmt.Println()
  // for _, e := range os.Environ() {
  //     pair := strings.SplitN(e, "=", 2)
  //    fmt.Println(pair[0])
  // }
}
