.PHONY: all
all: gen

.PHONY: gen
gen:
	go generate ./...

.PHONY: dev-deps
dev-deps:
	go get github.com/davecheney/godoc2md
