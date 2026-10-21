package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewArticleEntityFunc func(client *SpaceflightNewsSDK, entopts map[string]any) SpaceflightNewsEntity

var NewBlogEntityFunc func(client *SpaceflightNewsSDK, entopts map[string]any) SpaceflightNewsEntity

var NewInfoEntityFunc func(client *SpaceflightNewsSDK, entopts map[string]any) SpaceflightNewsEntity

var NewReportEntityFunc func(client *SpaceflightNewsSDK, entopts map[string]any) SpaceflightNewsEntity

