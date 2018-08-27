package dsl

type SourcePipeline struct {
	// TODO: define source queries
	ItemPipeline ItemPipeline
}

// should we be passing source query ?
func (source *SourcePipeline) Source() *SourcePipeline {
	// TODO: it will create a new source, and appends to the list

	return nil
}

func (source *SourcePipeline) Items() *ItemPipeline {
	return nil
}
