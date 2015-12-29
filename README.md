# ghstatus

[![Build Status](https://travis-ci.org/mlafeldt/ghstatus.svg?branch=master)](https://travis-ci.org/mlafeldt/ghstatus)
[![GoDoc](https://godoc.org/github.com/mlafeldt/ghstatus?status.svg)](https://godoc.org/github.com/mlafeldt/ghstatus)

This Go library allows you to check the system status of GitHub from your own
applications and monitoring services. The status information is retrieved from
GitHub's [system status API].

The project also comes with a simple command-line tool named `ghstatus` that
utilizes the Go library.

## Installation

First, make sure you have [Go] installed.

### Library

To download and install the ghstatus library from source, simply run:

    $ go get github.com/mlafeldt/ghstatus

### Client

To install the `ghstatus` command-line tool, run this:

    $ go get github.com/mlafeldt/ghstatus/cmd/ghstatus

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

You can run the tests this way:

    $ cd ghstatus/
    $ make test

## License

ghstatus is licensed under the terms of the MIT License. See [LICENSE](/LICENSE)
file.

## Contributing

Please see [CONTRIBUTING.md](/CONTRIBUTING.md) for details.


[Go]: http://golang.org/doc/install
[Godoc]: http://godoc.org/github.com/mlafeldt/ghstatus
[system status API]: https://status.github.com/api
