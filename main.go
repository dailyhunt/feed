package main

import (
	"github.com/dailyhunt/feed/dsl"
	"github.com/gin-gonic/gin"
	"github.com/dailyhunt/feed/cmd"
	"github.com/dailyhunt/feed/server"
	"github.com/dailyhunt/feed/engine"
)

var AppName = "feed"
var Version = "0.0.1"

func main() {
	// TODO: there needs to be top level data source and config loaders
	cmd.Execute(AppName, Version, endpointsBuilder)
}

// only method to be implemented
func endpointsBuilder() (endpoints []server.HttpEndpoint) {
	endpoints = make([]server.HttpEndpoint, 1)

	endpoints = append(endpoints,
		server.HttpEndpoint{Path: "inbox",
			Method: server.PostMethod,
			EngineContextBuilder: engine.ContextBuilder(requestBuilder, profileBuilder, configBuilder),
			FeedSet: inboxFeed()},
		server.HttpEndpoint{Path: "chrono",
			Method: server.PostMethod,
			EngineContextBuilder: engine.ContextBuilder(requestBuilder, profileBuilder, configBuilder),
			FeedSet: chronoFeed()},
	)

	return endpoints
}

func configBuilder(request *engine.RequestContext) (*engine.ConfigContext) {
	// TODO: add config builder
	return nil
}

func profileBuilder(request *engine.RequestContext) (*engine.ProfileContext) {
	// TODO: add profile builder
	return nil
}

func requestBuilder(c *gin.Context) (*engine.RequestContext) {
	return nil
}

// TODO: add data source builder... there can be multiple data sources
func dataSourceBuilder() {

}

func chronoFeed() (*dsl.FeedSet) {
	var feedSet = new(dsl.FeedSet)

	feedSet.
		Feed("news").
		Section("bnews").
		Pipeline("bnews_pipeline"). // pipeline
		Source(). // pipeline
		Source(). // pipeline
		Items(). // collection of sources of items
		Filter(). // filters at item level
		Filter(). // filters at item level
		Score(). // score at item level
		Sort(). // collect all and heap sort
		TopK() // bounded heap

	return feedSet;
}

func inboxFeed() (*dsl.FeedSet) {
	return nil
}
