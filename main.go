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

type server struct {
	pb.UnimplementedSysrepoServerServer
	connection *sysrepo.Connection
}

var (
	port     = flag.Int("port", 50051, "The server port")
	srServer *server
)

func (s *server) GetItems(_ context.Context, in *pb.GetItemsRequest) (*pb.ValuesList, error) {
	log.Printf("GetItems called with xpath: %s", in.Xpath)
	// Start session
	sess, err := sysrepo.StartSession(s.connection, sysrepo.DS_STARTUP)
	if err != nil {
		return nil, err
	}
	defer closeSession(sess)
	fmt.Println("Started session")
	r, err := sysrepo.GetValuesForModule(sess, in.Xpath)
	if err != nil {
		return nil, err
	}
	return &pb.ValuesList{Values: r}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srServer = &server{}
	// Open connection
	srServer.connection, err = sysrepo.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer disconnect()
	fmt.Println("Connected to sysrepo")
	fmt.Printf("Current log level: %v; setting to %v\n", sysrepo.GetLogLevel(), sysrepo.SR_LL_DBG)
	sysrepo.SetLogLevel(sysrepo.SR_LL_DBG)

	s := grpc.NewServer()
	pb.RegisterSysrepoServerServer(s, srServer)
	fmt.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main1() {

	/*	modules, err := sysrepo.GetDefaultModules()
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, module := range modules {
			fmt.Println(module)
		}
	*/

	//"ietf-interfaces", "ietf-datastores"
	/*	err := sysrepo.PrintCurrentConfig(sess, "ietf-netconf-acm")
		if err != nil {
			fmt.Println(err)
			return
		}
	*/
}

func disconnect() {
	err := sysrepo.Disconnect(srServer.connection)
	if err != nil {
		fmt.Println(err)
	}
	srServer.connection = nil
	fmt.Println("Disconnected from sysrepo")
}

func closeSession(sess *sysrepo.Session) {
	err := sysrepo.StopSession(sess)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Stopped session")
}
