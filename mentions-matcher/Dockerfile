FROM golang:1.14

COPY . /src/.

WORKDIR /src/mentions-matcher

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["twitter-crawler"]
CMD ["mentions-matcher"]

