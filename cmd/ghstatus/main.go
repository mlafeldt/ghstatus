package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/codegangsta/cli"
	"github.com/mlafeldt/go-ghstatus"
)

func printStatus(t time.Time, status, body string) {
	ts := t.Format(time.Stamp)
	if body != "" {
		fmt.Printf("[%s] %s %s\n", ts, status, body)
	} else {
		fmt.Printf("[%s] %s\n", ts, status)
	}
}

func exitWithStatus(status string) {
	code := map[string]int{
		ghstatus.Good:  0,
		ghstatus.Minor: 1,
		ghstatus.Major: 2,
	}[status]
	os.Exit(code)
}

func cmdStatus(c *cli.Context) {
	s, err := ghstatus.GetStatus()
	if err != nil {
		log.Fatal("error: failed to get status: ", err)
	}

	printStatus(s.LastUpdated, s.Status, "")

	if c != nil && c.Bool("exit-code") {
		exitWithStatus(s.Status)
	}
}

func cmdMessages(c *cli.Context) {
	messages, err := ghstatus.GetMessages()
	if err != nil {
		log.Fatal("error: failed to get messages: ", err)
	}

	for _, m := range messages {
		printStatus(m.CreatedOn, m.Status, m.Body)
	}
}

func cmdLastMessage(c *cli.Context) {
	m, err := ghstatus.GetLastMessage()
	if err != nil {
		log.Fatal("error: failed to get last message: ", err)
	}

	printStatus(m.CreatedOn, m.Status, m.Body)

	if c != nil && c.Bool("exit-code") {
		exitWithStatus(m.Status)
	}
}

func main() {
	log.SetFlags(0)

	app := cli.NewApp()
	app.Name = "ghstatus"
	app.Usage = "Check the system status of GitHub from the command line"
	app.Version = "1.5"
	app.Commands = []cli.Command{
		{
			Name:      "status",
			ShortName: "s",
			Action:    cmdStatus,
			Flags: []cli.Flag{
				cli.BoolFlag{
					"exit-code, e",
					"Make program exit with GitHub status as exit code",
				},
			},
		},
		{
			Name:      "messages",
			ShortName: "m",
			Action:    cmdMessages,
		},
		{
			Name:      "last",
			ShortName: "l",
			Action:    cmdLastMessage,
			Flags: []cli.Flag{
				cli.BoolFlag{
					"exit-code, e",
					"Make program exit with GitHub status as exit code",
				},
			},
		},
	}

	args := os.Args
	if len(args) < 2 {
		args = append(args, "status")
	}

	if err := app.Run(args); err != nil {
		log.Fatal(err)
	}
}
