FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/mainnode    ./cmd/mainnode
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/datanode ./cmd/datanode

FROM alpine:latest AS mainnode-runtime
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# copy only api binary into this image
COPY --from=builder /app/bin/mainnode ./bin/mainnode
EXPOSE 8080
CMD ["./bin/mainnode"]

FROM alpine:latest AS datanode-runtime
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# copy only api binary into this image
COPY --from=builder /app/bin/datanode ./bin/datanode
CMD ["./bin/datanode"]