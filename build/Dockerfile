FROM golang:alpine AS builder

# RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /bin/app ./main.go

FROM alpine:latest
COPY --from=builder /bin/app /app

ENTRYPOINT ["./app","rest"]