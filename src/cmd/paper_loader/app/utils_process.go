package app

import "github.com/marvin-hansen/arxiv/v1"

func (a *App) processPrintEntryHandler(entry *arxiv.Entry) {
	println("ID: ", entry.ID)
	println("Title: ", entry.Title)
	//println("Author: ", entry.Author)
	a.printAuthors(entry.Author)
	println("Published: ", entry.Published)
	println("Updated:", entry.Updated)
	a.printCategories(entry.Category)
}

func (a *App) printAuthors(people []*arxiv.Person) {
	for _, p := range people {
		println("Author: ", p.Name)
	}
}

func (a *App) printCategories(categories []*arxiv.Class) {

	for _, c := range categories {
		println("Category: " + c.Term.String())
	}
}
