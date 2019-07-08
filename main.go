package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	//GET request
	response, err := http.Get("https://api.coinbase.com/v2/prices/spot?currency=USD")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	//GET TYPE 2 - with API key
	req, err := http.NewRequest("GET", "https://YOUR_API_GATEWAY.execute-api.us-east-1.amazonaws.com/Dev/shared/getyears", nil)

	// add authorization header to the req
	req.Header.Add("x-api-key", "YOUR_API_KEY")

	// Send the req using http Client
	client2 := &http.Client{}
	resp, err := client2.Do(req)
	if err != nil {
		fmt.Printf("2. The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data))
	}

	//POST request
	jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
	jsonValue, _ := json.Marshal(jsonData)

	// post type 1
	response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))

	// post type 2
	request, _ := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err = client.Do(request)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

}
