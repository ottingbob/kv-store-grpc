### Instructions: 
Build a simple decomposed Key-Value store by implementing two services which communicate over gRPC.

The first service should implement a basic JSON Rest API to serve as the primary public interface.
This service should then externally communicate with a second service over gRPC, which implement a basic Key-Value store service that can:
- Store a value at a given key
- Retrieve the value for a given key
- Delete a given key

The JSON interface should at a minimum be able to expose and implement these three functions.

You can write this in whichever languages you choose, however Go would be preferred.
The final result should be built into two separate Docker containers which can be used to run each service independently.
