// These smoke tests start an internal web server that returns fake responses
// in order to check the output of the different ghstatus commands.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	ghstatus "github.com/mlafeldt/go-ghstatus"
)

func init() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "method must be GET", http.StatusMethodNotAllowed)
			return
		}
		content, err := ioutil.ReadFile("../../testdata" + r.URL.Path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(content))
	}))
	ghstatus.SetServiceURL(ts.URL)
}

func ExampleStatus() {
	runApp([]string{"ghstatus", "--status"})
	// Output:
	// [Mar 23 19:43:34] good
}

func ExampleMessages() {
	runApp([]string{"ghstatus", "--messages"})
	// Output:
	// [Mar 22 03:35:14] good Everything operating normally.
	// [Mar 22 01:04:47] minor As a result of our ongoing DDoS mitigation, we're experiencing high rates of packet loss from users in the Asia-Pacific region. We're working on reducing this disruption to service and will provide additional information when it becomes available.
	// [Mar 21 14:11:45] good GitHub.com remains stable as we continue to mitigate an ongoing DDoS attack. At this time we're discontinuing status updates, but will resume them if the situation changes.
	// [Mar 21 13:56:38] minor GitHub.com performance has been back to normal for some time. We continue to work to tune our mitigation to make sure that there is as little impact to legitimate traffic as possible.
	// [Mar 21 13:35:16] minor We are continuing to work to mitigate the attack and reduce the number of legitimate users who are flagged as attack traffic.
	// [Mar 21 13:13:42] minor We are currently working to mitigate an incoming DDoS attack. We'll provide additional information as it becomes available.
}

func ExampleLastMessage() {
	runApp([]string{"ghstatus", "--last"})
	// Output:
	// [Mar 22 03:35:14] good Everything operating normally.
}
