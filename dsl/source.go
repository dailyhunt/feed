package dsl

type SourcePipeline struct {
	// TODO: define source queries
	ItemPipeline ItemPipeline
}

// should we be passing source query ?
func (source *SourcePipeline) Source() (*SourcePipeline) {
	// it will create a new source

	return nil
}

func (source *SourcePipeline) Items() (*ItemPipeline) {
	return nil
}
