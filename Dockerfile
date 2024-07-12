FROM golang:1.22.5-alpine

WORKDIR /usr/src/app

COPY . .

RUN mkdir -p /usr/local/bin
RUN go mod tidy
RUN go build -v -o /usr/local/bin/app cmd/main.go

EXPOSE 8000

CMD ["app"]