bootstrap:
	go get -v -t ./...
	go get -v golang.org/x/tools/cmd/cover

build:
	gox \
		-output="build/ghstatus_{{.OS}}_{{.Arch}}" \
		-os="darwin linux windows" \
		./...

test:
	go test -v -cover ./...
