package app

import "github.com/marvin-hansen/arxiv/v1"

const maxPageNumber = 10

func getQueryBySearchTerms(terms string) *arxiv.Query {
	return &arxiv.Query{
		Terms:         terms,
		MaxPageNumber: maxPageNumber,
	}
}

func getQueryByArticleIDs(articleIDs []string) *arxiv.Query {
	return &arxiv.Query{
		ArticleIDs: articleIDs,
	}
}

// getQueryByTitleAndCategory returns a search query for the title keywords in that category
// titleKeywords: Comma separated string
// category: category enum
// maxPageNumber: int64 max number of pates to return. Each page contains ten results by default.
func getQueryByTitleAndCategory(titleKeywords string, category arxiv.Category) *arxiv.Query {
	return &arxiv.Query{
		Filters: []*arxiv.Filter{
			{
				Op: arxiv.OpOR,
				Fields: []*arxiv.Field{
					{Title: titleKeywords},
					{Category: category},
				},
			},
		},
		MaxPageNumber: maxPageNumber,
	}
}

func getQueryByAuthorAndTitle(authors, titleKeywords string) *arxiv.Query {
	return &arxiv.Query{
		Filters: []*arxiv.Filter{
			{
				Op: arxiv.OpAnd,
				Fields: []*arxiv.Field{
					{Author: authors},
					{Title: titleKeywords},
				},
			},
		},
		MaxPageNumber: maxPageNumber,
	}
}

func getQueryAllByAuthor(authors string) *arxiv.Query {
	return &arxiv.Query{
		Filters: []*arxiv.Filter{
			{
				Op: arxiv.OpAnd,
				Fields: []*arxiv.Field{
					{Author: authors},
				},
			},
		},
		MaxPageNumber: maxPageNumber,
	}
}

func getQueryAllByTitle(title string) *arxiv.Query {
	return &arxiv.Query{
		Filters: []*arxiv.Filter{
			{
				Op: arxiv.OpAll,
				Fields: []*arxiv.Field{
					{Title: title},
				},
			},
		},
		MaxPageNumber: maxPageNumber,
	}
}

func getQueryAllByTitleAndAbstractKeywords(keywords string) *arxiv.Query {
	return &arxiv.Query{
		Filters: []*arxiv.Filter{
			{
				Op: arxiv.OpAnd,
				Fields: []*arxiv.Field{
					{Title: keywords},
					{Abstract: keywords},
				},
			},
		},
		MaxPageNumber: maxPageNumber,
	}
}
