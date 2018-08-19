package dsl

type Pipeline struct {
	Id              string
	SourcePipelines SourcePipeline
}

// should we be passing source query ?
func (pipeline *Pipeline) Source() (*SourcePipeline) {
	// it will create a new source

	return nil
}
