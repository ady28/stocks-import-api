FROM golang:1.18.2-bullseye AS builder

WORKDIR /stocks

COPY go.* ./

RUN go get -d ./...

COPY . .

RUN go build -x -a -tags netgo -installsuffix netgo ./*.go

#---

FROM alpine:3.17.2

ENV PORT=8080 \
    MONGODBSERVERNAME=stocksdb \
    MONGODBSERVERPORT=27017 \
    COMPETITORS_NAME=stockscompetitors \
    COMPETITORS_PORT=8080

EXPOSE $PORT

WORKDIR /app

COPY --from=builder /stocks/stocks .

USER nobody

CMD ["./stocks"]