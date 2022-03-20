package app

import (
	"fmt"
	"github.com/jdkato/prose/v2"
	"github.com/marvin-hansen/arxiv/v1"
	"scigraph/src/utils/dbg_utils"
)

// https://github.com/jdkato/prose
const mtd = "processNLPHandler: "

func (a *App) processNLPHandler(entry *arxiv.Entry) {

	title := entry.Title
	println("Title: ", title)

	abstract := entry.Summary.Body
	println("Abstract: ", abstract)

	dbgPrint(mtd + "Parsing test")
	doc, err := prose.NewDocument(title)
	dbg_utils.CheckPrintErr(err, "error paring test! ")

	dbgPrint(mtd + "Iterate over the title's named-entities")
	for _, ent := range doc.Entities() {
		fmt.Println("Entity: ", ent.Text)
	}

	doc, err = prose.NewDocument(abstract)

	dbgPrint(mtd + "Iterate over the abstract's named-entities")
	for _, ent := range doc.Entities() {
		fmt.Println("Entity: ", ent.Text)
	}
}
