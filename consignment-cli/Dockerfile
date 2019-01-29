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
RUN export GOPATH=/app
COPY consignment.json /app/consignment.json
COPY --from=builder /go/src/consignment-cli/consignment-cli .
CMD ["./consignment-cli"]

# docker build -t consignment-cli .
# docker run --net="host" -e MICRO_REGISTRY=mdns consignment-cli consignment.json eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoiNTM1MWRiZGUtYTk3NC00MWFkLWIwNmQtNjA2NWQ1ZTk0MTM4IiwibmFtZSI6IkVkdWFyZG8gQ29uc3RhbnRpbiIsImNvbXBhbnkiOiJtaWNybyIsImVtYWlsIjoiZWR1Y29uc3RhbnRpbkBnbWFpbC5jb20iLCJwYXNzd29yZCI6IiQyYSQxMCRlMVgvYVR6OU1wSmZpdnZEM2pScWh1bmluZTlmVy9SaERSUFN6Y20xYnNjMjRwbUFlaDhneSJ9LCJleHAiOjE1NDg5NTE2NzIsImlzcyI6ImVudmlvcy51c2VyIn0.Evyz1iaqBkqSjGxIhlXYMDcbrFX3G8yx_Svd6iAEgrE