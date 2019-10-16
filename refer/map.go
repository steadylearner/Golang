package main

import "fmt"

func main() {

    m := make(map[string]string)

    m["website"] = "www.steadylearner.com"

    fmt.Println("map:", m)

    website := m["website"]
    fmt.Println("website: ", website)

    fmt.Println("len:", len(m))

    delete(m, "website")
    fmt.Println("map:", m)
}
