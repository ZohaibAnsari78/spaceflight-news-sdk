# Spaceflight News API

> The Spaceflight News API (SNAPI) is a product by [The Space Devs](https://thespacedevs.com) (TSD). It&#39;s the most complete and up-to-date spaceflight news API currently available.
>
> While this API is **free to use**, we do encourage developers to support us through [Patreon](https://www.patreon.com/TheSpaceDevs) to keep the API up and running.
>
>  ### GraphQL
>
>  The Spaceflight News API also has GraphQL support available! You can find the GraphiQl IDE [here](https://api.spaceflightnewsapi.net/v4/graphql/).
>
>  ### FAQs &amp; Tutorials
>
>  - [GitHub repository](https://github.com/TheSpaceDevs/Tutorials/): contains FAQs and tutorials for TSD APIs
>
>  - [TSD FAQ](https://github.com/TheSpaceDevs/Tutorials/blob/main/faqs/faq_TSD.md): TSD-specific FAQ (for example history, network, funding, etc.)
>
>  - [SNAPI FAQ](https://github.com/TheSpaceDevs/Tutorials/blob/main/faqs/faq_SNAPI.md): SNAPI-specific FAQ
>
>  ### Feedback &amp; Support
>
>  If you need any help with SNAPI, you can ask in the [`💬feedback-and-help`](https://discord.com/channels/676725644444565514/1019976345884827750) forum of the TSD [Discord server](https://discord.gg/p7ntkNA) or email [derk@spaceflightnewsapi.net](mailto:derk@spaceflightnewsapi.net).

## Start here

This guide introduces the API, the client libraries, and the companion tools in this repository. Start with the API capabilities, choose a client for your application, and use the linked reference when you need exact request and response details.

The selected API surface contains 4 entities and 8 HTTP routes. There are 3 SDK targets.

An entity groups related API operations. An operation can have several routes with different inputs or authentication requirements. The SDK exposes the entity and its operations using the conventions of the selected language.

## What the API provides

### Article

SDK operations: `list`, `load`.

### Blog

SDK operations: `list`, `load`.

### Info

SDK operations: `list`, `load`.

### Report

SDK operations: `list`, `load`.

### Route map

Use this map to locate a capability. Consult the entity reference before supplying request data; routes for the same operation can require different fields.

| Entity | SDK operation | HTTP route | Authentication |
| --- | --- | --- | --- |
| Article | `list` | `GET /v4/articles/` | Not required |
| Article | `load` | `GET /v4/articles/{id}/` | Not required |
| Blog | `list` | `GET /v4/blogs/` | Not required |
| Blog | `load` | `GET /v4/blogs/{id}/` | Not required |
| Info | `list` | `GET /v4/info/` | Not required |
| Info | `load` | `GET /v4/info/` | Not required |
| Report | `list` | `GET /v4/reports/` | Not required |
| Report | `load` | `GET /v4/reports/{id}/` | Not required |

## Connect to the API

- API server: `https://api.spaceflightnewsapi.net`

Check authentication for the route you plan to call. A route that declares no authentication can be used without credentials; this does not change the requirements of other routes. Keep credentials in environment variables or a configured secret provider, and keep them out of source control and logs.

## Make a first request

1. Choose the API server and an operation that matches your task.
2. Check the operation’s required input and authentication. Use values valid for your account and environment.
3. Send one request and inspect the returned data before adding retries, concurrency, or a larger batch.

A read request without required parameters or authentication is `GET /v4/articles/`. For example:

```sh
curl --fail-with-body --silent --show-error 'https://api.spaceflightnewsapi.net/v4/articles/'
```

Inspect the response using the Article reference. This checks the public route; authenticated operations need their own credentials and request data.

For an SDK call, install or build the chosen client, create a client instance with its documented configuration, and call the required entity operation. Language references describe the argument shape, asynchronous behaviour, and returned values.

## Choose an SDK

Choose the language already used by your application or service. The clients represent the same API model, while package setup, naming, and return types follow each language. Check the selected client’s reference and tests before integrating it into an existing application.

| Client | Repository directory | Distribution |
| --- | --- | --- |
| Golang | `go/` | Build from source |
| Python | `py/` | Build from source |
| TypeScript | `ts/` | Build from source |

Build-from-source entries are not marked as published in the project model. Follow the build instructions in that target’s README, then consume the resulting package using your language’s local dependency mechanism. Published entries give the installation command recorded for that client.

## Operational features

Features supply behaviour around API calls, such as request handling, diagnostics, or local testing. Inclusion in this project does not mean a feature is enabled at runtime. Check the selected SDK’s supported features and configuration defaults, then enable the behaviour your application needs.

- `test`: In-memory mock transport for testing without a live server

Start with the default client configuration. Add request limits and diagnostics as needed, test error paths, and review retry behaviour before using operations that change data. A retry can repeat an operation unless the API provides a suitable guarantee.

## Continue with the documentation

- Follow the first-call guide for the setup sequence.
- Read the authentication guide before using protected routes.
- Use the API reference for request schemas, response formats, and status codes.
- Check the chosen SDK or companion tool reference for its configuration and supported operations.

