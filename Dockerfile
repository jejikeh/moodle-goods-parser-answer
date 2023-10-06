FROM golang:1.21-alpine

WORKDIR /app

ADD . /app

RUN go mod download

COPY *.go ./

RUN go build -o /main

EXPOSE 8080

CMD [ "/main" ]