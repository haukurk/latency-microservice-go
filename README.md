LatencyREST 
=================

An extremly simple service that allows you to check latency between the server and some IP.
The project is based on fastping-go and the GIN framework to simplify ICMP and service communication.


# Endpoint specifications

/GET /ping/<string:hostname>
Returns:

200 OK
```
{
  "ip": "74.125.136.138",
  "rtt": 4,
  "status": "ok",
  "unit":"ms"
}
```
