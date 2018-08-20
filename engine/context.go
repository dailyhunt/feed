package engine

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

// TODO: store for engine's context
type Context struct {
	Request *RequestContext
	Profile *ProfileContext
	Config  *ConfigContext
}

type ContextBuilderFn func(c *gin.Context) (*Context)
type RequestContextBuilderFn func(c *gin.Context) (*RequestContext)
type ProfileContextBuilderFn func(request *RequestContext) (*ProfileContext)
type ConfigContextBuilderFn func(request *RequestContext) (*ConfigContext)

func ContextBuilder(requestBuilder RequestContextBuilderFn, profileBuilder ProfileContextBuilderFn, configBuilder ConfigContextBuilderFn) (ContextBuilderFn) {
	return func(c *gin.Context) (*Context) {
		return nil
	}
}
