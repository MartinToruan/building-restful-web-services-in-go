package main

import (
	"context"
	pb "github.com/MartinToruan/building-restful-web-services-in-go/ch06/2_grpc_example/datafiles"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:50051"
)

func main(){
	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil{
		log.Fatalf("can't connect to gRPC server: %v\n", err)
	}
	defer func(){
		if err := cc.Close(); err != nil{
			log.Fatalf("can't teardown the client: %v\n", err)
		}
	}()

	c := pb.NewMoneyTransactionClient(cc)

	resp, err := c.MakeTransaction(context.Background(), &pb.TransactionRequest{
		From: "Martin",
		To: "Kristopel",
		Amount: float32(15000),
	})
	if err != nil{
		log.Fatalf("got an error response from the server: %v\n", err)
	}

	log.Printf("Response confirmation: %v\n", resp.GetConfirmation())
}
