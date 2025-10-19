# Server gRPC streaming
## Purpose
In this example, we want to extend the [Unary gRPC](https://github.com/RaphSku/RoadToDistributedSystems/tree/main/networking/gRPC/unary_rpc)
scenario with server gRPC streaming. So to get the full context, have a look at the unary gRPC scenario.

## How it works
We have included one additional service method called `ListStorageLayoutOfCenter` to demonstrate the server streaming RPCs. This method retrieves and streams the availability status of storage blocks within a specific storage center. It looks up a given storage center ID in a map (StorageCenterManager) that contains information about available storage space, then sends structured responses for each storage layer and block to a client via a stream.

The difference to the unary gRPC scenario is that when the client invokes the `ListStorageLayoutOfCenter` method, the server does not send back a single response but sends back a stream of multiple responses over time that the client receives until the stream ends.

You can see this by having a look at the client code, there you can find the following code:
```go
resp, err := storageList.Recv()
```
where `storageList` is the streaming client that receives responses until either an `io.EOF` is thrown which indicates that the client has read all responses from the server and this is expected behavior or until any other error is returned which indicates that something went wrong.

## How to run
Just run
```
make start_server
```
first to start the server and then
```
make start_client
```
to start the client. Play around with this example, to get a feeling and since we only run one service method in the client,
try to run a sequence of those two service methods. You can also use `grpcurl` to interact with the gRPC server without
using the client.
