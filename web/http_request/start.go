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

	// [FORM]
	"net/url"

	// [UPLOAD FILE]
	"io"
	"mime/multipart"
	"os"

	// [TIMEOUT]
	"time"
)

func main() {
	// MakeGetRequest()
	// MakeCustomGetRequest()
	// MakePostRequest()
	// MakeFormRequest()
	MakeFileUploadRequest()
}

func logError(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

// if err != nil {
// 	log.Fatalln(err)
// }

// [GET]
func MakeGetRequest() {
	response, err := http.Get("https://httpbin.org/get")
	logError(err)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	logError(err)

	log.Println(string(body))
}

func MakeCustomGetRequest() {

	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	// type Client struct {
	// 	Transport RoundTripper
	// 	CheckRedirect func(req *Request, via []*Request) error
	// 	Jar CookieJar
	// 	Timeout time.Duration
	// }
	//  Should set Timeout part

	request, err := http.NewRequest("GET", "https://httpbin.org/get", nil) // nil part is body of the request
	logError(err)

	response, err := client.Do(request)
	logError(err)
	defer response.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)
	log.Println(result)
}

// [POST]
func MakePostRequest() {
	message := map[string]interface{}{
		"hello": "world",
		"life":  42,
		"embedded": map[string]string{
			"yes": "of course!",
		},
	}

	bytesRepresentation, err := json.Marshal(message)
	logError(err)

	response, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(bytesRepresentation))
	logError(err)
	defer response.Body.Close()

	var result map[string]interface{}

	json.NewDecoder(response.Body).Decode(&result)

	log.Println(result)
	log.Println("\n", result["data"])
}

// [FORM]
func MakeFormRequest() {

	formData := url.Values{
		"name": {"steadylearner"},
	}

	response, err := http.PostForm("https://httpbin.org/post", formData)
	logError(err)
	defer response.Body.Close()

	var result map[string]interface{}

	json.NewDecoder(response.Body).Decode(&result)

	log.Println(result["form"])
}

// [UPLOAD FILE] - should read more documenations
func MakeFileUploadRequest() {

	// You should make the file before with $touch name.txt
	// Open the file
	file, err := os.Open("name.txt")
	logError(err)
	// Close the file later
	defer file.Close()

	// Buffer to store our request body as bytes
	var requestBody bytes.Buffer

	// Create a multipart writer
	multiPartWriter := multipart.NewWriter(&requestBody)

	// Initialize the file field
	fileWriter, err := multiPartWriter.CreateFormFile("file_field", "name.txt")
	logError(err)

	// Copy the actual file content to the field field's writer
	_, err = io.Copy(fileWriter, file)
	logError(err)

	// Populate other fields
	fieldWriter, err := multiPartWriter.CreateFormField("normal_field")
	logError(err)

	_, err = fieldWriter.Write([]byte("Value"))
	logError(err)

	// We completed adding the file and the fields, let's close the multipart writer
	// So it writes the ending boundary
	multiPartWriter.Close()

	// By now our original request body should have been populated, so let's just use it with our custom request
	req, err := http.NewRequest("POST", "https://httpbin.org/post", &requestBody)
	logError(err)
	// We need to set the content type from the writer, it includes necessary boundary as well
	req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())

	// Do the request
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	// client := &http.Client{}
	response, err := client.Do(req)
	logError(err)
	defer response.Body.Close()

	var result map[string]interface{}

	json.NewDecoder(response.Body).Decode(&result)

	log.Println(result)
}
