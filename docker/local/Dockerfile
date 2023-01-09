FROM golang:1.18-alpine AS builder

WORKDIR /gitshock

ADD go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o gitshock main.go

FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates jq

WORKDIR /gitshock

COPY --from=builder /gitshock/gitshock ./
COPY ./docker/local/gitshock.sh ./

# Expose json-rpc, libp2p and grpc ports
EXPOSE 8545 9632 1478

ENTRYPOINT ["./gitshock.sh"]
