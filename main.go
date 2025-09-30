package main

import (
	"SysrepoRestAPI/sysrepo"
	"fmt"
)

//import "sysrepo"

func main() {
	fmt.Println("main")
	sysrepo.Sysrepotest()
	sysrepo.Connect()
}
