FROM golang:1.11 as builder
WORKDIR /go/src/user-cli
RUN export GOPATH=/go/src/user-cli
COPY . /go/src/user-cli
RUN go get -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/user-cli/user-cli .
CMD ["./user-cli"]

# docker build -t user-cli .
# docker run -e MICRO_REGISTRY=mdns user-cli
