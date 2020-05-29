# Mini BCR: mastodon-crawler

Crawl toots via mastodon Streaming API and publish them as mentions to kafka.

## Configuration

- `MASTODON_SERVER`
- `MASTODON_CLIENT_ID`
- `MASTODON_CLIENT_SECRET`
- `MASTODON_USERNAME`
- `MASTODON_PASSWORD`
- `KAFKA_HOST`
- `KAFKA_PUBLISH_TOPIC`

## Run

```sh
go run main.go
```

## License

[MIT © Josa Gesell](../LICENSE)

