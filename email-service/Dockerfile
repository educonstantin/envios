FROM golang:1.11 as builder

WORKDIR /go/src/github.com/educonstantin/envios/email-service

COPY . .

RUN go get -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/educonstantin/envios/email-service .

CMD ["./email-service"]

# docker build -t email-service .
# docker run --net="host" -p 50054 -e MICRO_SERVER_ADDRESS=:50054 -e MICRO_REGISTRY=mdns email-service