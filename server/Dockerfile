FROM golang:1.16 as base_image

# base image to build the go tree and ensure dependencies
FROM base_image as base
ENV GO111MODULE on
RUN mkdir -p /app/
WORKDIR /app/
ADD go.mod .
ADD go.sum .
RUN go mod download
ADD . .

# dev stage
FROM base as dev
RUN apt-get update && apt-get install -y curl git make ca-certificates bash apt-utils lsof
RUN GO111MODULE=off go get github.com/markbates/refresh
