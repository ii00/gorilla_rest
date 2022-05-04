package service

import (
	"errors"
	"fmt"
	"grpc_app/pb"
	"sync"

	"github.com/jinzhu/copier"
)

// ErrAlreadyExist is returned when a record with the same ID already exists in the store.
var ErrAlreadyExist = errors.New("record already exist")

// LaptopStore is an interface to store laptop.
type LaptopStore interface {
	// Save saves the laptop to the store
	Save(laptop *pb.Laptop) error
}

// InMemoryLaptopStore stores laptop in memory.
type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

// DBLaptopStore stores laptop in DB.
// THIS ONE FOR LATER!
type DBLaptopStore struct{}

// NewInMemoryLaptopStore returns a new InMemoryLaptopStore.
func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

// Save saves the laptop to the store
func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExist
	}

	// deep copy
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data: %w", err)
	}

	store.data[other.Id] = other
	return nil
}
