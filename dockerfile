FROM golang:1.20 AS builder


WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download


COPY . .


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fractal .


FROM alpine:latest

RUN apk add --no-cache libx11 libxrender libxext


COPY --from=builder /app/fractal .


CMD ["./fractal"]