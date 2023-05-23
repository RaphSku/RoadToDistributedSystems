package main

import (
	"context"
	"net"
	"os"

	"github.com/RaphSku/RoadToDistributedSystems/tree/main/networking/gRPC/unary_rpc/server/itemstore"
	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type StorageLayout [][]int

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

type Item struct {
	Id            int32  `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	StorageCenter string `json:"storageCenter"`
	StorageLayer  int32  `json:"storageLayout"`
	StorageBlock  int32  `json:"storageBlock"`
}

type Storage map[int32]map[int32][]Item

var StorageDelivery = map[string]Storage{
	"K":   {},
	"FRA": {},
	"HH":  {},
}

type ItemStore struct {
	logger hclog.Logger
	itemstore.UnimplementedItemStoreServiceServer
}

func NewItemStore(logger hclog.Logger) *ItemStore {
	return &ItemStore{logger, itemstore.UnimplementedItemStoreServiceServer{}}
}

func (itemStore *ItemStore) CheckStorageAvailability(ctx context.Context, req *itemstore.ItemStoreAvailabilityRequest) (*itemstore.ItemStoreAvailabilityResponse, error) {
	itemStore.logger.Info("Handle CheckStorageAvailability")

	if (req.StorageLayer < 1) || (req.StorageBlock < 1) {
		return nil, status.Errorf(codes.InvalidArgument, "StorageLayer or StorageBlock cannot have a value less than 1")
	}

	storageCenterId := req.StorageCenter.String()
	storageLayerId := req.StorageLayer - 1
	storageBlockId := req.StorageBlock - 1

	return &itemstore.ItemStoreAvailabilityResponse{
		Availability: int32(StorageCenterManager[storageCenterId][storageLayerId][storageBlockId]),
		StorageLayer: storageLayerId,
		StorageBlock: storageBlockId,
	}, nil
}

func (itemStore *ItemStore) DeliverItemToStorage(ctx context.Context, req *itemstore.ItemStoreRequest) (*itemstore.ItemStoreResponse, error) {
	itemStore.logger.Info("Delivering Item to Storage Location")

	if (req.StorageLayer < 1) || (req.StorageBlock < 1) {
		return nil, status.Errorf(codes.InvalidArgument, "StorageLayer or StorageBlock cannot have a value less than 1")
	}

	storageCenterId := req.StorageCenter.String()
	storageLayoutId := req.StorageLayer - 1
	storageBlockId := req.StorageBlock - 1

	storage, ok := StorageDelivery[storageCenterId]
	if ok {
		if StorageCenterManager[storageCenterId][storageLayoutId][storageBlockId] < 1 {
			return nil, status.Errorf(codes.OutOfRange, "Unfortunately this StorageLayout and StorageBlock have no availability left")
		}

		new_item := Item{
			Id:            req.Id,
			Name:          req.Name,
			Description:   req.Description,
			StorageCenter: req.StorageCenter.String(),
			StorageLayer:  req.StorageLayer,
			StorageBlock:  req.StorageBlock,
		}

		if storage[storageLayoutId] == nil {
			storage[storageLayoutId] = make(map[int32][]Item)
		}

		storage[storageLayoutId][storageBlockId] = append(storage[storageLayoutId][storageBlockId], new_item)

		StorageCenterManager[storageCenterId][storageLayoutId][storageBlockId] -= 1
	}

	return &itemstore.ItemStoreResponse{
		Id: req.Id,
	}, nil
}

func main() {
	logger := hclog.Default()

	gs := grpc.NewServer()
	is := NewItemStore(logger)

	itemstore.RegisterItemStoreServiceServer(gs, is)

	reflection.Register(gs)

	listener, err := net.Listen("tcp", "localhost:9099")
	if err != nil {
		logger.Error("Service cannot start...due to:", err)
		os.Exit(1)
	}

	gs.Serve(listener)
}
