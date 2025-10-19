# Unary gRPC
## Purpose
In this example, we want to demonstrate how unary gRPC works which closely resembles normal functions since
the client sends a single request and receives a single response back.

## Scenario
Suppose you got a regional storage system where each region has a number of storage layers and storage blocks.
Each storage block can hold a different number of items before its capacity is depleted. And you want to be
able to assign different items to different storage locations. Note, since this is an educational example that
we did not cover all edge cases possible. But it might be instructive to think yourself about all the different
kind of edge cases that could appear in such a system.

## How it works
Let's begin with the server and its design. The server has a hard-coded storage center layout
```go
var StorageCenterManager = map[string]StorageLayout{
	"K": {
		{20, 10, 5},
		{5, 15, 10},
	},
	"FRA": {
		{5, 5, 5},
		{5, 15, 5},
	},
	"HH": {
		{15, 5},
		{10, 15},
	},
}
```
with three regions: K stands for "Koeln", FRA stands for "Frankfurt" and HH stands for "Hamburg",
each having two storage layers which are the rows and storage blocks which are the respective column elements.
When we store an item, we need to say where we would like to store that item, so we need to provide three
information: the region, the storage layer and the storage block.

Additionally, an item is defined as
```go
type Item struct {
	Id            int32  `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	StorageCenter string `json:"storageCenter"`
	StorageLayer  int32  `json:"storageLayout"`
	StorageBlock  int32  `json:"storageBlock"`
}
```
where we can assign each item an id, a name and a description in addition to the locators.

Furthermore, our server offers two service methods named `DeliverItemToStorage` and `CheckStorageAvailability`.
Let's talk first about the `CheckStorageAvailability` service method which is very simple to understand.
It accepts as arguments the three locators where the storage layer and storage block are one-indexed and that is it.
It then maps the one-indexed locators to a zero-indexed value and uses them to retrieve the current availability
of that storage space. In the client, we run this service method and ask for the availability space
for Koeln in storage layer 1 and storage block 1. Now, since the server maps the storage layer and storage block back
to a zero-indexed number, we ask for Koeln and (0,0) which will return 20 if you have a look at our hard-coded storage layout.

The `DeliverItemToStorage` service method is also quite simple, we simply pass as arguments all the attributes that we need
to initialize an item and then check whether we have enough availability left, if not, then the request cannot be fulfilled.
If there is enough space, the availability is decreased and the item is stored in another data structure
```go
var StorageDelivery = map[string]Storage{
	"K":   {},
	"FRA": {},
	"HH":  {},
}
```
where `Storage` is defined as
```go
type Storage map[int32]map[int32][]Item
```
where those two `int32` correspond to the storage layer and storage block.

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
