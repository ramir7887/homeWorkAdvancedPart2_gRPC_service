package store

import "homeWorkAdvancedPart2_gRPC_serv/internal/model"

type DeliveryRepository interface {
	Create(delivery *model.Delivery) (int, error)
	Find(id int) (*model.Delivery, error)
	Cancel(id int) error
	Complete(id int) error
}

