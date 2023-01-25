# Request Forwarder
- Forwards request to a list of defined domains
- If one of the domains return a failure, it goes to other domains until it fulfills all the request.

# Build and Install
```
go build -o req-forwarder
```

# Running
```
./req-forwarder --domains=https://shuttle-4-bs1.estuary.tech,https://shuttle-4-bs1.estuary.tech
```

# Test
Run the following
https://shuttle-4.estuary.tech/gw/<>
