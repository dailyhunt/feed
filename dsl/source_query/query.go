package query

import (
	"fmt"
	"github.com/dailyhunt/feed/util"
	"strings"
)

const (
	QGenre     QueueType = "GENRE"
	QBreaking            = "BREAKING"
	QMostLiked           = "MOST_LIKED"
	QNPInbox             = "NP_INBOX"
	QInvalid             = "INVALID"

	// Filter Type
	OR  FilterType = "OR"
	AND            = "AND"
	NOT            = "NOT"

	// Field Type
	GENRE        FieldType = "GENRE"
	IG                     = "IG"
	CONTENT_TYPE           = "CONTENT_TYPE"
	NSFW                   = "NSFW"
)

type FilterType string

type Field string
type FieldType string

type QueueType string
type QueueCount map[string]int
type QueueQuery map[QueueType]QueueCount

type SourceQueryFilter interface {
	Type() FilterType
	Fields() []Field
}

type filter struct {
	fType  FilterType
	fields []Field
}

type SourceQuery struct {
	queueQueries QueueQuery
	language     util.Language
	orFilters    map[FieldType][]SourceQueryFilter
	andFilters   map[FieldType][]SourceQueryFilter
	notFilters   map[FieldType][]SourceQueryFilter
}

type SourceQueryBuilder struct {
	queueQueries QueueQuery
	language     util.Language
	orFilters    map[FieldType][]SourceQueryFilter
	andFilters   map[FieldType][]SourceQueryFilter
	notFilters   map[FieldType][]SourceQueryFilter
}

func (q QueueQuery) Add(qType QueueType, name string, count int) QueueQuery {
	if _, ok := q[qType]; !ok {
		qCount := QueueCount{}
		q[qType] = qCount
	}

	queueCount := q[qType]
	queueCount[name] = count

	return q
}

func (q QueueQuery) Count() int {
	return len(q)
}

func NewQueueQueryFromMap(qCountConfig map[string]map[string]int) QueueQuery {

	query := QueueQuery{}

	for qType, qConf := range qCountConfig {
		queueType := GetQueTypeFromString(qType)
		if queueType == QInvalid {
			panic(fmt.Sprintf("invalid queue type passed %s", qType))
		}

		for qName, count := range qConf {
			query.Add(queueType, qName, count)
		}
	}

	return query
}

func NewQueueQuery() *QueueQuery {
	return &QueueQuery{}
}

func GetQueTypeFromString(qType string) QueueType {
	types := map[string]QueueType{
		"GENRE":      QGenre,
		"BREAKING":   QBreaking,
		"MOST_LIKED": QMostLiked,
		"NP_INBOX":   QNPInbox,
	}

	upper := strings.ToUpper(qType)
	if t, ok := types[upper]; ok {
		return t
	}

	return QInvalid
}

func (f *filter) Type() FilterType {
	return f.fType
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

func (b *SourceQueryBuilder) Filter(fieldType FieldType, queryFilter SourceQueryFilter) *SourceQueryBuilder {
	switch fType := queryFilter.Type(); fType {
	case OR:
		b.addFilter(b.orFilters, fieldType, queryFilter)
	case AND:
		b.addFilter(b.andFilters, fieldType, queryFilter)
	case NOT:
		b.addFilter(b.notFilters, fieldType, queryFilter)
	default:
		fmt.Println("Invalid Filter list and Panic ")
	}

	return b
}

func (b *SourceQueryBuilder) addFilter(queryFilters map[FieldType][]SourceQueryFilter, fieldType FieldType, queryFilter SourceQueryFilter) {
	if filters, ok := queryFilters[fieldType]; ok {
		queryFilters[fieldType] = append(filters, queryFilter)
	} else {
		var filters []SourceQueryFilter
		queryFilters[fieldType] = append(filters, queryFilter)
	}
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
	return &SourceQueryBuilder{
		orFilters:  make(map[FieldType][]SourceQueryFilter),
		andFilters: make(map[FieldType][]SourceQueryFilter),
		notFilters: make(map[FieldType][]SourceQueryFilter),
	}
}
