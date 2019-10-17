// http://polyglot.ninja/golang-making-http-requests/

package main

import (
	// [GET]

	"io/ioutil"
	"log"
	"net/http"

	// [TIMEOUT]
	"time"
)

func logError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	MakeGetRequest()
}

func MakeGetRequest() {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("GET", "http://localhost:8080/from www.steadylearner.com", nil) // nil part is body of the request
	logError(err)

	response, err := client.Do(request)
	logError(err)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	logError(err)

	log.Println(string(body))
}

