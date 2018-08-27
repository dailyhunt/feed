package dsl

type FeedDefn struct {
	Id       string
	Sections []SectionDefn

	// TODO: add params to it for A/B, platform, or param based selection
}

func (feed *FeedDefn) Section(id string) *SectionDefn {
	// create a new section
	var sectionDefn = new(SectionDefn)
	sectionDefn.Id = id

	feed.Sections = append(feed.Sections, *sectionDefn)

	return sectionDefn
}
