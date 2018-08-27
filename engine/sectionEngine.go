package engine

import (
	"github.com/dailyhunt/feed/dsl"
)

type SectionEngine struct {
	// would be executed sequentially
	Pipelines []PipelineEngine
}

func BuildSectionEngine(defn *dsl.SectionDefn) (*SectionEngine) {
	var e = new(SectionEngine)

	for _, p := range defn.Pipelines {
		e.Pipelines = append(e.Pipelines, *BuildPipelineEngine(&p))
	}

	return e
}
