package engine

import (
	"github.com/dailyhunt/feed/dsl"
	"github.com/gin-gonic/gin"
	"github.com/dailyhunt/feed/context"
)

type Engine struct {
	// TODO: build effective search DS to find applicable FeedEngine
	Feeds            []FeedEngine
	RequestBuilderFn context.RequestBuilderFn
	ProfileBuilderFn context.ProfileBuilderFn
	ConfigBuilderFn  context.ConfigBuilderFn
}

func Build(defn *dsl.EngineDefn) (*Engine) {
	var e = new(Engine)

	e.RequestBuilderFn = defn.RequestBuilderFn
	e.ProfileBuilderFn = defn.ProfileBuilderFn
	e.ConfigBuilderFn = defn.ConfigBuilderFn

	for _, f := range defn.Feeds {
		e.Feeds = append(e.Feeds, *BuildFeedEngine(&f))
	}

	return e
}

func (engine *Engine) Serve(ginContext *gin.Context) {
	engineContext, err := engine.buildContext(ginContext)
	if err != nil {
		// handle error
	}

	// find relevant FeedEngine
	feedEngine, err := engine.findFeedEngine(engineContext)
	if err != nil {
		// handle error
	}

	feedEngine.Serve(ginContext, engineContext)
}

func (engine *Engine) findFeedEngine(engineContext *context.Context) (feedEngine *FeedEngine, err error) {
	return &engine.Feeds[0], nil
}

func (engine *Engine) buildContext(ginContext *gin.Context) (engineContext *context.Context, err error) {
	var requestContext = engine.RequestBuilderFn(ginContext)
	var configContext = engine.ConfigBuilderFn(requestContext)
	var profileContext = engine.ProfileBuilderFn(requestContext)

	engineContext = new(context.Context)
	engineContext.Request = requestContext
	engineContext.Config = configContext
	engineContext.Profile = profileContext

	return
}
