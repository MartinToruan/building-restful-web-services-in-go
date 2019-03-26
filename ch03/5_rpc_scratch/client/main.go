package main

import (
	"log"
	"net/rpc"
)

type Args struct {

}

func main(){
	var args Args
	var reply int64
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil{
		log.Fatalf("Can't connect to RPC server: %v", err)
	}
	if err := client.Call("TimeServer.GiveServerTime", args, &reply); err != nil{
		log.Fatalf("Got error response from server: %v", err )
	}
	log.Printf("Server time: %v", reply)
}
