package main

import (
	"encoding/json"
	"log"
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

	var m []Message
	err := json.Unmarshal(data, &m)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%d", len(m))
	log.Printf("%v", m)
}
