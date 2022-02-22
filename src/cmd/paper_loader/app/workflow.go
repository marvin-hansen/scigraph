package app

import (
	"github.com/marvin-hansen/arxiv/v1"
	"strings"
)

func (a *App) FetchAllPapersBySearchTerm(terms string) {
	q := getQueryBySearchTerms(terms)
	ph := a.state.handler

	searchAndProcess(q, ph)
}

func (a *App) FetchPaperByArticleIDs(articleIDs string) {
	// For bananas reasons, the arxiv go-wrapper only accept a string array,
	// thus the conversion here. see https://github.com/orijtech/arxiv/pulls
	s := strings.Split(articleIDs, ",")
	q := getQueryByArticleIDs(s)
	ph := a.state.handler

	searchAndProcess(q, ph)
}

func (a *App) FetchPaperByTitleAndCategory(titleKeywords string, category arxiv.Category) {
	q := getQueryByTitleAndCategory(titleKeywords, category)
	ph := a.state.handler

	searchAndProcess(q, ph)
}

func (a *App) FetchPaperByTitleAndAuthor(authors, titleKeywords string) {
	q := getQueryByAuthorAndTitle(authors, titleKeywords)
	ph := a.state.handler

	searchAndProcess(q, ph)
}

func (a *App) FetchAllPaperByAuthor(authors string) {
	q := getQueryAllByAuthor(authors)
	ph := a.state.handler

	searchAndProcess(q, ph)
}

func (a *App) FetchAllPaperByTitle(tile string) {
	q := getQueryAllByTitle(tile)
	ph := a.state.handler

	searchAndProcess(q, ph)
}

func (a *App) FetchAllPaperByKeywords(keywords string) {
	q := getQueryAllByTitleAndAbstractKeywords(keywords)
	ph := a.state.handler

	searchAndProcess(q, ph)
}
