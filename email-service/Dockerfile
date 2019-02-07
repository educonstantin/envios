FROM golang:1.11 as builder
RUN mkdir /go/src/email-service
RUN export GOPATH=/go/src/email-service
WORKDIR /go/src/email-service

COPY . .

RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/email-service .

CMD ["./email-service"]

# docker build -t email-service .
# docker run --net="host" -p 50054 -e MICRO_SERVER_ADDRESS=:50054 -e MICRO_REGISTRY=mdns email-service