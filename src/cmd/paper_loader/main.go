package main

import "scigraph/src/cmd/paper_loader/app"

func main() {
	a := app.NewApp()
	id := "2104.06643,"
	a.FetchPaperByArticleIDs(id)

}
