FROM golang:1.21.4-alpine AS builder

WORKDIR /build

COPY . .

ARG TARGETARCH

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=${TARGETARCH}

RUN go build -o cbef ./cmd/main.go

FROM alpine:3.18.4

COPY --from=builder /build/cbef .

CMD ["./cbef"]
