package server

import (
	"github.com/dailyhunt/feed/dsl"
)

type HttpMethod string

const (
	GetMethod  HttpMethod = "GET"
	PostMethod HttpMethod = "POST"
)

type HttpEndpoint struct {
	Path       string
	Method     HttpMethod
	EngineDefn *dsl.EngineDefn
}
