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
status, err := ghstatus.GetStatus()

// Get most recent human communications with status and timestamp.
messages, err := ghstatus.GetMessages()

// Get last human communication, status, and timestamp.
message, err := ghstatus.GetLastMessage()
```


License
-------

go-ghstatus is licensed under the terms of the MIT License. See [LICENSE] file.


Contributing
------------

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request


Contact
-------

* Web: <https://github.com/mlafeldt/go-ghstatus>
* Mail: <mathias.lafeldt@gmail.com>
* Twitter: [@mlafeldt](https://twitter.com/mlafeldt)


[Go 1]: http://golang.org/doc/install
[LICENSE]: https://github.com/mlafeldt/go-ghstatus/blob/master/LICENSE
[System Status API]: https://status.github.com/api
