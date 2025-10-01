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

func StartSession(conn Connection, ds Datastore) (Session, error) {
	var rc C.int = C.SR_ERR_OK
	var sess = Session{session: nil, datastore: ds.to_sr_datastore_t()}
	rc = C.sr_session_start(conn.connection, sess.datastore, &sess.session)
	if C.SR_ERR_OK != rc {
		return sess, fmt.Errorf("error by sr_session_start: %v", C.sr_strerror(rc))
	}
	return sess, nil
}

func StopSession(sess Session) error {
	var rc C.int = C.SR_ERR_OK
	rc = C.sr_session_stop(sess.session)
	if C.SR_ERR_OK != rc {
		return fmt.Errorf("error by sr_stop_session: %v", C.sr_strerror(rc))
	}
	return nil
}
