package sysrepo

/*
 #include <sysrepo.h>
*/
import "C"
import (
	"fmt"
	"runtime"
)

//goland:noinspection GoExportedElementShouldHaveComment
const (
	SR_LL_NONE = iota
	SR_LL_ERR
	SR_LL_WRN
	SR_LL_INF
	SR_LL_DBG
)

func GetLogLevel() int {
	var level C.sr_log_level_t
	level = C.sr_log_get_stderr()
	return int(level)
}

func SetLogLevel(level int) {
	C.sr_log_stderr(C.sr_log_level_t(level))
}

func CheckForError(rc C.int) error {
	if C.SR_ERR_OK != rc {
		var callerFuncName string
		c_err_msg := C.sr_strerror(rc)
		// Do not free c_err_msg (from above) as it returns values from statically allocated array

		err_msg := C.GoString(c_err_msg)
		pc, _, _, ok := runtime.Caller(1)
		if !ok {
			callerFuncName = "unknown caller"
		}
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			callerFuncName = "unknown caller"
		}
		callerFuncName = fn.Name()

		return fmt.Errorf("failed to call %s: %v", callerFuncName, err_msg)
	}
	return nil
}
