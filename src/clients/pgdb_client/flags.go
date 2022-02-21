package pgdb_client

import "scigraph/src/utils/dbg_utils"

const (
	debug = false
	main  = "PGdbClient: "
)

func dbgPrint(msg string) {
	dbg_utils.DbgPrint(debug, main+msg)
}
