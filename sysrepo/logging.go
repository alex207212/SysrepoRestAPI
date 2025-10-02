package sysrepo

/*
 #include <sysrepo.h>
*/
import "C"

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
