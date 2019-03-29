package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	pb "github.com/MartinToruan/building-restful-web-services-in-go/ch06/2_grpc_example/datafiles"
	"net"
)

const (
	port = ":50051"
)

type server struct {

}

func (s *server) MakeTransaction(ctx context.Context, req *pb.TransactionRequest) (*pb.TransactionResponse, error){
	log.Println("Start MakeTransaction")
	defer log.Println("Finish MakeTransaction")
	log.Printf("Amount: %f, From A/c: %s, To A/c: %s\n", req.GetAmount(), req.GetFrom(), req.GetTo())

	return &pb.TransactionResponse{
		Confirmation: true,
	}, nil
}

func main(){
	lis, err := net.Listen("tcp", port)
	if err != nil{
		log.Fatalf("can't listen on the port %v, err: %v", port, err)
	}

	// Create gRPC server
	s := grpc.NewServer()
	pb.RegisterMoneyTransactionServer(s, &server{})

	// Add Reflection
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("error while start the server: %v\n", err)
	}
}
