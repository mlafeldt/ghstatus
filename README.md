go-ghstatus
===========

This Go library allows you to check the system status of GitHub from your own
applications. The status information is retrieved from GitHub's [system status
API].


Installation
------------

First, make sure you have [Go 1] installed.

To download and install go-ghstatus from source, simply run:

    $ go get github.com/mlafeldt/go-ghstatus


Usage
-----

For usage and examples see the [Godoc] for this package.


Testing
-------

The tests require [testify] to be installed, so get it first:

    $ go get github.com/stretchr/testify

Now you can run the tests as usual:

    $ cd go-ghstatus/
    $ go test ./...

Alternatively, use the `test` script which does all of the steps above:

    $ cd go-ghstatus/
    $ ./script/test


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
[Godoc]: http://godoc.org/github.com/mlafeldt/go-ghstatus
[LICENSE]: https://github.com/mlafeldt/go-ghstatus/blob/master/LICENSE
[system status API]: https://status.github.com/api
[testify]: https://github.com/stretchr/testify
