FROM golang:latest

LABEL maintainer "Quique<mayurkhairnar325@gmail.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8080

RUN go build

CMD ["./Dockerimage2"]