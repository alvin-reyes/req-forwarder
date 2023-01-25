# Request Forwarder

![image](https://user-images.githubusercontent.com/4479171/214675514-fa7a4a69-fb34-49a9-9089-325aebf38b61.png)


- Forwards request to a list of defined domains
- If one of the domains return a failure, it goes to other domains until it fulfills all the request.

# Build and Install
```
go build -tags netgo -ldflags '-s -w' -o req-forwarder
```

# Running
```
./req-forwarder --domains=https://shuttle-4-bs1.estuary.tech,https://shuttle-4-bs2.estuary.tech
```

# Test
Run the following
https://shuttle-4.estuary.tech/gw/<>
