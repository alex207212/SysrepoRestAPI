package main

import (
	pb "SysrepoRestAPI/demo_proto"
	"SysrepoRestAPI/sysrepo"
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedSysrepoServerServer
}

func (s *server) GetItems(_ context.Context, in *pb.GetItemsRequest) (*pb.ValuesList, error) {
	log.Printf("GetItems called with xpath: %s", in.Xpath)
	return &pb.ValuesList{Values: []string{"123", "456", "789", "012"}}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSysrepoServerServer(s, &server{})
	fmt.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main1() {
	// Open connection
	conn, err := sysrepo.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer disconnect(conn)
	fmt.Println("Connected to sysrepo")

	// Start session
	sess, err := sysrepo.StartSession(conn, sysrepo.DS_STARTUP)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer closeSession(sess)
	fmt.Println("Started session")

	/*	modules, err := sysrepo.GetDefaultModules()
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, module := range modules {
			fmt.Println(module)
		}
	*/
	fmt.Printf("Current log level: %v; setting to %v\n", sysrepo.GetLogLevel(), sysrepo.SR_LL_DBG)
	sysrepo.SetLogLevel(sysrepo.SR_LL_DBG)

	//"ietf-interfaces", "ietf-datastores"
	err = sysrepo.PrintCurrentConfig(sess, "ietf-netconf-acm")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func disconnect(conn sysrepo.Connection) {
	err := sysrepo.Disconnect(conn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Disconnected from sysrepo")
}

func closeSession(sess sysrepo.Session) {
	err := sysrepo.StopSession(sess)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Stopped session")
}
