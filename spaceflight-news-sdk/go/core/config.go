package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "SpaceflightNews",
			"slug": "spaceflight-news",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"entity": "`$MAP`",
					"net": "`$MAP`",
				},
				"strict": false,
				"transport": "base",
			},
		},
		"options": map[string]any{
			"base": "https://api.spaceflightnewsapi.net",
			"auth": map[string]any{
				"prefix": "Bearer",
			},
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"article": map[string]any{},
				"blog": map[string]any{},
				"info": map[string]any{},
				"report": map[string]any{},
			},
		},
		"entity": map[string]any{
			"article": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "authors",
						"title": "Authors",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "events",
						"title": "Events",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "featured",
						"title": "Featured",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$INTEGER`",
						"req": true,
						"readOnly": true,
					},
					map[string]any{
						"name": "image_url",
						"title": "Image Url",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
					map[string]any{
						"name": "launches",
						"title": "Launches",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "news_site",
						"title": "News Site",
						"type": "`$STRING`",
						"req": true,
						"readOnly": true,
					},
					map[string]any{
						"name": "published_at",
						"title": "Published At",
						"type": "`$STRING`",
						"req": true,
						"format": "date-time",
					},
					map[string]any{
						"name": "summary",
						"title": "Summary",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"req": true,
						"readOnly": true,
						"format": "date-time",
					},
					map[string]any{
						"name": "url",
						"title": "Url",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "article",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/v4/articles/",
								"segments": []any{
									map[string]any{
										"lit": "v4",
									},
									map[string]any{
										"lit": "articles",
									},
								},
								"parts": []any{
									"v4",
									"articles",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.results`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "event",
											"orig": "event",
											"type": "`$ARRAY`",
											"kind": "query",
										},
										map[string]any{
											"name": "has_event",
											"orig": "has_event",
											"type": "`$BOOLEAN`",
											"kind": "query",
										},
										map[string]any{
											"name": "has_launch",
											"orig": "has_launch",
											"type": "`$BOOLEAN`",
											"kind": "query",
										},
										map[string]any{
											"name": "is_featured",
											"orig": "is_featured",
											"type": "`$BOOLEAN`",
											"kind": "query",
										},
										map[string]any{
											"name": "launch",
											"orig": "launch",
											"type": "`$ARRAY`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "news_site",
											"orig": "news_site",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "news_site_exclude",
											"orig": "news_site_exclude",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "offset",
											"orig": "offset",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "ordering",
											"orig": "ordering",
											"type": "`$ARRAY`",
											"kind": "query",
										},
										map[string]any{
											"name": "published_at_gt",
											"orig": "published_at_gt",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "published_at_gte",
											"orig": "published_at_gte",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "published_at_lt",
											"orig": "published_at_lt",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "published_at_lte",
											"orig": "published_at_lte",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "summary_contain",
											"orig": "summary_contains",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "summary_contains_all",
											"orig": "summary_contains_all",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "summary_contains_one",
											"orig": "summary_contains_one",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "title_contain",
											"orig": "title_contains",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "title_contains_all",
											"orig": "title_contains_all",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "title_contains_one",
											"orig": "title_contains_one",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "updated_at_gt",
											"orig": "updated_at_gt",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "updated_at_gte",
											"orig": "updated_at_gte",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "updated_at_lt",
											"orig": "updated_at_lt",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "updated_at_lte",
											"orig": "updated_at_lte",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/v4/articles/{id}/",
								"segments": []any{
									map[string]any{
										"lit": "v4",
									},
									map[string]any{
										"lit": "articles",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"v4",
									"articles",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$INTEGER`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"blog": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "authors",
						"title": "Authors",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "events",
						"title": "Events",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "featured",
						"title": "Featured",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$INTEGER`",
						"req": true,
						"readOnly": true,
					},
					map[string]any{
						"name": "image_url",
						"title": "Image Url",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
					map[string]any{
						"name": "launches",
						"title": "Launches",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "news_site",
						"title": "News Site",
						"type": "`$STRING`",
						"req": true,
						"readOnly": true,
					},
					map[string]any{
						"name": "published_at",
						"title": "Published At",
						"type": "`$STRING`",
						"req": true,
						"format": "date-time",
					},
					map[string]any{
						"name": "summary",
						"title": "Summary",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"req": true,
						"readOnly": true,
						"format": "date-time",
					},
					map[string]any{
						"name": "url",
						"title": "Url",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "blog",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/v4/blogs/",
								"segments": []any{
									map[string]any{
										"lit": "v4",
									},
									map[string]any{
										"lit": "blogs",
									},
								},
								"parts": []any{
									"v4",
									"blogs",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.results`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "event",
											"orig": "event",
											"type": "`$ARRAY`",
											"kind": "query",
										},
										map[string]any{
											"name": "has_event",
											"orig": "has_event",
											"type": "`$BOOLEAN`",
											"kind": "query",
										},
										map[string]any{
											"name": "has_launch",
											"orig": "has_launch",
											"type": "`$BOOLEAN`",
											"kind": "query",
										},
										map[string]any{
											"name": "is_featured",
											"orig": "is_featured",
											"type": "`$BOOLEAN`",
											"kind": "query",
										},
										map[string]any{
											"name": "launch",
											"orig": "launch",
											"type": "`$ARRAY`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "news_site",
											"orig": "news_site",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "news_site_exclude",
											"orig": "news_site_exclude",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "offset",
											"orig": "offset",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "ordering",
											"orig": "ordering",
											"type": "`$ARRAY`",
											"kind": "query",
										},
										map[string]any{
											"name": "published_at_gt",
											"orig": "published_at_gt",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "published_at_gte",
											"orig": "published_at_gte",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "published_at_lt",
											"orig": "published_at_lt",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "published_at_lte",
											"orig": "published_at_lte",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "summary_contain",
											"orig": "summary_contains",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "summary_contains_all",
											"orig": "summary_contains_all",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "summary_contains_one",
											"orig": "summary_contains_one",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "title_contain",
											"orig": "title_contains",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "title_contains_all",
											"orig": "title_contains_all",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "title_contains_one",
											"orig": "title_contains_one",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "updated_at_gt",
											"orig": "updated_at_gt",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "updated_at_gte",
											"orig": "updated_at_gte",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "updated_at_lt",
											"orig": "updated_at_lt",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "updated_at_lte",
											"orig": "updated_at_lte",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/v4/blogs/{id}/",
								"segments": []any{
									map[string]any{
										"lit": "v4",
									},
									map[string]any{
										"lit": "blogs",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"v4",
									"blogs",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$INTEGER`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"info": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "news_sites",
						"title": "News Sites",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "version",
						"title": "Version",
						"type": "`$STRING`",
						"req": true,
					},
				},
				"name": "info",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/v4/info/",
								"segments": []any{
									map[string]any{
										"lit": "v4",
									},
									map[string]any{
										"lit": "info",
									},
								},
								"parts": []any{
									"v4",
									"info",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.news_sites`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/v4/info/",
								"segments": []any{
									map[string]any{
										"lit": "v4",
									},
									map[string]any{
										"lit": "info",
									},
								},
								"parts": []any{
									"v4",
									"info",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"report": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "authors",
						"title": "Authors",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$INTEGER`",
						"req": true,
						"readOnly": true,
					},
					map[string]any{
						"name": "image_url",
						"title": "Image Url",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
					map[string]any{
						"name": "news_site",
						"title": "News Site",
						"type": "`$STRING`",
						"req": true,
						"readOnly": true,
					},
					map[string]any{
						"name": "published_at",
						"title": "Published At",
						"type": "`$STRING`",
						"req": true,
						"format": "date-time",
					},
					map[string]any{
						"name": "summary",
						"title": "Summary",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "title",
						"title": "Title",
						"type": "`$STRING`",
						"req": true,
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"req": true,
						"readOnly": true,
						"format": "date-time",
					},
					map[string]any{
						"name": "url",
						"title": "Url",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "report",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/v4/reports/",
								"segments": []any{
									map[string]any{
										"lit": "v4",
									},
									map[string]any{
										"lit": "reports",
									},
								},
								"parts": []any{
									"v4",
									"reports",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.results`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "news_site",
											"orig": "news_site",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "news_site_exclude",
											"orig": "news_site_exclude",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "offset",
											"orig": "offset",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "ordering",
											"orig": "ordering",
											"type": "`$ARRAY`",
											"kind": "query",
										},
										map[string]any{
											"name": "published_at_gt",
											"orig": "published_at_gt",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "published_at_gte",
											"orig": "published_at_gte",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "published_at_lt",
											"orig": "published_at_lt",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "published_at_lte",
											"orig": "published_at_lte",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "summary_contain",
											"orig": "summary_contains",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "summary_contains_all",
											"orig": "summary_contains_all",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "summary_contains_one",
											"orig": "summary_contains_one",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "title_contain",
											"orig": "title_contains",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "title_contains_all",
											"orig": "title_contains_all",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "title_contains_one",
											"orig": "title_contains_one",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "updated_at_gt",
											"orig": "updated_at_gt",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "updated_at_gte",
											"orig": "updated_at_gte",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "updated_at_lt",
											"orig": "updated_at_lt",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "updated_at_lte",
											"orig": "updated_at_lte",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/v4/reports/{id}/",
								"segments": []any{
									map[string]any{
										"lit": "v4",
									},
									map[string]any{
										"lit": "reports",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"v4",
									"reports",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$INTEGER`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
