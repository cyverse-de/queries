FROM golang:1.17-alpine

WORKDIR /go/src/github.com/cyverse-de/queries
COPY . /go/src/github.com/cyverse-de/queries

CMD ["go", "test", "github.com/cyverse-de/queries"]
