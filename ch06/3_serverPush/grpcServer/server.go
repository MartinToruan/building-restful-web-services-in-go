package main

import (
	"fmt"
	pb "github.com/MartinToruan/building-restful-web-services-in-go/ch06/3_serverPush/datafiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"time"
)

const(
	port= ":50051"
)

type server struct {

}

var steps = map[int32]string{
	0: "initial",
	1: "approved",
	2: "queued",
	3: "completed",
}

func (s *server) MakeTransaction(req *pb.TransactionRequest, stream pb.MoneyTransaction_MakeTransactionServer) error{
	log.Println("start MakeTransaction...")
	defer log.Println("finish MakeTransaction...")

	for k, v := range steps{
		if err := stream.Send(&pb.TransactionResponse{
			Status: v,
			Step: k,
			Description: fmt.Sprintf("Transaction %s, Status: %s", req.GetTo(), v),
		}); err != nil{
			log.Fatalf("error while send responses to client: %v\n", err)
			return status.Errorf(codes.Internal,
				fmt.Sprintf("Internal error: %v\n", err))
		}
		time.Sleep(1 * time.Second)
	}

	return nil
}

func main(){
	// Create Listener
	l, err := net.Listen("tcp", port)
	if err != nil{
		log.Fatalf("error while connect to addres %v. err: %v\n", port, err)
	}

	// Create gRPC server
	s := grpc.NewServer()
	pb.RegisterMoneyTransactionServer(s, &server{})

	// Start the server
	if err := s.Serve(l); err != nil{
		log.Fatalf("error while start the server: %v\n", err)
	}
}