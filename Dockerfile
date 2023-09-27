FROM golang:1.21.1-alpine3.18

LABEL authors="yunusemremert"

WORKDIR /app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build -o api-fiber-gorm .

EXPOSE 3000

CMD ["./api-fiber-gorm"]