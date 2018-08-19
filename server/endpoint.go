package server

import (
	"github.com/gin-gonic/gin"
	"github.com/dailyhunt/feed/dsl"
)

type HttpMethod string

const (
	GetMethod  HttpMethod = "GET"
	PostMethod HttpMethod = "POST"
)

type HttpEndpoint struct {
	Path               string
	Method             HttpMethod
	FeedContextBuilder FeedContextBuilder
	FeedDsl            *dsl.Feed
}

type FeedContextBuilder func(c *gin.Context)
