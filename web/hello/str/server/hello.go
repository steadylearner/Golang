package main

// https://scene-si.org/2018/01/25/go-tips-and-tricks-almost-everything-about-imports/
import (
	"fmt"

	"log"
	"net/http"

	. "github.com/logrusorgru/aurora"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %s\n", r.URL.Path[1:]) // Remove / from r.URL.Path with [1:]
	})

	const port = ":8080"

	fmt.Printf("\n🚀 Golang Server ready at %s", Blue("http://localhost"+port))
	log.Fatal(http.ListenAndServe(port, nil))
}
