FROM golang:1.19

WORKDIR /usr/app/

# copy relevant go files from current src dir to image work folder:
COPY assets/ assets/
COPY configs/ configs/
COPY internal/ internal/
COPY pkg/ pkg/
COPY go.mod go.mod

RUN go mod tidy && \
    go build ./... && \
    go install ./...

# Run tests
CMD CGO_ENABLED=0 go test ./...