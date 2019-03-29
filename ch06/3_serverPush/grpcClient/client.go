package main

import (
	"context"
	pb "github.com/MartinToruan/building-restful-web-services-in-go/ch06/3_serverPush/datafiles"
	"google.golang.org/grpc"
	"io"
	"log"
)

const (
	address= "localhost:50051"
)

func main(){
	// Connect to server
	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can't connect to the gRPC server: %v\n", err)
	}
	defer cc.Close()

	// Create a gRPC Client
	c := pb.NewMoneyTransactionClient(cc)

	// Hit the server
	stream, err := c.MakeTransaction(context.Background(), &pb.TransactionRequest{
		From: "Martin",
		To: "Kristopel",
		Amount: 50000,
	})
	if err != nil{
		log.Fatalf("error while hit the api: %v", err)
	}

	for {
		resp, err := stream.Recv()

		if err == io.EOF{
			log.Println("Finish consume response from the server.")
			break
		}

		if err != nil{
			log.Fatalf("got error response from the server: %v\n", err)
		}

		log.Printf("step: %d status:%s description: %s\n", resp.GetStep(), resp.GetStatus(), resp.GetDescription())
	}
}