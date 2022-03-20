package app

import "scigraph/src/utils/dbg_utils"

func (a App) init() {
	a.state.handler = a.processNLPGraphHandler // a.processNLPHandler //a.processPrintEntryHandler
	a.verifyInit()
}

func (a App) verifyInit() {
	dbg_utils.NilCheck(a.state.handler, main+"NPE: Process Handler is NIL. Fix init. ")
}
