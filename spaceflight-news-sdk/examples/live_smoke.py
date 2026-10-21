# Live smoke test against the real Spaceflight News API (no key needed).
# Install first: pip install -e py, then: python examples/live_smoke.py
from spaceflightnews_sdk import SpaceflightNewsSDK

client = SpaceflightNewsSDK({})

articles = client.Article().list({"limit": 3})
print(f"articles: {len(articles)}")
for article in articles:
    data = article.data_get()
    print(f"  {data['id']}  [{data['news_site']}]  {data['title']}")

first = client.Article().load({"id": articles[0].data_get()["id"]})
print(f"loaded article {first.data_get()['id']}: {first.data_get()['url']}")

reports = client.Report().list({"limit": 2})
print(f"reports: {len(reports)}")

info = client.Info().load()
print(f"api version: {info.data_get()['version']}")
