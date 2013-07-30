package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const StatusApiUrl = "https://status.github.com/api"

const (
	StatusGood  = "good"
	StatusMinor = "minor"
	StatusMajor = "major"
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

func sendRequest(endpoint string, v interface{}) error {
	resp, err := http.Get(StatusApiUrl + endpoint + ".json")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &v)
}

func GetStatus() (*Status, error) {
	var status *Status
	if err := sendRequest("/status", &status); err != nil {
		return nil, err
	}
	return status, nil
}

func GetMessages() ([]Message, error) {
	var messages []Message
	if err := sendRequest("/messages", &messages); err != nil {
		return nil, err
	}
	return messages, nil
}

func GetLastMessage() (*Message, error) {
	var message *Message
	if err := sendRequest("/last-message", &message); err != nil {
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
