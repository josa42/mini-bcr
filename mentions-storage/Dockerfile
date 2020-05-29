FROM golang:1.14

COPY . /src/.

WORKDIR /src/mentions-storage

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["mentions-storage"]

