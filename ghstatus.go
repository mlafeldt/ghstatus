package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Status struct {
	Status      string `json:"status"`
	LastUpdated string `json:"last_updated"`
}

type Message struct {
	Status    string `json:"status"`
	Body      string `json:"body"`
	CreatedOn string `json:"created_on"`
}

func sendRequest(endpoint string) ([]byte, error) {
	resp, err := http.Get("https://status.github.com/api/" + endpoint + ".json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func GetStatus() (*Status, error) {
	body, err := sendRequest("status")
	if err != nil {
		return nil, err
	}

	var status *Status
	err = json.Unmarshal(body, &status)
	if err != nil {
		return nil, err
	}
	return status, nil
}

func GetMessages() ([]Message, error) {
	body, err := sendRequest("messages")
	if err != nil {
		return nil, err
	}

	var messages []Message
	err = json.Unmarshal(body, &messages)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func GetLastMessage() (*Message, error) {
	body, err := sendRequest("last-message")
	if err != nil {
		return nil, err
	}

	var message *Message
	err = json.Unmarshal(body, &message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func main() {
	log.SetFlags(0)

	status, err := GetStatus()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", status)

	messages, err := GetMessages()
	if err != nil {
		log.Fatal(err)
	}
	for _, m := range messages {
		log.Printf("%+v", m)
	}

	message, err := GetLastMessage()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", message)
}
