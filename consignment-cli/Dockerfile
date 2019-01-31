FROM golang:1.11 as builder
WORKDIR /go/src/consignment-cli
COPY . .
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
COPY consignment.json /app/consignment.json
COPY --from=builder /go/src/consignment-cli .
ENTRYPOINT ["./consignment-cli"]

# docker build -t consignment-cli .
# docker run --net="host" -e MICRO_REGISTRY=mdns consignment-cli consignment.json <TOKEN_HERE>