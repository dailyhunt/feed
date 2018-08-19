package main

import (
	"github.com/dailyhunt/feed/dsl"
	"github.com/gin-gonic/gin"
	"github.com/dailyhunt/feed/cmd"
	"github.com/dailyhunt/feed/server"
)

var AppName = "feed"
var Version = "0.0.1"

func main() {
	cmd.Execute(AppName, Version, endpointsBuilder)
}

// only method to be implemented
func endpointsBuilder() (endpoints []server.HttpEndpoint) {
	endpoints = make([]server.HttpEndpoint, 1)

	endpoints = append(endpoints,
		server.HttpEndpoint{Path: "inbox", Method: server.PostMethod, FeedContextBuilder: feedContextBuilder, FeedDsl: inboxFeed()},
		server.HttpEndpoint{Path: "chrono", Method: server.PostMethod, FeedContextBuilder: feedContextBuilder, FeedDsl: chronoFeed()},
	)

	return endpoints
}

// TODO: add config builder
// TODO: add profile builder
// TODO: add data source builder... there can be multiple data sources

func chronoFeed() (f *dsl.Feed) {
	f = dsl.New()

	// attaches a pipeline
	f.
		Pipeline("bnews_pipeline"). // pipeline
		Source(). // pipeline
		Source(). // pipeline
		Items(). // collection of sources of items
		Filter(). // filters at item level
		Filter(). // filters at item level
		Score(). // score at item level
		Sort(). // collect all and heap sort
		TopK() // bounded heap

	return f
}

func inboxFeed() (feed *dsl.Feed) {
	return
}

func feedContextBuilder(c *gin.Context) {

}
