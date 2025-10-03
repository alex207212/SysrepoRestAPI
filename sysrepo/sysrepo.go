package sysrepo

/*
#cgo pkg-config: sysrepo
#include <sysrepo.h>
#include <sysrepo/values.h>
#include "helper.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func Connect() (*Connection, error) {
	var conn = Connection{connection: nil, connOpts: C.SR_CONN_DEFAULT}
	var rc C.int = C.SR_ERR_OK

	var connCtx *C.sr_conn_ctx_t
	rc = C.sr_connect(conn.connOpts, &connCtx)
	if err := CheckForError(rc); err != nil {
		return nil, err
	}
	conn.connection = connCtx

	return &conn, nil
}

func Disconnect(conn *Connection) error {
	var rc C.int = C.SR_ERR_OK
	rc = C.sr_disconnect(conn.connection)
	if C.SR_ERR_OK != rc {
		return fmt.Errorf("error by sr_connect: %v", C.sr_strerror(rc))
	}
	return nil
}

func StartSession(conn *Connection, ds Datastore) (*Session, error) {
	var rc C.int = C.SR_ERR_OK
	var sess = Session{session: nil, datastore: ds.to_sr_datastore_t()}
	rc = C.sr_session_start(conn.connection, sess.datastore, &sess.session)
	if C.SR_ERR_OK != rc {
		return nil, fmt.Errorf("error by sr_session_start: %v", C.sr_strerror(rc))
	}
	return &sess, nil
}

func StopSession(sess *Session) error {
	var rc C.int = C.SR_ERR_OK
	rc = C.sr_session_stop(sess.session)
	if C.SR_ERR_OK != rc {
		return fmt.Errorf("error by sr_stop_session: %v", C.sr_strerror(rc))
	}
	return nil
}

func GetDefaultModules() ([]string, error) {
	modules_ptr := C.sr_get_module_ds_default()
	if modules_ptr == nil {
		return nil, fmt.Errorf("error by sr_get_module_ds_default")
	}
	modules := [6]string{}
	for idx := 0; idx < 6; idx++ {
		modules[idx] = C.GoString((*modules_ptr).plugin_name[idx])
	}
	return modules[:], nil
}

func GetValuesForModule(sess *Session, module string) ([]string, error) {
	xpath := C.CString(fmt.Sprintf("/%s:*//*", module))
	defer C.free(unsafe.Pointer(xpath))

	var values *C.sr_val_t = nil
	var count C.size_t = 0
	var rc C.int = C.SR_ERR_OK

	rc = C.sr_get_items(sess.session, xpath, 0, 0, &values, &count)
	if err := CheckForError(rc); err != nil {
		return nil, err
	}
	defer C.sr_free_values(values, count)

	result := make([]string, count)
	var i C.size_t = 0
	for i = 0; i < count; i++ {
		val := C.get_val(values, i)
		result[i] = sysrepo_get_value(val)
	}

	return result, nil
}

func sysrepo_get_value(value *C.sr_val_t) string {
	var mem *C.char = nil
	rc := C.sr_print_val_mem(&mem, value)
	if C.SR_ERR_OK != rc {
		return fmt.Sprintf("Error by sr_print_val_mem: %d", C.sr_strerror(rc))
	}
	return fmt.Sprintf("%s", C.GoString(mem))
}

/*
func print_value(value *C.sr_val_t) {
	fmt.Printf("%s ", C.GoString(value.xpath))

	switch value._type {
	case C.SR_CONTAINER_T:
	case C.SR_CONTAINER_PRESENCE_T:
		fmt.Printf("(container)\n")
		break
	case C.SR_LIST_T:
		fmt.Printf("(list instance)\n")
		break
	case C.SR_STRING_T:
		val := (**C.char)(unsafe.Pointer(&value.data))
		fmt.Printf("= %s\n", C.GoString(*val))
		break
	case C.SR_BOOL_T:
		bool_val := (*C.bool)(unsafe.Pointer(&value.data))
		if *bool_val == C.bool(true) {
			fmt.Printf("= true\n")
		} else {
			fmt.Printf("= false\n")
		}
		break
	case C.SR_ENUM_T:
		val := (**C.char)(unsafe.Pointer(&value.data))
		fmt.Printf("= %s\n", C.GoString(*val))
		break
	case C.SR_DECIMAL64_T:
		val := (*C.double)(unsafe.Pointer(&value.data))
		fmt.Printf("= %d\n", *val)
		break
	case C.SR_INT8_T:
		val := (*C.int8_t)(unsafe.Pointer(&value.data))
		fmt.Printf("= %d\n", *val)
		break
	case C.SR_INT16_T:
		val := (*C.int16_t)(unsafe.Pointer(&value.data))
		fmt.Printf("= %d\n", *val)
		break
	case C.SR_INT32_T:
		val := (*C.int32_t)(unsafe.Pointer(&value.data))
		fmt.Printf("= %d\n", *val)
		break
	case C.SR_INT64_T:
		val := (*C.int64_t)(unsafe.Pointer(&value.data))
		fmt.Printf("= %d\n", *val)
		break
	case C.SR_UINT8_T:
		val := (*C.uint8_t)(unsafe.Pointer(&value.data))
		fmt.Printf("= %d\n", *val)
		break
	case C.SR_UINT16_T:
		val := (*C.uint16_t)(unsafe.Pointer(&value.data))
		fmt.Printf("= %d\n", *val)
		break
	case C.SR_UINT32_T:
		val := (*C.uint32_t)(unsafe.Pointer(&value.data))
		fmt.Printf("= %d\n", *val)
		break
	case C.SR_UINT64_T:
		val := (*C.uint64_t)(unsafe.Pointer(&value.data))
		fmt.Printf("= %d\n", *val)
		break
	case C.SR_IDENTITYREF_T:
		val := (**C.char)(unsafe.Pointer(&value.data))
		fmt.Printf("= %s\n", C.GoString(*val))
		break
	case C.SR_BITS_T:
		val := (**C.char)(unsafe.Pointer(&value.data))
		fmt.Printf("= %s\n", C.GoString(*val))
		break
	case C.SR_BINARY_T:
		val := (**C.char)(unsafe.Pointer(&value.data))
		fmt.Printf("= %s\n", C.GoString(*val))
		break
	default:
		fmt.Printf("(unprintable)\n")
	}
}
*/
