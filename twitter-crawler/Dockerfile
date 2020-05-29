FROM golang:1.14

COPY . /src/.

WORKDIR /src/twitter-crawler

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["twitter-crawler"]
