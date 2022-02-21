package app

import (
	"github.com/marvin-hansen/arxiv/v1"
	"testing"
)

func TestFetchAllPapersBySearchTerm(t *testing.T) {
	a := NewApp()
	terms := "Generative Causal Explanations, Graph Neural Networks"
	a.FetchAllPapersBySearchTerm(terms)
}

func TestFetchAllPapersByArticleIDs(t *testing.T) {
	a := NewApp()
	id := "2104.06643,"
	a.FetchPaperByArticleIDs(id)
}

func TestFetchPaperByTitleAndCategory(t *testing.T) {
	a := NewApp()
	title := "Generative Causal Explanations for Graph Neural Networks"
	category := arxiv.CSLearning
	a.FetchPaperByTitleAndCategory(title, category)
}

func TestFetchPaperByTitleAndAuthor(t *testing.T) {
	a := NewApp()
	title := "Generative Causal Explanations for Graph Neural Networks"
	authors := "Wanyu Lin, Hao Lan, Baochun Li"
	a.FetchPaperByTitleAndAuthor(authors, title)
}

func TestFetchPaperByAuthor(t *testing.T) {
	a := NewApp()
	authors := "Wanyu Lin"
	a.FetchAllPaperByAuthor(authors)
}
func TestFetchPaperByTitle(t *testing.T) {
	a := NewApp()
	title := "Dynamic Uncertain Causality Graph"
	a.FetchAllPaperByTitle(title)
}

func TestFetchPaperByTitleAndAbstractKeywords(t *testing.T) {
	a := NewApp()
	keywords := "Cubic, Dynamic Uncertain Causality Graph, Graph, Complex, "

	a.FetchAllPaperByKeywords(keywords)
}
