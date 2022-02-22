package app

import "scigraph/src/utils/dbg_utils"

const (
	main = "app: "
	dbg  = true
)

func dbgPrint(msg string) {
	dbg_utils.DbgPrint(dbg, main+msg)
}
