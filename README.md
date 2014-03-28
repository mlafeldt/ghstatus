go-ghstatus
===========

This Go library allows you to check the system status of GitHub from your own
applications. The status information is retrieved from GitHub's [system status
API].

The project also comes with a simple command-line tool named `ghstatus` that
utilizes the Go library.

## Installation

First, make sure you have [Go] installed.

### Library

To download and install the go-ghstatus library from source, simply run:

    $ go get github.com/mlafeldt/go-ghstatus

### Client

To install the `ghstatus` command-line tool, run this:

    $ go get github.com/mlafeldt/go-ghstatus/cmd/ghstatus

## Usage

### Library

For usage and examples, see the [Godoc] for this package.

### Client

To learn how to use the `ghstatus` tool, run `ghstatus --help`. The output will
look like this:

```
NAME:
   ghstatus - Check the system status of GitHub from the command line

USAGE:
   ghstatus [global options] command [command options] [arguments...]

COMMANDS:
   status, s	Show current system status (default command)
   messages, m	Show recent human communications
   last, l	Show last human communication
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --version, -v	print the version
   --help, -h		show help
```

To find out more about a specific command, execute `ghstatus help <command>`.

## Testing

[![Build Status](https://travis-ci.org/mlafeldt/go-ghstatus.svg?branch=master)](https://travis-ci.org/mlafeldt/go-ghstatus)

You can run the tests this way:

    $ cd go-ghstatus/
    $ ./script/test

## License and Author

Author:: Mathias Lafeldt (<mathias.lafeldt@gmail.com>)

Copyright:: 2013-2014, Mathias Lafeldt

go-ghstatus is licensed under the terms of the MIT License. See `LICENSE` file.

## Contributing

Please see `CONTRIBUTING.md` for details.


[Go]: http://golang.org/doc/install
[Godoc]: http://godoc.org/github.com/mlafeldt/go-ghstatus
[system status API]: https://status.github.com/api
