FROM golang:1.17.6-alpine3.14
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
RUN mkdir -p /build
WORKDIR /build
COPY go.* /build/
RUN go mod download
COPY . .
RUN go build -o nuri-sochainAPI
CMD ["./nuri-sochainAPI"]
EXPOSE 8080