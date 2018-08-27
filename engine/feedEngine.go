package engine

import (
	"github.com/dailyhunt/feed/dsl"
	"github.com/gin-gonic/gin"
	"github.com/dailyhunt/feed/context"
)

type FeedEngine struct {
	// could be executed parallel.
	// TODO: build DAG basis depends on.
	Sections []SectionEngine
}

func BuildFeedEngine(defn *dsl.FeedDefn) (*FeedEngine) {
	var e = new(FeedEngine)

	for _, s := range defn.Sections {
		e.Sections = append(e.Sections, *BuildSectionEngine(&s))
	}

	return e
}

func (engine *FeedEngine) Serve(ginContext *gin.Context, engineContext *context.Context) {

}
