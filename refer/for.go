package main

import "fmt"

func main() {
  i := 1
  for i <= 3 {
    fmt.Println(i)
    i++
  }

  for j := 7; j <= 8; j++ {
    fmt.Println(j)
  }

  for {
    fmt.Println("While")
    break
  }

  for n := 0; n <= 5; n++ {
    if n % 5 == 0 {
      continue
    }

    fmt.Println(n)
  }
}
