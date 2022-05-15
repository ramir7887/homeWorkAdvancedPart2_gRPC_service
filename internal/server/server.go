package server

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"homeWorkAdvancedPart2_gRPC_serv/internal/model"
	"homeWorkAdvancedPart2_gRPC_serv/internal/pb"
	"homeWorkAdvancedPart2_gRPC_serv/internal/store"
	"log"
	"math/rand"
	"time"
)

type Server struct {
	s store.Store
	pb.UnimplementedDeliveryServiceServer

}

func NewServer(s store.Store) *Server {
	return &Server{
		s: s,
	}
}

func (s *Server) CreateDelivery(ctx context.Context, in *pb.Order) (*pb.DeliveryID, error) {
	log.Println("CreateDelivery")
	rand.Seed(time.Now().Unix())
	randDay := time.Duration(rand.Int63n(15))
	deliveryDate := time.Now().Add(randDay * 24 * time.Hour)

	delivery := &model.Delivery{
		OrderID: int(in.Id),
		DeliveryDate: deliveryDate,
		Complete: false,
		Address: in.Address,
	}

	id, err := s.s.Delivery().Create(delivery)
	if err != nil {
		return nil, err
	}

	return &pb.DeliveryID{Value: int32(id)}, nil
}

func (s *Server) GetDelivery(ctx context.Context, in *pb.DeliveryID) (*pb.Delivery, error) {
	log.Println("GetDelivery")
	delivery, err := s.s.Delivery().Find(int(in.Value))
	if err != nil {
		return nil, err
	}
	return &pb.Delivery{
		Id: int32(delivery.ID),
		DeliveryDate: timestamppb.New(delivery.DeliveryDate),
		Complete: delivery.Complete,
		Address: delivery.Address,
	}, nil
}

func (s *Server) DeleteDelivery(ctx context.Context, in *pb.DeliveryID) (*pb.Complete, error) {
	log.Println("DeleteDelivery")
	err := s.s.Delivery().Cancel(int(in.Value))
	if err != nil {
		return nil, err
	}
	return &pb.Complete{Value: true}, err
}

func (s *Server) CompleteDelivery(ctx context.Context, in *pb.DeliveryID) (*pb.Complete, error) {
	log.Println("CompleteDelivery")
	err := s.s.Delivery().Complete(int(in.Value))
	if err != nil {
		return nil, err
	}
	return &pb.Complete{Value: true}, err
}



