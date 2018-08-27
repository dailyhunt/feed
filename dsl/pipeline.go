package dsl

type PipelineDefn struct {
	Id              string
	SourcePipelines SourcePipeline
}

// should we be passing source query ?
func (pipeline *PipelineDefn) Source() (*SourcePipeline) {
	// it will create a new source

	return nil
}
