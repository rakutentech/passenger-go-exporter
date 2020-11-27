FROM golang:1.14 AS builder
WORKDIR /go/src/github.com/rakutentech/passenger-go-exporter

COPY go.mod go.sum ./
RUN go mod download && go get github.com/goreleaser/goreleaser
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main . 

ARG GITHUB_TOKEN
RUN goreleaser

FROM scratch
WORKDIR /app
COPY --from=builder /main /app/main
ENTRYPOINT ["/app/main"]
