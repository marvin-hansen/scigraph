package printing_utils

import "scigraph/src/utils/dbg_utils"

func PrintInitHeader(debug bool, main, msg string) {
	DbgPrint(debug, main)
	DbgPrint(debug, "===========================")
	DbgPrint(debug, msg)
	DbgPrint(debug, "===========================")
}

func DbgPrint(debug bool, msg string) {
	dbg_utils.DbgPrint(debug, msg)
}
