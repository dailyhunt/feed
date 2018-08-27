package context

import "github.com/gin-gonic/gin"

type RequestContext struct {
	// captures request specific data
}

type ProfileContext struct {
	// captures request specific profile
}

type ConfigContext struct {
	// captures request specific config
}

// store for engine's context
type Context struct {
	Request *RequestContext
	Profile *ProfileContext
	Config  *ConfigContext
}

type RequestBuilderFn func(c *gin.Context) (*RequestContext)
type ProfileBuilderFn func(request *RequestContext) (*ProfileContext)
type ConfigBuilderFn func(request *RequestContext) (*ConfigContext)