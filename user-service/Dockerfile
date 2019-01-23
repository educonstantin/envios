FROM golang:1.11 as builder
RUN mkdir /go/src/user-service
RUN export GOPATH=/go/src/user-service
WORKDIR /go/src/user-service
COPY . .
RUN go get -v
#RUN CGO_ENABLED=0 GOOS=linux go build -o user-service -a -installsuffix cgo main.go repository.go handler.go database.go token_service.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/user-service/user-service .
CMD ["./user-service"]

# docker build -t user-service .
# docker run -p 50053:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns user-service