package dsl

type Feed struct {
	Id string
	Sections []Section

	// TODO: add params to it for A/B, platform, or param based selection
}

func New() (*Feed) {
	return nil
}

func (feed *Feed) Section(id string) (section *Section) {
	// create a new section
	return
}
