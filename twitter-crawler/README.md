# Mini BCR: twitter-crawler

Crawl tweets via Twitter Streaming API and publish them as mentions to kafka.

## Configuration

- `TWITTER_CONSUMER_KEY`
- `TWITTER_CONSUMER_SECRET`
- `TWITTER_TOKEN`
- `TWITTER_TOKEN_SECRET`
- `KAFKA_HOST`
- `KAFKA_PUBLISH_TOPIC`

## Run

```sh
go run main.go
```

## License

[MIT © Josa Gesell](../LICENSE)

