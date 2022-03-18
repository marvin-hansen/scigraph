package kb_types

import (
	"github.com/marvin-hansen/arxiv/v1"
)

func convertTimeStrToString(entry arxiv.TimeStr) string {
	return TimeStr(entry).String()
}
