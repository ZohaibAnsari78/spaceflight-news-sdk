# SpaceflightNews Golang SDK



The Golang SDK for the SpaceflightNews API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

It exposes the API as capitalised, semantic **Entities** — e.g. `client.Article(nil)` — each with the same small set of operations (`List`, `Load`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Also generated from this model: `py`, `ts` — see
> the [top-level README](../README.md).


## Install
```bash
go get github.com/ZohaibAnsari78/spaceflight-news-sdk/go@latest
```

The Go module proxy resolves the version from the `go/vX.Y.Z` GitHub
release tag — see [Tags](https://github.com/ZohaibAnsari78/spaceflight-news-sdk/tags) for the available versions.

To vendor from a local checkout instead, clone this repo alongside your
project and add a `replace` directive pointing at the checked-out
`go/` directory:

```bash
go mod edit -replace github.com/ZohaibAnsari78/spaceflight-news-sdk/go=../spaceflight-news-sdk/go
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### Quickstart

A complete program: create a client, then call the entity operations.
Each operation returns `(value, error)` — the value is the entity, and for
`List` a `[]any` of entities, one per record (there is no `{ok, data}`
wrapper), so check `err` and read a record through the entity's
`Data()`.

```go
package main

import (
    "fmt"
    "os"
    sdk "github.com/ZohaibAnsari78/spaceflight-news-sdk/go"
)

func main() {
    client := sdk.NewSpaceflightNewsSDK(map[string]any{
        "apikey": os.Getenv("SPACEFLIGHT_NEWS_APIKEY"),
    })

    // List article records — the value is a []any of entities, one per record.
    articles, err := client.Article(nil).List(nil, nil)
    if err != nil {
        panic(err)
    }
    for _, item := range articles.([]any) {
        fmt.Println(item.(sdk.Entity).Data())
    }

    // Load a single article — the value is the entity; Data() reads its record.
    article, err := client.Article(nil).Load(map[string]any{"id": 1}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(article.(sdk.Entity).Data())
}
```


## Error handling

Every entity operation returns `(value, error)`. Check `err` before
using the value — there is no exception to catch:

```go
blogs, err := client.Blog(nil).List(nil, nil)
if err != nil {
    // handle err
    return
}
_ = blogs
```

`Direct` follows the same `(value, error)` convention:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example_id"},
})
if err != nil {
    // handle err
}
_ = result
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

if result["ok"] == true {
    fmt.Println(result["status"]) // 200
    fmt.Println(result["data"])   // response body
}
```

### Prepare a request without sending it

```go
fetchdef, err := client.Prepare(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "DELETE",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

fmt.Println(fetchdef["url"])
fmt.Println(fetchdef["method"])
fmt.Println(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```go
client := sdk.Test()

blogs, err := client.Blog(nil).List(
    nil, nil,
)
if err != nil {
    panic(err)
}
// A []any of entities, one per mock record.
for _, item := range blogs.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```go
mockFetch := func(url string, init map[string]any) (map[string]any, error) {
    return map[string]any{
        "status":     200,
        "statusText": "OK",
        "headers":    map[string]any{},
        "json": (func() any)(func() any {
            return map[string]any{"id": "mock01"}
        }),
    }, nil
}

client := sdk.NewSpaceflightNewsSDK(map[string]any{
    "base": "http://localhost:8080",
    "system": map[string]any{
        "fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
    },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
SPACEFLIGHT_NEWS_TEST_LIVE=TRUE
SPACEFLIGHT_NEWS_APIKEY=<your-key>
```

Then run:

```bash
cd go && go test ./test/...
```


## Reference

### NewSpaceflightNewsSDK

```go
func NewSpaceflightNewsSDK(options map[string]any) *SpaceflightNewsSDK
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `"apikey"` | `string` | API key for authentication. |
| `"base"` | `string` | Base URL of the API server. |
| `"prefix"` | `string` | URL path prefix prepended to all requests. |
| `"suffix"` | `string` | URL path suffix appended to all requests. |
| `"feature"` | `map[string]any` | Feature activation flags. |
| `"extend"` | `[]any` | Additional Feature instances to load. |
| `"system"` | `map[string]any` | System overrides (e.g. custom `"fetch"` function). |

### TestSDK

```go
func TestSDK(testopts map[string]any, sdkopts map[string]any) *SpaceflightNewsSDK
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### SpaceflightNewsSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `OptionsMap` | `() map[string]any` | Deep copy of current SDK options. |
| `GetUtility` | `() *Utility` | Copy of the SDK utility object. |
| `Prepare` | `(fetchargs map[string]any) (map[string]any, error)` | Build an HTTP request definition without sending. |
| `Direct` | `(fetchargs map[string]any) (map[string]any, error)` | Build and send an HTTP request. |
| `Article` | `(data map[string]any) SpaceflightNewsEntity` | Create an Article entity instance. |
| `Blog` | `(data map[string]any) SpaceflightNewsEntity` | Create a Blog entity instance. |
| `Info` | `(data map[string]any) SpaceflightNewsEntity` | Create an Info entity instance. |
| `Report` | `(data map[string]any) SpaceflightNewsEntity` | Create a Report entity instance. |

### Entity interface (SpaceflightNewsEntity)

All entities implement the `SpaceflightNewsEntity` interface.

| Method | Signature | Description |
| --- | --- | --- |
| `Load` | `(reqmatch, ctrl map[string]any) (any, error)` | Load a single entity by match criteria, and return it. |
| `List` | `(reqmatch, ctrl map[string]any) (any, error)` | List entities matching the criteria, one per record. |
| `Data` | `(args ...any) any` | Get or set entity data. |
| `Match` | `(args ...any) any` | Get or set entity match criteria. |
| `Make` | `() Entity` | Create a new instance with the same options. |
| `GetName` | `() string` | Return the entity name. |

### Result shape

Entity operations return `(value, error)`. The `value` is the entity
itself — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `Load` | the entity, whose `Data()` reads its record (`map[string]any`) |
| `List` | a `[]any` of entities, one per record |

Check `err` first, then use the value directly (or the typed
`...Typed` variants, which return the entity's model struct and a typed
slice):

    article, err := client.Article(nil).List(nil, nil)
    if err != nil { /* handle */ }
    // article is a []any of entities, one per record

Only `Direct()` returns a response envelope — a `map[string]any` with
`"ok"`, `"status"`, `"headers"`, and `"data"` keys.

### Entities

#### Article

| Field | Description |
| --- | --- |
| `"authors"` |  |
| `"events"` |  |
| `"featured"` |  |
| `"id"` |  |
| `"image_url"` |  |
| `"launches"` |  |
| `"news_site"` |  |
| `"published_at"` |  |
| `"summary"` |  |
| `"title"` |  |
| `"updated_at"` |  |
| `"url"` |  |

Operations: List, Load.

API path: `/v4/articles/`

#### Blog

| Field | Description |
| --- | --- |
| `"authors"` |  |
| `"events"` |  |
| `"featured"` |  |
| `"id"` |  |
| `"image_url"` |  |
| `"launches"` |  |
| `"news_site"` |  |
| `"published_at"` |  |
| `"summary"` |  |
| `"title"` |  |
| `"updated_at"` |  |
| `"url"` |  |

Operations: List, Load.

API path: `/v4/blogs/`

#### Info

| Field | Description |
| --- | --- |
| `"news_sites"` |  |
| `"version"` |  |

Operations: List, Load.

API path: `/v4/info/`

#### Report

| Field | Description |
| --- | --- |
| `"authors"` |  |
| `"id"` |  |
| `"image_url"` |  |
| `"news_site"` |  |
| `"published_at"` |  |
| `"summary"` |  |
| `"title"` |  |
| `"updated_at"` |  |
| `"url"` |  |

Operations: List, Load.

API path: `/v4/reports/`



## Entities


### Article

Create an instance: `article := client.Article(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authors` | `[]any` |  |
| `events` | `[]any` |  |
| `featured` | `bool` |  |
| `id` | `int` |  |
| `image_url` | `string` |  |
| `launches` | `[]any` |  |
| `news_site` | `string` |  |
| `published_at` | `string` |  |
| `summary` | `string` |  |
| `title` | `string` |  |
| `updated_at` | `string` |  |
| `url` | `string` |  |

#### Example: Load

```go
article, err := client.Article(nil).Load(map[string]any{"id": 1}, nil)
if err != nil {
    panic(err)
}
fmt.Println(article.(sdk.Entity).Data()) // the loaded entity's record
```

#### Example: List

```go
articles, err := client.Article(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range articles.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### Blog

Create an instance: `blog := client.Blog(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authors` | `[]any` |  |
| `events` | `[]any` |  |
| `featured` | `bool` |  |
| `id` | `int` |  |
| `image_url` | `string` |  |
| `launches` | `[]any` |  |
| `news_site` | `string` |  |
| `published_at` | `string` |  |
| `summary` | `string` |  |
| `title` | `string` |  |
| `updated_at` | `string` |  |
| `url` | `string` |  |

#### Example: Load

```go
blog, err := client.Blog(nil).Load(map[string]any{"id": 1}, nil)
if err != nil {
    panic(err)
}
fmt.Println(blog.(sdk.Entity).Data()) // the loaded entity's record
```

#### Example: List

```go
blogs, err := client.Blog(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range blogs.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### Info

Create an instance: `info := client.Info(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `news_sites` | `[]any` |  |
| `version` | `string` |  |

#### Example: Load

```go
info, err := client.Info(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(info.(sdk.Entity).Data()) // the loaded entity's record
```

#### Example: List

```go
infos, err := client.Info(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range infos.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```


### Report

Create an instance: `report := client.Report(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authors` | `[]any` |  |
| `id` | `int` |  |
| `image_url` | `string` |  |
| `news_site` | `string` |  |
| `published_at` | `string` |  |
| `summary` | `string` |  |
| `title` | `string` |  |
| `updated_at` | `string` |  |
| `url` | `string` |  |

#### Example: Load

```go
report, err := client.Report(nil).Load(map[string]any{"id": 1}, nil)
if err != nil {
    panic(err)
}
fmt.Println(report.(sdk.Entity).Data()) // the loaded entity's record
```

#### Example: List

```go
reports, err := client.Report(nil).List(nil, nil)
if err != nil {
    panic(err)
}
// A []any of entities, one per record.
for _, item := range reports.([]any) {
    fmt.Println(item.(sdk.Entity).Data())
}
```

## Features

This SDK ships 1 optional features. Each is **inactive until you
switch it on**, so an SDK you have not configured behaves exactly as if none of
them existed — no retries, no cache, no logging, no measurable overhead.

Activate a feature by name in the client options, alongside the options shown
above:

| Feature | What it does |
|---|---|
| [`test`](#test) | Test transport |

### test

Test transport.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.test.active` to enable it, then override any of the options above.


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature implements the
`Feature` interface and provides hooks — functions keyed by pipeline
stage names.

The SDK ships with built-in features:

- **TestFeature**: Test transport

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as maps

The Go SDK uses `map[string]any` throughout rather than typed structs.
This mirrors the dynamic nature of the API and keeps the SDK
flexible — no code generation is needed when the API schema changes.

An operation returns the entity, and its `Data()` returns the record. Use
`core.ToMapAny()` to safely cast that record, or data nested in it, to
`map[string]any`: it returns `nil` for anything else, an entity included.

### Package structure

```
github.com/ZohaibAnsari78/spaceflight-news-sdk/go/
├── spaceflight-news.go        # Root package — type aliases and constructors
├── core/               # SDK core — client, types, pipeline
├── entity/             # Entity implementations
├── feature/            # Built-in features (Base, Test, Log)
├── utility/            # Utility functions and struct library
└── test/               # Test suites
```

The root package (`github.com/ZohaibAnsari78/spaceflight-news-sdk/go`) re-exports everything needed
for normal use. Import sub-packages only when you need specific types
like `core.ToMapAny`.

### Entity state

Entity instances are stateful. After a successful `List`, the entity
stores the returned data and match criteria internally.

```go
blog := client.Blog(nil)
blog.List(nil, nil)

// blog.Data() now returns the blog data from the last list
// blog.Match() returns the last match criteria
```

Call `Make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`Direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `Prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
