package engine

import "github.com/dailyhunt/feed/dsl"

type PipelineEngine struct {
}

func BuildPipelineEngine(defn *dsl.PipelineDefn) *PipelineEngine {
	var e = new(PipelineEngine)

	return e
}
