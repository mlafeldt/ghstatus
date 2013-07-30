package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Message struct {
	Status      string
	Body        string
	CreatedOn   string
}

var data = []byte(`
[
    {
        "body": "Everything operating normally.",
        "created_on": "2013-07-29T22:23:19Z",
        "status": "good"
    },
    {
        "body": "We are continuing to work on the increased exception rate on the GitHub API. We will update again as soon as have the source of these exceptions resolved. ",
        "created_on": "2013-07-29T21:09:54Z",
        "status": "minor"
    }
]
`)

func main() {
	log.SetFlags(0)

	resp, err := http.Get("https://status.github.com/api/messages.json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var messages []Message
	err = json.Unmarshal(body, &messages)
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range messages {
		log.Printf("%v", m)
	}
}
