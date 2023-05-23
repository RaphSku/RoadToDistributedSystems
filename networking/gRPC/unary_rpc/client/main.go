package main

import (
	"context"
	"os"

	"github.com/RaphSku/RoadToDistributedSystems/tree/main/networking/gRPC/unary_rpc/client/itemstore"
	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

type ItemStore struct {
	logger hclog.Logger
	ics    itemstore.ItemStoreServiceClient
}

func NewItemStore(logger hclog.Logger, ics itemstore.ItemStoreServiceClient) *ItemStore {
	return &ItemStore{logger, ics}
}

func main() {
	logger := hclog.Default()

	conn, err := grpc.Dial("localhost:9099", grpc.WithInsecure())
	if err != nil {
		logger.Error("Cannot reach gRPC server...due to:", err)
		os.Exit(1)
	}
	defer conn.Close()

	isc := itemstore.NewItemStoreServiceClient(conn)

	is := NewItemStore(logger, isc)

	isa, err := is.ics.CheckStorageAvailability(context.Background(), &itemstore.ItemStoreAvailabilityRequest{
		StorageCenter: 0,
		StorageLayer:  1,
		StorageBlock:  1,
	})
	if err != nil {
		logger.Error("Request could not be processed by gRPC server...due to:", err)
		os.Exit(1)
	}

	logger.Info(isa.String())
}
