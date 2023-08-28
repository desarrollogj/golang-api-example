FROM golang:1.20.7-alpine AS build_base

ARG ARTIFACT=example-api
ARG VERSION=0.0.1

# Set the Current Working Directory inside the container
WORKDIR /tmp/go-app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
RUN CGO_ENABLED=0 go test -v

# Swagger
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init

# Build the Go app
RUN go build -o ./out/$ARTIFACT-$VERSION .

# Start fresh from a smaller image
FROM alpine:3.17

# Install curl for health checks
RUN apk --no-cache add curl

ARG ARTIFACT=example-api
ARG VERSION=0.0.1
ARG PROFILE=DEVELOP
ARG PORT=9090

ENV APP_ARTIFACT=$ARTIFACT
ENV APP_VERSION=$VERSION
ENV APP_PROFILE=$PROFILE

COPY --from=build_base /tmp/go-app/out/$ARTIFACT-$VERSION /example-app
COPY --from=build_base /tmp/go-app/config /config

# This container exposes port to the outside world
EXPOSE $PORT

# Run the binary program produced by `go install`
CMD ["/example-app"]