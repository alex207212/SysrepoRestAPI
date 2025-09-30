package main

import (
	"SysrepoRestAPI/sysrepo"
	"fmt"
	"log"
)

//import "sysrepo"

func main() {
	conn, err := sysrepo.Connect()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to sysrepo")
	err = sysrepo.Disconnect(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Disconnected from sysrepo")
}
