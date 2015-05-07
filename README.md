Latency Analyser Microservice
=================

An extremly simple service that allows you to check latency between the server and some IP.
The project is based on fastping-go and the GIN framework to simplify ICMP and service communication.

Frameworks and libraries used:
 * Gin
 * cli.go

# Installation

Building solution:
```
make deps && make 
```

Run tests
```
make test
```

# Running the service

To start-up the service use:
```
cmd/server/server --config config.json
```

# Client library

The solution includes a client library to interact with the service.

```
cmd/latency-cli/latency-cli --host $HOSTNAME -r $HOSTTOCHECKLATENCY
```


# Endpoint specifications

/GET /latency/<string:hostname>
Returns:

200 OK
```
{
  "ip": "74.125.136.138",
  "rtt": 4.243ms,
  "status": "ok",
}

```

# Considerations

The ICMP library uses raw sockets, therefore needs root privileges to function properly.

*I discourage users using this as a public service.*

Make sure to run tests before using.
