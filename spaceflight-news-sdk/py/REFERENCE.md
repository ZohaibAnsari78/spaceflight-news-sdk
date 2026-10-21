# SpaceflightNews Python SDK Reference

Complete API reference for the SpaceflightNews Python SDK.


## SpaceflightNewsSDK

### Constructor

```python
from spaceflightnews_sdk import SpaceflightNewsSDK

client = SpaceflightNewsSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["apikey"]` | `str` | API key for authentication. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `SpaceflightNewsSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = SpaceflightNewsSDK.test()
```


### Instance Methods

#### `Article(data=None)`

Create a new `ArticleEntity` instance. Pass `None` for no initial data.

#### `Blog(data=None)`

Create a new `BlogEntity` instance. Pass `None` for no initial data.

#### `Info(data=None)`

Create a new `InfoEntity` instance. Pass `None` for no initial data.

#### `Report(data=None)`

Create a new `ReportEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## ArticleEntity

```python
article = client.Article()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authors` | `list` | Yes |  |
| `events` | `list` | Yes |  |
| `featured` | `bool` | No |  |
| `id` | `int` | Yes |  |
| `image_url` | `str` | Yes |  |
| `launches` | `list` | Yes |  |
| `news_site` | `str` | Yes |  |
| `published_at` | `str` | Yes |  |
| `summary` | `str` | Yes |  |
| `title` | `str` | Yes |  |
| `updated_at` | `str` | Yes |  |
| `url` | `str` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[ArticleEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Article().list()
for article in results:
    print(article.data_get())
```

#### `load(reqmatch, ctrl=None) -> ArticleEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Article().load({"id": 1})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ArticleEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## BlogEntity

```python
blog = client.Blog()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authors` | `list` | Yes |  |
| `events` | `list` | Yes |  |
| `featured` | `bool` | No |  |
| `id` | `int` | Yes |  |
| `image_url` | `str` | Yes |  |
| `launches` | `list` | Yes |  |
| `news_site` | `str` | Yes |  |
| `published_at` | `str` | Yes |  |
| `summary` | `str` | Yes |  |
| `title` | `str` | Yes |  |
| `updated_at` | `str` | Yes |  |
| `url` | `str` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[BlogEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Blog().list()
for blog in results:
    print(blog.data_get())
```

#### `load(reqmatch, ctrl=None) -> BlogEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Blog().load({"id": 1})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `BlogEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## InfoEntity

```python
info = client.Info()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `news_sites` | `list` | Yes |  |
| `version` | `str` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[InfoEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Info().list()
for info in results:
    print(info.data_get())
```

#### `load(reqmatch, ctrl=None) -> InfoEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Info().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `InfoEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ReportEntity

```python
report = client.Report()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authors` | `list` | Yes |  |
| `id` | `int` | Yes |  |
| `image_url` | `str` | Yes |  |
| `news_site` | `str` | Yes |  |
| `published_at` | `str` | Yes |  |
| `summary` | `str` | No |  |
| `title` | `str` | Yes |  |
| `updated_at` | `str` | Yes |  |
| `url` | `str` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list[ReportEntity]`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list of entities, one per record, and raises on error.

```python
results = client.Report().list()
for report in results:
    print(report.data_get())
```

#### `load(reqmatch, ctrl=None) -> ReportEntity`

Load a single entity matching the given criteria. Returns the entity, whose record `data_get()` reads, and raises on error.

```python
result = client.Report().load({"id": 1})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ReportEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | Test transport |


Features are activated via the `feature` option:

```python
client = SpaceflightNewsSDK({
    "feature": {
        "test": {"active": True},
    },
})
```


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### `test`

Test transport.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

| Option | Type |
|---|---|
| `entity` | map |
| `net` | map |

These take no default: the feature behaves one way when you supply them and
another when you do not.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

