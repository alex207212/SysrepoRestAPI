package main

import (
	"SysrepoRestAPI/sysrepo"
	"fmt"
	"log"
)

//import "sysrepo"

func main() {
	// Open connection
	conn, err := sysrepo.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer disconnect(conn)
	fmt.Println("Connected to sysrepo")

	// Start session
	sess, err := sysrepo.StartSession(conn, sysrepo.DS_STARTUP)
	if err != nil {
		log.Fatal(err)
	}
	defer closeSession(sess)
	fmt.Println("Started session")
}

func disconnect(conn sysrepo.Connection) {
	err := sysrepo.Disconnect(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Disconnected from sysrepo")
}

func closeSession(sess sysrepo.Session) {
	err := sysrepo.StopSession(sess)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Stopped session")
}
