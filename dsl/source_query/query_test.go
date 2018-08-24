package query

import (
	"github.com/dailyhunt/feed/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

// query from config map
var qCountConfig = map[string]map[string]int{
	"breaking": {
		"breaking": 10},
	"genre": {
		"genre-1": 20,
		"genre-2": 5},
}

func TestNewQueueQuery(t *testing.T) {

	// TODO: check for all other wise intersection

	// test Empty
	emptyQuery := NewQueueQuery()
	assert.NotNil(t, emptyQuery)

	queueQuery := NewQueueQueryFromMap(qCountConfig)

	assert.Equal(t, 1, len(queueQuery[QBreaking]))
	assert.Equal(t, 2, len(queueQuery[QGenre]))

	assert.Equal(t, 10, queueQuery[QBreaking]["breaking"])
	assert.Equal(t, 20, queueQuery[QGenre]["genre-1"], "genre-1 count should be 20")
	assert.Equal(t, 5, queueQuery[QGenre]["genre-2"], "genre-2 count should be 5")
}

func TestInstanceOfSourceQuery(t *testing.T) {

	queueQuery := NewQueueQueryFromMap(qCountConfig)
	queueQuery.Add(QGenre, "genre-3", 5)

	sourceQuery := NewSourceQueryBuilder().
		Select(queueQuery).
		From(util.HINDI).
		Filter(GENRE, MatchAny("Genre1", "Genre100")).
		Filter(GENRE, MatchAll("Genre3", "Genre4", "Genre5")).
		Filter(IG, MatchAll("IG3", "IG4")).
		Filter(NSFW, MatchNone("Racy")).
		Filter(NSFW, MatchNone("Adult")).
		Build()

	assert.NotNil(t, sourceQuery)

	genreAnyFilters := sourceQuery.orFilters[GENRE]
	assert.Equal(t, 1, len(genreAnyFilters))
	assert.Equal(t, 2, len(genreAnyFilters[0].Fields()))

	igAllFilters := sourceQuery.andFilters[IG]
	assert.Equal(t, 1, len(igAllFilters))
	assert.Equal(t, 2, len(igAllFilters[0].Fields()))

	genreAllFilters := sourceQuery.andFilters[GENRE]
	assert.Equal(t, 1, len(genreAllFilters))
	assert.Equal(t, 3, len(genreAllFilters[0].Fields()))

	notNsfwFilters := sourceQuery.notFilters[NSFW]
	assert.Equal(t, 2, len(notNsfwFilters))
	assert.Equal(t, 1, len(notNsfwFilters[1].Fields()))

}
