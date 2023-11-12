FROM golang:1.21.4-alpine AS builder

WORKDIR /build

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o cbef ./cmd/main.go

FROM scratch

COPY --from=builder /build/cbef .

CMD ["./cbef"]
