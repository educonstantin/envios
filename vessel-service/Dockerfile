FROM golang:1.11 as builder
RUN mkdir /go/src/vessel-service
RUN export GOPATH=/go/src/vessel-service
WORKDIR /go/src/vessel-service
COPY . .
RUN go get -v
RUN CGO_ENABLED=0 GOOS=linux go build -o vessel-service -a -installsuffix cgo main.go repository.go handler.go datastore.go


FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/vessel-service/vessel-service .
CMD ["./vessel-service"]

# docker build -t vessel-service .
# docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns vessel-service