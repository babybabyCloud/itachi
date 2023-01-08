.PHONY: build test update_dep cover clean

SHELL := /bin/bash
binary := itachi

build: update_dep
	go build -o ${binary}

test: update_dep
	if [ "" == "${coverfile}" ]; then \
		go test -v ./...; \
	else \
		go test -v -coverprofile=${coverfile}  ./...; \
	fi

update_dep:
	go mod tidy

cover: export coverfile := cover.out
cover: test
	go tool cover -html=${coverfile}

clean:
	rm ${binary} ${coverfile}
