# Build image
FROM golang:alpine as builder
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR $GOPATH/src/github.com/vkd/interview-back-message-board/
COPY . .
RUN go mod download
RUN go mod verify

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/message-board cmd/message-board/main.go

# Prod image
# TODO: use scratch, now is not possible because the target system is unknown
FROM alpine
# COPY --fromakem=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/message-board /go/bin/message-board
COPY --from=builder /go/src/github.com/vkd/interview-back-message-board/import_messages.csv /import_messages.csv

CMD ["/go/bin/message-board", "-import-file", "/import_messages.csv"]
