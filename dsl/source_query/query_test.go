package query

import (
	"github.com/dailyhunt/feed/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewQueueQuery(t *testing.T) {
	queueQuery := NewQueueQuery().
		Add("breaking", 10).
		Add("genre1", 20)

	assert.Equal(t, 2, len(queueQuery))
	assert.Equal(t, 20, queueQuery["genre1"])
}

func TestInstanceOfSourceQuery(t *testing.T) {

	queueQuery := NewQueueQuery().
		Add("breaking", 10).
		Add("genre1", 30).
		Add("genre2", 5)

	sourceQuery := NewSourceQueryBuilder().
		Select(queueQuery).
		From(util.HINDI).
		Filter(Or("IG1", "IG2")).
		Filter(And("IG1", "IG2")).
		Build()

	assert.NotNil(t, sourceQuery)

}
