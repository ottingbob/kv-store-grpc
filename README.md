# Key Value Store over GRPC

This simple project aims to create a decomposed Key-Value store by implementing two services, Rest API and RPC, which communicate over gRPC.

The solution is written in golang and provides a solution using the `grpc-gateway` package in order to setup handlers to validate the JSON requests via API and convert them to proto messages in the RPC.

The Key-Value store service can:
- Store a value at a given key
- Retrieve the value for a given key
- Delete a given key

The application is setup using `docker-compose` which uses a single docker file and creates the specific application based on command line arguments passed into the built application.

### Running using docker-compose

The solution should work based on using the following command and expect and output where both services print that they are initialized:

```
$ docker-compose up --build
...
kv-rest | Key-Value REST Service Initializing...
kv-rpc  | Key-Value RPC Service Initializing...
```

### Testing out the application

The application takes JSON requests on a local endpoint exposed over `localhost:8080`. The `Makefile` provides a way to quickly test out the three methods by using the `make test` command:

```
$ make test
curl -X POST localhost:8080/v1/kv -d '{"key": "nice", "value": "gotem" }' && echo
{"key":"nice", "value":"gotem"}
curl -X GET localhost:8080/v1/kv/nice && echo
{"key":"nice", "value":"gotem"}
curl -X DELETE localhost:8080/v1/kv/nice && echo
{}
# I should expect an error here...
curl -X GET localhost:8080/v1/kv/nice && echo
{"code":5, "message":"Key-Value pair could not be found at key nice", "details":[]}
# I should expect an error here...
curl -X DELETE localhost:8080/v1/kv/nice && echo
{"code":5, "message":"Key-Value pair could not be found at key nice", "details":[]}
```

Additional `curl` requests may also be used to add key/value pairs under new keys and the proto validation should be able to eliminate invalid data types when hitting the various endpoints.