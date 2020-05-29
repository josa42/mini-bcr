module github.com/josa42/mini-bcr/twitter-crawler

go 1.14

require (
	github.com/Shopify/sarama v1.26.4 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/dghubble/go-twitter v0.0.0-20190719072343-39e5462e111f
	github.com/dghubble/oauth1 v0.6.0
	github.com/josa42/mini-bcr/pkg v0.0.0-00010101000000-000000000000
	github.com/klauspost/compress v1.10.6 // indirect
	github.com/pierrec/lz4 v2.5.2+incompatible // indirect
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0 // indirect
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37 // indirect
	golang.org/x/net v0.0.0-20200528225125-3c3fba18258b // indirect
)

replace github.com/josa42/mini-bcr/pkg => ../pkg
