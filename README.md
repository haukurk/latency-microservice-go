[![Build Status](https://travis-ci.org/haukurk/latency-microservice-go.svg?branch=master)](https://travis-ci.org/haukurk/latency-microservice-go)
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


# JSON specifications (Examples)

/GET /latency/<string:hostname>

200 OK
```
{
  "ip": "74.125.136.138",
  "rtt": 4.243,
  "status": "ok",
  "unit": "ms"
}

```

404 NOT FOUND
```
{
  "error":"cannot resolve remote address",
  "status":"fail"
}
```

/GET /stats

200 OK
```
{
  "pid":1953,
  "uptime":"1h42m44.212770594s",
  "uptime_sec":6164.212770594,
  "time":"2015-05-10 10:33:17.482820233 -0400 EDT",
  "unixtime":1431268397,
  "status_code_count":{},
  "total_status_code_count":{"200":4},
  "count":0,
  "total_count":4,
  "total_response_time":"4.199562871s",
  "total_response_time_sec":4.199562871,
  "average_response_time":"1.049890717s",
  "average_response_time_sec":1.049890717
}
```

# Considerations

The ICMP library uses raw sockets, therefore needs root privileges to function properly.

*I discourage users using this as a public service.*

Make sure to run tests before using.
