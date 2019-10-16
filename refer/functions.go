package main

import "fmt"

func plus(a int, b int) int {
  return a + b
}

func morePlus(a, b, c int) int {
  return a + b + c
}

func vals() (int, int, int) {
  return 3, 7, 10
}

// variadic with (args ...type)

func sum(nums ...int) {
  fmt.Print(nums, "\n")
  total := 0
  for _, num := range nums {
    total += num
  }
  fmt.Println(total)
}

func main() {
  res := plus(1, 2)
  fmt.Println("1 + 2 =", res)

  res = morePlus(1, 2, 3)
  fmt.Println("1 + 2 + 3 =", res)
 
  _, b, c := vals()

  fmt.Println(b, c)

  nums := []int{1, 2, 3, 4}

  sum(nums...)
}
