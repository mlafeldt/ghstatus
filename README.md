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

[![Build Status](https://travis-ci.org/mlafeldt/go-ghstatus.svg?branch=master)](https://travis-ci.org/mlafeldt/go-ghstatus)

You can run the tests this way:

    $ cd go-ghstatus/
    $ ./script/test


License
-------

go-ghstatus is licensed under the terms of the MIT License. See [LICENSE] file.


Contributing
------------

Please see `CONTRIBUTING.md` for details.


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
