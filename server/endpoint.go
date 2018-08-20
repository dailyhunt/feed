package server

import (
	"github.com/dailyhunt/feed/dsl"
	"github.com/dailyhunt/feed/engine"
)

type HttpMethod string

const (
	GetMethod  HttpMethod = "GET"
	PostMethod HttpMethod = "POST"
)

type HttpEndpoint struct {
	Path                 string
	Method               HttpMethod
	EngineContextBuilder engine.ContextBuilderFn
	FeedSet              *dsl.FeedSet
}
