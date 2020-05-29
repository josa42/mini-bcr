FROM golang:1.14

COPY . /src/.

WORKDIR /src/mastodon-crawler

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["mastodon-crawler"]
