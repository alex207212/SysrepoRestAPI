package sysrepo

/*
#cgo pkg-config: sysrepo
#include <sysrepo.h>
#include "helper.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func Sysrepotest() {
	fmt.Println("sysrepotest")
}

func Connect() {
	var connection *C.sr_conn_ctx_t = nil
	var rc C.int = C.SR_ERR_OK
	var connOpts C.sr_conn_options_t = C.SR_CONN_DEFAULT

	moduleName := C.CString("ietf-interfaces")
	defer C.free(unsafe.Pointer(moduleName))

	rc = C.sr_connect(connOpts, &connection)
	if C.SR_ERR_OK != rc {
		fmt.Printf("Error by sr_connect: %s\n", C.sr_strerror(rc))
		return
	} else {
		defer C.sr_disconnect(connection)
	}

	fmt.Println("Successfully connected to sysrepo")

}
