# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /url-short ./app

## Deploy
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /url-short /url-short
COPY --from=build /app/.env ./.env

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/url-short"]