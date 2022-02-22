package app

import "github.com/marvin-hansen/arxiv/v1"

type State struct {
	handler func(entry *arxiv.Entry)
}

func newState() *State {
	return &State{
		handler: nil,
	}
}
