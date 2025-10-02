package main

import (
	"SysrepoRestAPI/sysrepo"
	"fmt"
)

//import "sysrepo"

func main() {
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
