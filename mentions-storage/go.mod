module github.com/josa42/mini-bcr/mentions-storage

go 1.14

require (
	github.com/elastic/go-elasticsearch v0.0.0
	github.com/elastic/go-elasticsearch/v7 v7.7.0
	github.com/josa42/mini-bcr/pkg v0.0.0-00010101000000-000000000000
)

replace github.com/josa42/mini-bcr/pkg => ../pkg
