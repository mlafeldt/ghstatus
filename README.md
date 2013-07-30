go-ghstatus
===========

A Go library for GitHub's [System Status API].


Installation
------------

First, make sure you have [Go 1] installed.

To download and install go-ghstatus from source, simply run:

    $ go get github.com/mlafeldt/go-ghstatus


Usage
-----

```go
import "github.com/mlafeldt/go-ghstatus"

// Get current system status (one of good, minor, or major) and timestamp.
status, err := GetStatus()

// Get most recent human communications with status and timestamp.
messages, err := GetMessages()

// Get last human communication, status, and timestamp.
message, err := GetLastMessage()
```


License
-------

go-ghstatus is licensed under the terms of the MIT License. See [LICENSE] file.


Contact
-------

* Web: <https://github.com/mlafeldt/go-ghstatus>
* Mail: <mathias.lafeldt@gmail.com>
* Twitter: [@mlafeldt](https://twitter.com/mlafeldt)


[Go 1]: http://golang.org/doc/install
[LICENSE]: https://github.com/mlafeldt/go-ghstatus/blob/master/LICENSE
[System Status API]: https://status.github.com/api
