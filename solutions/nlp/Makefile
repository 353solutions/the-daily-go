all:
	$(error please pick a target)

test:
	go test -v ./...


nlpd: cmd/nlpd
	go build \
	    -ldflags="-X main.Version=$(shell git rev-parse --short HEAD)" \
	    ./cmd/nlpd

stop_words:
	go generate


grpc:
	mkdir -p pb
	protoc nlp.proto --go_out=plugins=grpc:pb

circleci:
	docker build -f Dockerfile.test .
