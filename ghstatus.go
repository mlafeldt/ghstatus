package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Message struct {
	Status    string
	Body      string
	CreatedOn string
}

func sendRequest(endpoint string) ([]byte, error) {
	resp, err := http.Get("https://status.github.com/api/" + endpoint + ".json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	log.SetFlags(0)

	body, err := sendRequest("messages")
	if err != nil {
		log.Fatal(err)
	}

	var messages []Message
	err = json.Unmarshal(body, &messages)
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range messages {
		log.Printf("%v", m)
	}
}
