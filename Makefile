
VERSION ?= unknown
NOTIFIER_PATH := $(GOPATH)/src/github.com/diskordanz/weather_notifier/cmd/notifier

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
	ENABLE_WEATHER_PUSH_REQUESTS=true \
	API_KEY=b04ad8db6f75cbd1a02e6e4c8e1e1272 \
	./main
 
cli:
	API_KEY=b04ad8db6f75cbd1a02e6e4c8e1e1272 \
	go run client/client.go

image: 
	build
	CGO_ENABLED=0 go build -a -ldflags '-s -X main.serviceVersion=$(VERSION)' -installsuffix cgo -o $(NOTIFIER_PATH)/notifier $(NOTIFIER_PATH)/
	docker build -t notifier $(NOTIFIER_PATH)
	docker build -t darksky .
	rm $(NOTIFIER_PATH)/notifier
	rm darksky

up:
	docker-compose up

down:
	docker-compose down

