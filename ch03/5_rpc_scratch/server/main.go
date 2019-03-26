package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct {

}

type TimeServer int64

func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error{
	// Fill reply pointer to send the data back
	*reply = time.Now().Unix()
	return nil
}

func main(){
	timeServer := new(TimeServer)
	if err := rpc.Register(timeServer); err != nil{
		log.Fatal("Can't register rpc")
	}
	rpc.HandleHTTP()

	// Start TCP Server
	l, err := net.Listen("tcp", ":1234")
	if err != nil{
		log.Fatal("Can't start TCP Server")
	}

	log.Fatal(http.Serve(l, nil))
}
