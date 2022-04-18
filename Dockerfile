FROM golang:1.17.6-alpine3.15

RUN mkdir /app

WORKDIR /app

EXPOSE 8080

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

RUN go build

CMD ["go", "run", "."]