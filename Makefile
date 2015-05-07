default: build

clean:
	rm -f cmd/server/server
	rm -f cmd/latency-cli/latency-cli

build:
	cd cmd/server; \
	go get
	cd cmd/server; \
	go build
	cd cmd/latency-cli; \
	go build

deps:
	cd cmd/server; \
	go get

test:
	./cmd/server/server --config config.json server & \
	pid=$$!; \
	go test; \
	kill $$pid
