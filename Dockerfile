FROM golang:1.22-alpine AS builder

WORKDIR /src

COPY app/go.mod ./
RUN go mod download

COPY app/ .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /app .

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /

COPY --from=builder /app /app

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app"]