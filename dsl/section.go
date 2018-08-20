package dsl

type Section struct {
	Id        string
	Pipelines []Pipeline
}

func (section *Section) Pipeline(id string) (pipeline *Pipeline) {
	// create a new pipeline
	return
}
