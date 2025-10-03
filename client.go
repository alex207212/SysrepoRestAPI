package main

import (
	pb "SysrepoRestAPI/demo_proto"
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect. Error: %v", err)
	}
	defer conn.Close()
	client := pb.NewSysrepoServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.GetItems(ctx, &pb.GetItemsRequest{Xpath: "ietf-netconf-acm"})
	if err != nil {
		log.Fatalf("could not get items. Error: %v", err)
	}
	fmt.Printf("GetItems: %v", r)
}
