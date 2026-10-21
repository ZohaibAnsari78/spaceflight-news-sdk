# SpaceflightNews SDK configuration


# The sekreto plugin DEFINITIONS the model selected per feature, imported
# above by name from the modules the catalogue's active `plugin.def`
# entries declare. Handed to each feature (secrets builds its Sekreto
# with them): a provider kind not listed here is unknown to that SDK.
FEATURE_PLUGINS = {
}


_shared_config = None


def shared_config():
    """Return the process-wide config, built once on first use.

    The SDK reads the config on every request and never writes to it, so one
    instance is shared by every client rather than rebuilt per client.

    The returned dict is shared: treat it as read-only. Callers that need to
    mutate should use make_config, which always returns a fresh copy.
    """
    global _shared_config
    if _shared_config is None:
        _shared_config = make_config()
    return _shared_config


def make_config():
    """Build a fresh, fully materialised config dict.

    Every call rebuilds the whole structure, so prefer shared_config unless
    you need a private copy you intend to mutate.
    """
    return {
        "main": {
            "name": "SpaceflightNews",
            "slug": "spaceflight-news",
            "version": "0.0.1",
            "target": "py",
        },
        "feature": {
            "test": {
        "options": {
          "active": False,
        },
        "optspec": {
          "entity": "`$MAP`",
          "net": "`$MAP`",
        },
        "strict": False,
        "transport": "base",
      },
        },
        "options": {
            "base": "https://api.spaceflightnewsapi.net",
            "auth": {
                "prefix": "Bearer",
            },
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "article": {},
                "blog": {},
                "info": {},
                "report": {},
            },
        },
        "entity": {
      "article": {
        "fields": [
          {
            "name": "authors",
            "title": "Authors",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "events",
            "title": "Events",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "featured",
            "title": "Featured",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "id",
            "title": "Id",
            "type": "`$INTEGER`",
            "req": True,
            "readOnly": True,
          },
          {
            "name": "image_url",
            "title": "Image Url",
            "type": "`$STRING`",
            "req": True,
            "format": "uri",
          },
          {
            "name": "launches",
            "title": "Launches",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "news_site",
            "title": "News Site",
            "type": "`$STRING`",
            "req": True,
            "readOnly": True,
          },
          {
            "name": "published_at",
            "title": "Published At",
            "type": "`$STRING`",
            "req": True,
            "format": "date-time",
          },
          {
            "name": "summary",
            "title": "Summary",
            "type": "`$STRING`",
            "req": True,
          },
          {
            "name": "title",
            "title": "Title",
            "type": "`$STRING`",
            "req": True,
          },
          {
            "name": "updated_at",
            "title": "Updated At",
            "type": "`$STRING`",
            "req": True,
            "readOnly": True,
            "format": "date-time",
          },
          {
            "name": "url",
            "title": "Url",
            "type": "`$STRING`",
            "req": True,
            "format": "uri",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
        },
        "name": "article",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/v4/articles/",
                "segments": [
                  {
                    "lit": "v4",
                  },
                  {
                    "lit": "articles",
                  },
                ],
                "parts": [
                  "v4",
                  "articles",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.results`",
                },
                "args": {
                  "query": [
                    {
                      "name": "event",
                      "orig": "event",
                      "type": "`$ARRAY`",
                      "kind": "query",
                    },
                    {
                      "name": "has_event",
                      "orig": "has_event",
                      "type": "`$BOOLEAN`",
                      "kind": "query",
                    },
                    {
                      "name": "has_launch",
                      "orig": "has_launch",
                      "type": "`$BOOLEAN`",
                      "kind": "query",
                    },
                    {
                      "name": "is_featured",
                      "orig": "is_featured",
                      "type": "`$BOOLEAN`",
                      "kind": "query",
                    },
                    {
                      "name": "launch",
                      "orig": "launch",
                      "type": "`$ARRAY`",
                      "kind": "query",
                    },
                    {
                      "name": "limit",
                      "orig": "limit",
                      "type": "`$INTEGER`",
                      "kind": "query",
                    },
                    {
                      "name": "news_site",
                      "orig": "news_site",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "news_site_exclude",
                      "orig": "news_site_exclude",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "offset",
                      "orig": "offset",
                      "type": "`$INTEGER`",
                      "kind": "query",
                    },
                    {
                      "name": "ordering",
                      "orig": "ordering",
                      "type": "`$ARRAY`",
                      "kind": "query",
                    },
                    {
                      "name": "published_at_gt",
                      "orig": "published_at_gt",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "published_at_gte",
                      "orig": "published_at_gte",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "published_at_lt",
                      "orig": "published_at_lt",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "published_at_lte",
                      "orig": "published_at_lte",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "search",
                      "orig": "search",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "summary_contain",
                      "orig": "summary_contains",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "summary_contains_all",
                      "orig": "summary_contains_all",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "summary_contains_one",
                      "orig": "summary_contains_one",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "title_contain",
                      "orig": "title_contains",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "title_contains_all",
                      "orig": "title_contains_all",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "title_contains_one",
                      "orig": "title_contains_one",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "updated_at_gt",
                      "orig": "updated_at_gt",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "updated_at_gte",
                      "orig": "updated_at_gte",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "updated_at_lt",
                      "orig": "updated_at_lt",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "updated_at_lte",
                      "orig": "updated_at_lte",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {},
                "response": {
                  "kind": "json",
                  "media": "application/json",
                },
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/v4/articles/{id}/",
                "segments": [
                  {
                    "lit": "v4",
                  },
                  {
                    "lit": "articles",
                  },
                  {
                    "var": "id",
                  },
                ],
                "parts": [
                  "v4",
                  "articles",
                  "{id}",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$INTEGER`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "response": {
                  "kind": "json",
                  "media": "application/json",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "blog": {
        "fields": [
          {
            "name": "authors",
            "title": "Authors",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "events",
            "title": "Events",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "featured",
            "title": "Featured",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "id",
            "title": "Id",
            "type": "`$INTEGER`",
            "req": True,
            "readOnly": True,
          },
          {
            "name": "image_url",
            "title": "Image Url",
            "type": "`$STRING`",
            "req": True,
            "format": "uri",
          },
          {
            "name": "launches",
            "title": "Launches",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "news_site",
            "title": "News Site",
            "type": "`$STRING`",
            "req": True,
            "readOnly": True,
          },
          {
            "name": "published_at",
            "title": "Published At",
            "type": "`$STRING`",
            "req": True,
            "format": "date-time",
          },
          {
            "name": "summary",
            "title": "Summary",
            "type": "`$STRING`",
            "req": True,
          },
          {
            "name": "title",
            "title": "Title",
            "type": "`$STRING`",
            "req": True,
          },
          {
            "name": "updated_at",
            "title": "Updated At",
            "type": "`$STRING`",
            "req": True,
            "readOnly": True,
            "format": "date-time",
          },
          {
            "name": "url",
            "title": "Url",
            "type": "`$STRING`",
            "req": True,
            "format": "uri",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
        },
        "name": "blog",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/v4/blogs/",
                "segments": [
                  {
                    "lit": "v4",
                  },
                  {
                    "lit": "blogs",
                  },
                ],
                "parts": [
                  "v4",
                  "blogs",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.results`",
                },
                "args": {
                  "query": [
                    {
                      "name": "event",
                      "orig": "event",
                      "type": "`$ARRAY`",
                      "kind": "query",
                    },
                    {
                      "name": "has_event",
                      "orig": "has_event",
                      "type": "`$BOOLEAN`",
                      "kind": "query",
                    },
                    {
                      "name": "has_launch",
                      "orig": "has_launch",
                      "type": "`$BOOLEAN`",
                      "kind": "query",
                    },
                    {
                      "name": "is_featured",
                      "orig": "is_featured",
                      "type": "`$BOOLEAN`",
                      "kind": "query",
                    },
                    {
                      "name": "launch",
                      "orig": "launch",
                      "type": "`$ARRAY`",
                      "kind": "query",
                    },
                    {
                      "name": "limit",
                      "orig": "limit",
                      "type": "`$INTEGER`",
                      "kind": "query",
                    },
                    {
                      "name": "news_site",
                      "orig": "news_site",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "news_site_exclude",
                      "orig": "news_site_exclude",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "offset",
                      "orig": "offset",
                      "type": "`$INTEGER`",
                      "kind": "query",
                    },
                    {
                      "name": "ordering",
                      "orig": "ordering",
                      "type": "`$ARRAY`",
                      "kind": "query",
                    },
                    {
                      "name": "published_at_gt",
                      "orig": "published_at_gt",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "published_at_gte",
                      "orig": "published_at_gte",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "published_at_lt",
                      "orig": "published_at_lt",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "published_at_lte",
                      "orig": "published_at_lte",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "search",
                      "orig": "search",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "summary_contain",
                      "orig": "summary_contains",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "summary_contains_all",
                      "orig": "summary_contains_all",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "summary_contains_one",
                      "orig": "summary_contains_one",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "title_contain",
                      "orig": "title_contains",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "title_contains_all",
                      "orig": "title_contains_all",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "title_contains_one",
                      "orig": "title_contains_one",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "updated_at_gt",
                      "orig": "updated_at_gt",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "updated_at_gte",
                      "orig": "updated_at_gte",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "updated_at_lt",
                      "orig": "updated_at_lt",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "updated_at_lte",
                      "orig": "updated_at_lte",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {},
                "response": {
                  "kind": "json",
                  "media": "application/json",
                },
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/v4/blogs/{id}/",
                "segments": [
                  {
                    "lit": "v4",
                  },
                  {
                    "lit": "blogs",
                  },
                  {
                    "var": "id",
                  },
                ],
                "parts": [
                  "v4",
                  "blogs",
                  "{id}",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$INTEGER`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "response": {
                  "kind": "json",
                  "media": "application/json",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "info": {
        "fields": [
          {
            "name": "news_sites",
            "title": "News Sites",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "version",
            "title": "Version",
            "type": "`$STRING`",
            "req": True,
          },
        ],
        "name": "info",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/v4/info/",
                "segments": [
                  {
                    "lit": "v4",
                  },
                  {
                    "lit": "info",
                  },
                ],
                "parts": [
                  "v4",
                  "info",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.news_sites`",
                },
                "args": {},
                "select": {},
                "response": {
                  "kind": "json",
                  "media": "application/json",
                },
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/v4/info/",
                "segments": [
                  {
                    "lit": "v4",
                  },
                  {
                    "lit": "info",
                  },
                ],
                "parts": [
                  "v4",
                  "info",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {},
                "response": {
                  "kind": "json",
                  "media": "application/json",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "report": {
        "fields": [
          {
            "name": "authors",
            "title": "Authors",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "id",
            "title": "Id",
            "type": "`$INTEGER`",
            "req": True,
            "readOnly": True,
          },
          {
            "name": "image_url",
            "title": "Image Url",
            "type": "`$STRING`",
            "req": True,
            "format": "uri",
          },
          {
            "name": "news_site",
            "title": "News Site",
            "type": "`$STRING`",
            "req": True,
            "readOnly": True,
          },
          {
            "name": "published_at",
            "title": "Published At",
            "type": "`$STRING`",
            "req": True,
            "format": "date-time",
          },
          {
            "name": "summary",
            "title": "Summary",
            "type": "`$STRING`",
          },
          {
            "name": "title",
            "title": "Title",
            "type": "`$STRING`",
            "req": True,
          },
          {
            "name": "updated_at",
            "title": "Updated At",
            "type": "`$STRING`",
            "req": True,
            "readOnly": True,
            "format": "date-time",
          },
          {
            "name": "url",
            "title": "Url",
            "type": "`$STRING`",
            "req": True,
            "format": "uri",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
        },
        "name": "report",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/v4/reports/",
                "segments": [
                  {
                    "lit": "v4",
                  },
                  {
                    "lit": "reports",
                  },
                ],
                "parts": [
                  "v4",
                  "reports",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.results`",
                },
                "args": {
                  "query": [
                    {
                      "name": "limit",
                      "orig": "limit",
                      "type": "`$INTEGER`",
                      "kind": "query",
                    },
                    {
                      "name": "news_site",
                      "orig": "news_site",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "news_site_exclude",
                      "orig": "news_site_exclude",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "offset",
                      "orig": "offset",
                      "type": "`$INTEGER`",
                      "kind": "query",
                    },
                    {
                      "name": "ordering",
                      "orig": "ordering",
                      "type": "`$ARRAY`",
                      "kind": "query",
                    },
                    {
                      "name": "published_at_gt",
                      "orig": "published_at_gt",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "published_at_gte",
                      "orig": "published_at_gte",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "published_at_lt",
                      "orig": "published_at_lt",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "published_at_lte",
                      "orig": "published_at_lte",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "search",
                      "orig": "search",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "summary_contain",
                      "orig": "summary_contains",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "summary_contains_all",
                      "orig": "summary_contains_all",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "summary_contains_one",
                      "orig": "summary_contains_one",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "title_contain",
                      "orig": "title_contains",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "title_contains_all",
                      "orig": "title_contains_all",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "title_contains_one",
                      "orig": "title_contains_one",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "updated_at_gt",
                      "orig": "updated_at_gt",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "updated_at_gte",
                      "orig": "updated_at_gte",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "updated_at_lt",
                      "orig": "updated_at_lt",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "updated_at_lte",
                      "orig": "updated_at_lte",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {},
                "response": {
                  "kind": "json",
                  "media": "application/json",
                },
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/v4/reports/{id}/",
                "segments": [
                  {
                    "lit": "v4",
                  },
                  {
                    "lit": "reports",
                  },
                  {
                    "var": "id",
                  },
                ],
                "parts": [
                  "v4",
                  "reports",
                  "{id}",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$INTEGER`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "response": {
                  "kind": "json",
                  "media": "application/json",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
    },
    }
