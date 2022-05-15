package inmemorystore

import (
	"homeWorkAdvancedPart2_gRPC_serv/internal/model"
	"homeWorkAdvancedPart2_gRPC_serv/internal/store"
	"time"
)

type Store struct {
	store []model.Delivery
	repository *Repository
}

func NewStore() *Store {
	delivers := []model.Delivery{
		{ID: 1, OrderID: 1, DeliveryDate: time.Now().Add(time.Hour * 24 * 1), Complete: false, Address: "SomeAddress"},
		{ID: 2, OrderID: 2, DeliveryDate: time.Now().Add(time.Hour * 24 * 7), Complete: false, Address: "SomeAddress"},
		{ID: 3, OrderID: 3, DeliveryDate: time.Now().Add(time.Hour * 24 * 4), Complete: false, Address: "SomeAddress"},
		{ID: 4, OrderID: 4, DeliveryDate: time.Now().Add(time.Hour * 24 * 4), Complete: false, Address: "SomeAddress"},
		{ID: 5, OrderID: 5, DeliveryDate: time.Now().Add(time.Hour * 24 * 5), Complete: false, Address: "SomeAddress"},
		{ID: 6, OrderID: 6, DeliveryDate: time.Now().Add(time.Hour * 24 * 3), Complete: false, Address: "SomeAddress"},
		{ID: 8, OrderID: 7, DeliveryDate: time.Now(), Complete: true},
	}
	return &Store{
		store: delivers,
	}
}

func (s *Store) Delivery() store.DeliveryRepository {
	if s.repository != nil {
		return s.repository
	}

	s.repository = &Repository{
		store: s,
	}
	return s.repository
}