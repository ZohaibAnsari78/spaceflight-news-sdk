# Report: Spaceflight News SDK built with Voxgig sdkgen

**API:** [Spaceflight News API v4](https://api.spaceflightnewsapi.net/v4/docs/), which needs no
authentication. It has 7 paths and 4 entities (Article, Blog, Report, Info). It is not in the
`voxgig-sdk` org.
**Tool:** `@voxgig/create-sdkgen@0.30.8` / `@voxgig/sdkgen@4.35.0`, Node 24, Windows 11.
**Targets:** TypeScript, Python, Go, plus the `test` feature.

## What I did

```bash
npx @voxgig/create-sdkgen spaceflight-news -d ./snapi-openapi.json -o ./spaceflight-news-sdk -t ts,py,go -f test
cd spaceflight-news-sdk/.sdk && npm run generate
cd ../ts && npm install && npm run build && npm test
cd ../py && python -m pytest
```

The scaffold took about 2m20s, almost all of it `npm install`. Generation took about 15s.

## Results

| Check                              | Result                                                               |
| ---------------------------------- | -------------------------------------------------------------------- |
| TypeScript offline tests           | 242 passed, 0 failed, 12 skipped                                     |
| Python offline tests               | 226 passed, 0 failed, 91 skipped                                     |
| Go                                 | Generated but not run, because the Go toolchain is not installed     |
| Request URLs (stubbed `fetch`)     | All correct, e.g. `GET /v4/articles/?limit=3&search=mars`            |
| Live API                           | The upstream API returned 504 on most calls during the session (this happened with plain `curl` too) |

All the skipped tests belong to features I did not select (retry, timeout, and similar).

## Issues found, and the fixes I made

1. **The spec has no `servers` URL, and generation still exits 0.** The log says `PARTIAL BUILD!`
   but `npm run generate` succeeds. The SDK would then ask every caller for a `base` URL.
   _Fix:_ I added `server: 'https://api.spaceflightnewsapi.net'` to `.sdk/build/apidef.js`.
   _Suggestion:_ add a `--server` flag to `create-sdkgen`, and exit non-zero on a partial build.
2. **The defaults assume Voxgig owns the SDK.** The generated `LICENSE` said "Copyright (c) 2026
   Voxgig". The packages were `@voxgig-sdk/...` and `voxgig-sdk-...`, the Go module was
   `github.com/voxgig-sdk/...`, and the README clone URLs pointed to the voxgig-sdk org.
   _Fix:_ I set `main.kit.publisher` and `main.kit.repo.path` in `project.aontu`. The scope comes
   from `origin: 'voxgig-sdk'`, which is hardcoded in `sdk.aontu`, so I edited that file even
   though it is not the file the docs tell you to edit.
   _Suggestion:_ add `--publisher`, `--org` and `--repo` flags, or read them from `git config`.
3. **Single-object endpoint mis-modelled (a data-loss bug).** `/v4/info/` returns one object,
   `{version, news_sites}`. apidef made it a `list` op that returns `body.news_sites`, so
   `Info().list()` returned the site names as "entities" and dropped `version`, with no error.
   _Fix:_ I added a `load` op in `.sdk/model/guide/guide.aontu`. `Info().load()` now returns the
   whole object. The bad `list` op is still there, because I found no way to remove it from the
   guide.
4. **The docs point to the wrong place for edits.** The create-sdkgen docs say to "shape the SDK"
   by editing `.sdk/model/entity/*.aon`. Those files are regenerated, and my edit was silently
   overwritten. The file that is kept is `model/guide/guide.aontu`, and the overlay is
   `project.aontu`, not `project.aon`.
5. **The README shows an API key the API doesn't use.** The README examples pass
   `apikey: process.env.SPACEFLIGHT_NEWS_APIKEY`, but the spec has no `securitySchemes`.
6. **Minor issues:**
   - The landing page describes an interactive walkthrough, but the docs describe flags only.
   - There is no root `.gitattributes`, so on Windows git prints a CRLF warning for each of the
     roughly 1,000 generated files.
   - The 91 skipped Python tests look alarming until you read why they were skipped.

## What worked well

- One command produced three languages, offline tests, `AGENTS.md`, `REFERENCE.md`, CI
  workflows, and tests that compile and run the README examples.
- The entity API (`client.Article().list({ limit: 3 })`) is easy to use. Pagination envelopes
  (`body.results`) were detected correctly for the collection endpoints.
- The model and guide layering is powerful once you find the right file.

## How AI was used

I used Claude Code to research the docs, run the toolchain, read the generated model, and draft
this report. I checked the work this way:

- I ran the generated test suites.
- I stubbed `fetch` to check every request URL against the spec.
- I reproduced the `Info` bug before and after the fix.
- I probed the live API with `curl` to separate upstream outages from SDK bugs.

## Time

The time-box was 30 minutes of human work. Fixing issues 1–3 took most of it. Go testing and a
full live run are not done, because there was no Go toolchain and the upstream API was returning
504s.
