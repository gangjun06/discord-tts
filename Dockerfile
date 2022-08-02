## Build
FROM golang:1.18-alpine AS build

WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN apk add --update --no-cache pkgconfig opus-dev build-base

COPY . ./

RUN go build -o main .

## Deploy
FROM alpine

RUN apk add --no-cache sox ffmpeg

WORKDIR /

COPY --from=build /build/main /main
COPY  __files /__files

RUN mkdir __temp

CMD ["./main"]