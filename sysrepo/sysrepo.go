package sysrepo

/*
#cgo pkg-config: sysrepo
#include <sysrepo.h>
#include "helper.h"
*/
import "C"
import (
	"fmt"
)

type Connection struct {
	connection *C.sr_conn_ctx_t
	connOpts   C.sr_conn_options_t
}

func Connect() (Connection, error) {
	var conn = Connection{connection: nil, connOpts: C.SR_CONN_DEFAULT}
	var rc C.int = C.SR_ERR_OK

	/*	moduleName := C.CString("ietf-interfaces")
		defer C.free(unsafe.Pointer(moduleName))
	*/
	rc = C.sr_connect(conn.connOpts, &conn.connection)
	if C.SR_ERR_OK != rc {
		return conn, fmt.Errorf("error by sr_connect: %v", C.sr_strerror(rc))
	}

	return conn, nil
}

func Disconnect(conn Connection) error {
	var rc C.int = C.SR_ERR_OK
	rc = C.sr_disconnect(conn.connection)
	if C.SR_ERR_OK != rc {
		return fmt.Errorf("error by sr_connect: %v", C.sr_strerror(rc))
	}
	return nil
}
