// http://polyglot.ninja/golang-making-http-requests/

package main

import (
	// [GET]
	"io/ioutil"
	"log"
	"net/http"

	// [POST]
	"bytes"
	"encoding/json"
)

func main() {
	MakePostRequest()
}

func logError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}


func MakePostRequest() {
	message := map[string]interface{}{
		"payload": "from www.steadylearner.com",
	}

	bytesRepresentation, err := json.Marshal(message)
	logError(err)

	response, err := http.Post("http://localhost:8080/decode", "application/json", bytes.NewBuffer(bytesRepresentation))
	logError(err)

	body, err := ioutil.ReadAll(response.Body)
	logError(err)
	defer response.Body.Close()

	log.Println(string(body))
}