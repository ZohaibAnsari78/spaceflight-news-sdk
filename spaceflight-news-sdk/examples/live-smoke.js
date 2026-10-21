// Live smoke test against the real Spaceflight News API (no key needed).
// Build first: (cd ts && npm install && npm run build), then: node examples/live-smoke.js
const { SpaceflightNewsSDK } = require('../ts/dist/SpaceflightNewsSDK.js')

async function main() {
  const client = new SpaceflightNewsSDK()

  const articles = await client.Article().list({ limit: 3 })
  console.log(`articles: ${articles.length}`)
  for (const article of articles) {
    const { id, title, news_site } = article.data()
    console.log(`  ${id}  [${news_site}]  ${title}`)
  }

  const first = await client.Article().load({ id: articles[0].data().id })
  console.log(`loaded article ${first.data().id}: ${first.data().url}`)

  const reports = await client.Report().list({ limit: 2 })
  console.log(`reports: ${reports.length}`)

  const info = await client.Info().load()
  console.log(`api version: ${info.data().version}`)
}

main().catch((err) => {
  console.error('live smoke failed:', err.message || err)
  process.exitCode = 1
})
