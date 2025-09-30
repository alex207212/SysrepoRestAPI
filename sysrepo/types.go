package sysrepo

/*
 #include <sysrepo_types.h>
*/
import "C"
import "log"

type Datastore int

const (
	DS_STARTUP = iota
	DS_RUNNING
	DS_CANDIDATE
	DS_OPERATIONAL
	DS_FACTORY_DEFAULT
)

func (ds Datastore) String() string {
	return [...]string{"Startup", "Running", "Candidate", "Operational", "FactoryDefault"}[ds]
}

func (ds Datastore) to_sr_datastore_t() C.sr_datastore_t {
	switch ds {
	case DS_STARTUP:
		return C.SR_DS_STARTUP
	case DS_RUNNING:
		return C.SR_DS_RUNNING
	case DS_CANDIDATE:
		return C.SR_DS_CANDIDATE
	case DS_OPERATIONAL:
		return C.SR_DS_OPERATIONAL
	case DS_FACTORY_DEFAULT:
		return C.SR_DS_FACTORY_DEFAULT
	}
	log.Fatal("Unsupported datastore type", ds)
	return 0
}

type Session struct {
	session   *C.sr_session_ctx_t
	datastore C.sr_datastore_t
}
