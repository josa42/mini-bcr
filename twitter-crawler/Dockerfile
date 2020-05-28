FROM golang:1.14

WORKDIR /srv
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["twitter-crawler"]
