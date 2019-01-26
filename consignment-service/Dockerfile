FROM golang:1.11 as builder
RUN mkdir /go/src/consignment-service
RUN export GOPATH=/go/src/consignment-service
WORKDIR /go/src/consignment-service
COPY . .
RUN go get -v
RUN CGO_ENABLED=0 GOOS=linux go build -o consignment-service -a -installsuffix cgo main.go repository.go handler.go datastore.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/consignment-service/consignment-service .
CMD ["./consignment-service"]

# docker build -t consignment-service .
# docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns consignment-service