package app

import (
	"context"
	"github.com/marvin-hansen/arxiv/v1"
	"scigraph/src/utils/dbg_utils"
)

// searchAndProcess searches for the given query and processes each entry with the given handler.
func searchAndProcess(query *arxiv.Query, handler func(entry *arxiv.Entry)) {
	ctx := context.Background()
	resChan, cancel, searchErr := arxiv.Search(ctx, query)
	dbg_utils.CheckPrintErr(searchErr, "search query failed")

	for resPage := range resChan {
		if err := resPage.Err; err != nil {
			dbg_utils.CheckPrintErr(err, "search query failed")
			continue
		}

		feed := resPage.Feed
		for _, entry := range feed.Entry {
			handler(entry)
		}

		if resPage.PageNumber >= 2 {
			cancel()
		}
	}
}
