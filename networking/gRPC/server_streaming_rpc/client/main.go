package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/RaphSku/RoadToDistributedSystems/tree/main/networking/gRPC/server_streaming_rpc/client/itemstore"
	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	conn, err := grpc.NewClient("localhost:9099", grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	storageList, err := is.ics.ListStorageLayoutOfCenter(context.Background(), &itemstore.StorageCenterLayoutRequest{
		StorageCenter: 0,
	})
	if err != nil {
		logger.Error("Request could not be processed by gRPC server...due to:", err)
		os.Exit(1)
	}

	done := make(chan bool)
	go func() {
		for {
			resp, err := storageList.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				logger.Error("Cannot receive message due to:", err)
				continue
			}
			message := fmt.Sprintf(
				"Message received and the result is:\n"+
					"Availability: %v\n"+
					"StorageLayer: %v\n"+
					"StorageBlock: %v\n",
				resp.Availability,
				resp.StorageLayer,
				resp.StorageBlock,
			)
			logger.Info(message)
		}
	}()
	<-done

	isresp, err := is.ics.DeliverItemToStorage(context.Background(), &itemstore.ItemStoreRequest{
		Id:            0,
		Name:          "ReTake Shoe Power",
		Description:   "Shoes that take your feets on the next level",
		StorageCenter: 1,
		StorageLayer:  2,
		StorageBlock:  3,
	})
	if err != nil {
		logger.Error("Request could not be processed by gRPC server...due to:", err)
		os.Exit(1)
	}
	logger.Info(fmt.Sprintf("Item with Id: %v was delivered successfully!", isresp.Id))

	isa, err = is.ics.CheckStorageAvailability(context.Background(), &itemstore.ItemStoreAvailabilityRequest{
		StorageCenter: 1,
		StorageLayer:  2,
		StorageBlock:  3,
	})
	if err != nil {
		logger.Error("Request could not be processed by gRPC server...due to:", err)
		os.Exit(1)
	}
	logger.Info(isa.String())
}
