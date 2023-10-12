FROM golang:1.21.1-alpine3.18 as builder

LABEL authors="yunusemremert"

WORKDIR /app

COPY . .

RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -o api-fiber-gorm main.go

FROM scratch

COPY --from=builder /app/api-fiber-gorm .

EXPOSE 3000

CMD ["./api-fiber-gorm"]