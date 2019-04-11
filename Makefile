VERSION ?= unknown

all: deps test build

fmt:
	goimports -w $$(ls -d */ | grep -v vendor)
	goimports -w main.go

test:
	go test -v --cover --race -short `glide novendor | grep -v ./proto`

build:
	CGO_ENABLED=0 go build -a -ldflags '-s -X main.serviceVersion=$(VERSION)' -installsuffix cgo -o main .

deps:
	glide install

run:
	./main
