## v1.7 (Dec 29 2015)

* Rename project to ghstatus.
* Introduce `Client` type that allows to configure a custom `http.Client` and
  service URL.
* Set lower HTTP timeout in CLI tool.
* Use `http.FileServer` for testing rather than re-implementing it.
* Replace scripts with `Makefile`.

## v1.6 (Mar 28 2014)

* Add `ghstatus` command-line tool.
* Add `script/build` script for easy cross-compilation.

## v1.5 (Mar 24 2014)

* Use real on-disk test data dumped from status.github.com instead of minimal
  fixtures in code.
* Run tests against Go 1.2 + tip on Travis.
* Add useful scripts under `script/` and remove Makefile.
* Overhaul README.

## v1.4 (Aug 26 2013)

* Use testify asserts in tests. Add "Testing" section to README.
* Makefile: make -x optional to reduce verbosity

## v1.3 (Aug 21 2013)

* Remove superfluous "Status" prefix from status names (API-breaking change).
* Add `serveTestResponses` to make test code more readable.

## v1.2 (Aug 10 2013)

* Add functions `ServiceURL` and `SetServiceURL` to get and set service URL
  (API-breaking change).
* Remove empty line before `package` clause to fix Godoc.

## v1.1 (Aug 6 2013)

* Store timestamps in `Status` and `Message` structs as `time.Time`
  (API-breaking change).
* Update Godoc.

## v1.0 (Aug 1 2013)

* Initial public release.
