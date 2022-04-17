package app

import (
	"encoding/json"
	"fmt"
	"github.com/jdkato/prose/v2"
	"github.com/marvin-hansen/arxiv/v1"
	"scigraph/src/types/kb_types"
	"scigraph/src/utils/dbg_utils"
)

func (a *App) processNLPGraphHandler(entry *arxiv.Entry) {

	p := kb_types.NewPublication(entry)
	var conceptArray []*kb_types.Concept

	title := entry.Title
	doc, err := prose.NewDocument(title)
	dbg_utils.CheckPrintErr(err, "error paring test! ")
	for _, ent := range doc.Entities() {
		if !contains(conceptArray, ent.Text) {
			c := kb_types.NewConcept(ent.Text)
			conceptArray = append(conceptArray, c)
		}
	}

	abstract := entry.Summary.Body
	doc, err = prose.NewDocument(abstract)
	for _, ent := range doc.Entities() {
		if !contains(conceptArray, ent.Text) {
			c := kb_types.NewConcept(ent.Text)
			conceptArray = append(conceptArray, c)
		}
	}

	p.Concept = conceptArray

	//println(p.String())
	//
	// to print generated JSON, we typecast it to a string
	data, _ := json.MarshalIndent(p, "", " ")
	fmt.Println(string(data))

}
