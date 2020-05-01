# build stage
FROM golang as builder

ENV GO111MODULE=on

WORKDIR /workspace

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# final stage
FROM alpine
WORKDIR /app
COPY --from=builder /workspace/joke-server .
COPY --from=builder /workspace/config.yaml .
ENTRYPOINT ["/app/joke-server"]
