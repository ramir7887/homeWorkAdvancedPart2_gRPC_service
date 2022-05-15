package store


type Store interface {
	Delivery() DeliveryRepository
}