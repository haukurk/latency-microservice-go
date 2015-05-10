default: build

clean:
	rm -f cmd/latency-server/latency-server
	rm -f cmd/latency-cli/latency-cli

build:
	cd cmd/latency-server; \
	go build
	cd cmd/latency-cli; \
	go build

deps:
	cd cmd/latency-server; \
	go get

test:
	./cmd/latency-server/latency-server --config config.json server & \
	pid=$$!; \
	go test; \
	kill $$pid

install:
	cd cmd/latency-server; \
	go install
