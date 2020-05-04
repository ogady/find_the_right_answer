FROM golang:1.14

ENV GO111MODULE=on

WORKDIR /go/src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server server.go


FROM alpine:3.10
RUN apk --no-cache add ca-certificates && \
    export ENV=prod
WORKDIR /root/
COPY --from=0 /go/src .
CMD ["./server"]