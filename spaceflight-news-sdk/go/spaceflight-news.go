package zohaibansari78spaceflightnewssdk

import (
	"github.com/ZohaibAnsari78/spaceflight-news-sdk/go/core"
	"github.com/ZohaibAnsari78/spaceflight-news-sdk/go/entity"
	"github.com/ZohaibAnsari78/spaceflight-news-sdk/go/feature"
	_ "github.com/ZohaibAnsari78/spaceflight-news-sdk/go/utility"
)

// Type aliases preserve external API.
type SpaceflightNewsSDK = core.SpaceflightNewsSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type SpaceflightNewsEntity = core.SpaceflightNewsEntity
type StreamItem = core.StreamItem
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type SpaceflightNewsError = core.SpaceflightNewsError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewArticleEntityFunc = func(client *core.SpaceflightNewsSDK, entopts map[string]any) core.SpaceflightNewsEntity {
		return entity.NewArticleEntity(client, entopts)
	}
	core.NewBlogEntityFunc = func(client *core.SpaceflightNewsSDK, entopts map[string]any) core.SpaceflightNewsEntity {
		return entity.NewBlogEntity(client, entopts)
	}
	core.NewInfoEntityFunc = func(client *core.SpaceflightNewsSDK, entopts map[string]any) core.SpaceflightNewsEntity {
		return entity.NewInfoEntity(client, entopts)
	}
	core.NewReportEntityFunc = func(client *core.SpaceflightNewsSDK, entopts map[string]any) core.SpaceflightNewsEntity {
		return entity.NewReportEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewSpaceflightNewsSDK = core.NewSpaceflightNewsSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig
var SharedConfig = core.SharedConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewSpaceflightNewsSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *SpaceflightNewsSDK  { return NewSpaceflightNewsSDK(nil) }
func Test() *SpaceflightNewsSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
