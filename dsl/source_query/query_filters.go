package query

type FilterType string
type Field string

const (
	OR  FilterType = "OR"
	AND            = "AND"
	NOT            = "NOT"
)

type SourceQueryFilter interface {
	Type() FilterType
	Fields() []Field
}

type filter struct {
	fType  FilterType
	fields []Field
}

func (f *filter) Type() FilterType {
	return OR
}

func (f *filter) Fields() []Field {
	return f.fields
}

func MatchAny(fields ...Field) SourceQueryFilter {
	return &filter{OR, fields}
}

func MatchAll(fields ...Field) SourceQueryFilter {
	return &filter{AND, fields}
}

func MatchNone(fields ...Field) SourceQueryFilter {
	return &filter{NOT, fields}
}
