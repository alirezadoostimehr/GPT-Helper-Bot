FROM golang:1.22.4-alpine as builder
WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .
RUN go build -v -o server


FROM alpine as runner
COPY --from=builder /app/server /app/server

ENTRYPOINT ["/app/server"]