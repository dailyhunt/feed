package dsl

type SectionDefn struct {
	Id        string
	Pipelines []PipelineDefn
}

func (section *SectionDefn) Pipeline(id string) (*PipelineDefn) {
	// create a new pipeline
	var pipelineDefn = new(PipelineDefn)
	pipelineDefn.Id = id

	section.Pipelines = append(section.Pipelines, *pipelineDefn)

	return pipelineDefn
}
