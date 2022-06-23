FROM golang:1.18 as build-env

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...

RUN go build -o /go/bin/app

FROM heroku/heroku:20
COPY --from=build-env /go/bin/app /
COPY --from=build-env /go/src/app/config/config.yaml /config/config.yaml
EXPOSE 8080
CMD ["/app"]
