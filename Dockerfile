FROM golang:1.19

WORKDIR /usr/app/

# copy relevant go files from current src dir to image work folder:
COPY assets/ assets/
COPY configs/ configs/
COPY internal/ internal/
COPY pkg/ pkg/
COPY go.mod go.mod

# https://es.wikipedia.org/wiki/Uname
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN go mod tidy && \
    go build ./... && \
    go install ./...

# Run tests
CMD CGO_ENABLED=0 go test ./...