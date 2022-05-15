package inmemorystore

import (
	"homeWorkAdvancedPart2_gRPC_serv/internal/model"
	"log"
)

var (
	rnfError = RecordNotFoundError{}
)

type RecordNotFoundError struct {
}

func (r RecordNotFoundError) Error() string {
	return "record not found"
}

type Repository struct {
	store *Store
}

func (r *Repository) Create(delivery *model.Delivery) (int, error) {
	id := r.nextID()
	delivery.ID = id
	r.store.store = append(r.store.store, *delivery)

	log.Printf("Records: %v", r.store.store)

	return id, nil
}

func (r *Repository) Find(id int) (*model.Delivery, error) {
	for _, val := range r.store.store {
		if val.ID == id {
			delivery := val
			return &delivery, nil
		}
	}
	return nil, rnfError
}

func (r *Repository) Cancel(id int) error {
	index := -1
	for idx, val := range r.store.store {
		if val.ID == id {
			index = idx
		}
	}
	if index == -1 {
		return rnfError
	}
	r.store.store = append(r.store.store[:index], r.store.store[index + 1:]...)
	log.Printf("Record count: %v", len(r.store.store))
	return nil
}

func (r *Repository) Complete(id int) error {
	for _, val := range r.store.store {
		if val.ID == id {
			val.Complete = true
			return nil
		}
	}
	return rnfError
}

func (r Repository) lastID() int {
	id := 0
	for _, val := range r.store.store {
		if val.ID > id {
			id = val.ID
		}
	}
	return id
}

func (r Repository) nextID() int {
	return r.lastID() + 1
}
