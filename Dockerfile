FROM golang:1.19-alpine

WORKDIR /usr/app/

# create lib dirs under work space:
RUN mkdir assets configs internal pkg

# copy all go files from current src dir to image work folder:
RUN cp assets/* assets/ && \
    cp configs/* configs/ && \
    cp internal/* internal/ && \
    cp pkg/* pkg/

# check folders were created successfully
RUN ls /usr/app/assets && \
    ls /usr/app/assets && \
    ls /usr/app/assets && \
    ls /usr/app/pkg

# Run tests
CMD CGO_ENABLED=0 go test ./...