package main

import (
	"encoding/json"
	"fmt"
	//	"encoding/json"
	"io/ioutil"
	"log"
	//	"net/url"
	"net/http"
	//	"time"
)


func main() {
	//getRequest()
	//MakeRequest()
	getRequest2("tevora.com")
}

type wayurl struct {
	url	string
}

func getRequest() ([]wayurl, error) {
	resp, err := http.Get( "http://web.archive.org/cdx/search/cdx?url=*.tevora.com/*&output=json&fl=original&collapse=urlkey")
	if err != nil {
		log.Fatal("Error getting response. ", err)
	}
	/*client := &http.Client{Timeout: time.Second * 5}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	} */
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log .Fatal("Error reading body. ", err)
	}

	// create mother of all slices to hold slices.
	var wrap [][]string
	err = json.Unmarshal(body, &wrap)
	// create array with size of response len. Refers to OG wayurl struct.
	out := make([]wayurl, 0, len(wrap))
	for _, urls := range wrap {
		out = append(out, wayurl{url: urls[0]})
	}
	fmt.Printf("%s\n", out)

	return out, nil
}

func getRequest2(domain string) ([]wayurl, error) {
	resp, err := http.Get( "http://web.archive.org/cdx/search/cdx?url=*.tevora.com/*&output=json&fl=original&collapse=urlkey")
	if err != nil {
		log.Fatal("Error getting response. ", err)
	}
	defer resp.Body.Close()

	// make slice with 0 length and no cap. An empty slice.
	output := make([]wayurl, 0)

	decode := json.NewDecoder(resp.Body)

/*	// create a composite literal of type struct{}. Costs no memory.
	wrapper := struct {
		url string
	}{} */

	// make empty slice holding slice
	var wrapper [][]string

	err = decode.Decode(&wrapper)
	if err != nil {
		return []wayurl{}, err
	}

	for _, urls := range wrapper {
		output = append(output,wayurl{url: urls[0]})
	}
	fmt.Printf("%s\n", output)
	return output, nil

}


func MakeRequest() {

	client := http.Client{}
	resp, err := client.Get("http://web.archive.org/cdx/search/cdx?url=*.tevora.com/*&output=json&fl=original&collapse=urlkey")
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println(result)
}

/*
func oldGetRequest() {
	resp, err := http.Get( "http://web.archive.org/cdx/search/cdx?url=*.tevora.com/*&output=json&fl=original&collapse=urlkey")
	if err != nil {
		log.Fatal("Error getting response. ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log .Fatal("Error reading body. ", err)
	}
	fmt.Printf("%s\n", body)
} */
