package main

import (
	"github.com/dailyhunt/feed/dsl"
	"github.com/gin-gonic/gin"
	"github.com/dailyhunt/feed/cmd"
	"github.com/dailyhunt/feed/server"
	"github.com/dailyhunt/feed/context"
)

var AppName = "feed"
var Version = "0.0.1"

func main() {
	// TODO: there needs to be top level data source and config loaders
	cmd.Execute(AppName, Version, endpointsBuilder)
}

// only method to be implemented
func endpointsBuilder() ([]server.HttpEndpoint) {
	var endpoints = make([]server.HttpEndpoint, 0)

	endpoints = append(endpoints,
		server.HttpEndpoint{Path: "inbox", Method: server.PostMethod, EngineDefn: inboxFeed()},
		server.HttpEndpoint{Path: "chrono", Method: server.PostMethod, EngineDefn: chronoFeed()},
	)

	return endpoints
}

func configBuilder(request *context.RequestContext) (*context.ConfigContext) {
	// TODO: add config builder
	return nil
}

func profileBuilder(request *context.RequestContext) (*context.ProfileContext) {
	// TODO: add profile builder
	return nil
}

func requestBuilder(ginContext *gin.Context) (*context.RequestContext) {
	return nil
}

// TODO: add data source builder... there can be multiple data sources
func dataSourceBuilder() {

}

func chronoFeed() (*dsl.EngineDefn) {
	var feedEngine = new(dsl.EngineDefn)

	feedEngine.
		RequestBuilder(requestBuilder).
		ProfileBuilder(profileBuilder).
		ConfigBuilder(configBuilder).
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

	return feedEngine;
}

func inboxFeed() (*dsl.EngineDefn) {
	var feedEngine = new(dsl.EngineDefn)

	feedEngine.
		RequestBuilder(requestBuilder).
		ProfileBuilder(profileBuilder).
		ConfigBuilder(configBuilder).
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

	return feedEngine;
}
