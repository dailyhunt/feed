package dsl

import (
	"github.com/dailyhunt/feed/context"
)

type EngineDefn struct {
	Feeds            []FeedDefn
	RequestBuilderFn context.RequestBuilderFn
	ProfileBuilderFn context.ProfileBuilderFn
	ConfigBuilderFn  context.ConfigBuilderFn
}

func (defn *EngineDefn) RequestBuilder(fn context.RequestBuilderFn) *EngineDefn {
	defn.RequestBuilderFn = fn

	return defn
}

func (defn *EngineDefn) ProfileBuilder(fn context.ProfileBuilderFn) *EngineDefn {
	defn.ProfileBuilderFn = fn

	return defn
}

func (defn *EngineDefn) ConfigBuilder(fn context.ConfigBuilderFn) *EngineDefn {
	defn.ConfigBuilderFn = fn

	return defn
}

func (defn *EngineDefn) Feed(id string) *FeedDefn {
	// create a new feed
	var feedDefn = new(FeedDefn)
	feedDefn.Id = id

	defn.Feeds = append(defn.Feeds, *feedDefn)

	return feedDefn
}
