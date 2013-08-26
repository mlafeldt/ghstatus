v1.4 (Aug 26 2013)
------------------

* Use testify asserts in tests. Add "Testing" section to README.
* Makefile: make -x optional to reduce verbosity

v1.3 (Aug 21 2013)
------------------

* Remove superfluous "Status" prefix from status names (API-breaking change).
* Add `serveTestResponses` to make test code more readable.

v1.2 (Aug 10 2013)
------------------

* Add functions `ServiceURL` and `SetServiceURL` to get and set service URL
  (API-breaking change).
* Remove empty line before `package` clause to fix Godoc.

v1.1 (Aug 6 2013)
-----------------

* Store timestamps in `Status` and `Message` structs as `time.Time`
  (API-breaking change).
* Update Godoc.

v1.0 (Aug 1 2013)
-----------------

* Initial public release.
