FROM golang:1.11 as builder
WORKDIR /go/src/consignment-cli
RUN export GOPATH=/go/src/consignment-cli
COPY . /go/src/consignment-cli
RUN go get -v
RUN CGO_ENABLED=0 GOOS=linux go build -o consignment-cli -a -installsuffix cgo .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
COPY consignment.json /app/consignment.json
COPY --from=builder /go/src/consignment-cli/consignment-cli .
CMD ["./consignment-cli"]

# docker build -t consignment-cli .
# docker run -e MICRO_REGISTRY=mdns consignment-cli
