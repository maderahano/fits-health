FROM golang:1.17 as builder

WORKDIR /app

COPY . .

RUN go build -tags netgo -o main.app ./cmd/app/main.go

FROM alpine:latest

WORKDIR /pack

COPY --from=builder /app/main.app ./cmd/app/main.go

CMD [ "/pack/main.app" ]