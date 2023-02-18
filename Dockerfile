FROM golang:1-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY *.go ./

RUN go build -o /app/app

## Deploy
FROM alpine

COPY --from=build /app/app /app
RUN apk upgrade --no-cache && apk add --no-cache bash jq vim \
    && chmod +x /app
EXPOSE 4000
ENTRYPOINT ["/app"]
