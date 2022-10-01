FROM golang:1.16 AS base

FROM golang:1.16 AS build
WORKDIR /go/app
COPY . /go/app
RUN GO111MODULE=on  CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /go/app/main ./main.go

FROM base
COPY --from=build /go/app/main /main

CMD ["/main"]