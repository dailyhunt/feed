package query

import (
	"fmt"
	"github.com/dailyhunt/feed/util"
)

type QueueQuery map[string]int

func (q QueueQuery) Add(name string, count int) QueueQuery {
	if _, ok := q[name]; !ok {
		q[name] = count
	}
	return q
}

func (q QueueQuery) Count() int {
	return len(q)
}

func NewQueueQuery() *QueueQuery {
	return &QueueQuery{}
}

type SourceQuery struct {
	queueQueries QueueQuery
	language     util.Language
	orFilters    []SourceQueryFilter
	andFilters   []SourceQueryFilter
	notFilters   []SourceQueryFilter
}

type SourceQueryBuilder struct {
	queueQueries QueueQuery
	language     util.Language
	orFilters    []SourceQueryFilter
	andFilters   []SourceQueryFilter
	notFilters   []SourceQueryFilter
}

func (b *SourceQueryBuilder) Select(queueQueries QueueQuery) *SourceQueryBuilder {
	// Todo : Should handle nil or empty map and panic
	if b.queueQueries == nil {
		b.queueQueries = queueQueries
	}
	return b
}

func (b *SourceQueryBuilder) From(language util.Language) *SourceQueryBuilder {
	if b.language == util.UNKNOWN {
		b.language = language
	}
	return b
}

func (b *SourceQueryBuilder) Filter(filter SourceQueryFilter) *SourceQueryBuilder {
	switch fType := filter.Type(); fType {
	case OR:
		fmt.Println("Add to OR Filter list")
	case AND:
		fmt.Println("Add to OR Filter list")
	case NOT:
		fmt.Println("Add to NOT Filter list")
	default:
		fmt.Println("Invalid Filter list and Panic ")
	}

	return b
}
func (b *SourceQueryBuilder) Build() *SourceQuery {
	return &SourceQuery{
		queueQueries: b.queueQueries,
		language:     b.language,
		orFilters:    b.orFilters,
		andFilters:   b.andFilters,
		notFilters:   b.notFilters,
	}

}

func NewSourceQueryBuilder() *SourceQueryBuilder {
	return &SourceQueryBuilder{}
}
