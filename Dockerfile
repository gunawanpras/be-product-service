# Dockerfile definition for Backend application service.

# From which image we want to build. This is basically our environment.
FROM golang:1.23-alpine as Build

# This will copy all the files in our repo to the inside the container at root location.
COPY . .

# Copy the config file
COPY config.yaml /config.yaml

# Build our binary at root location.
RUN GOPATH= go build -o /main cmd/main.go

####################################################################
# This is the actual image that we will be using in production.
FROM alpine:latest

# We need to copy the binary from the build image to the production image.
COPY --from=Build /main .

# Copy the config file
COPY --from=Build /config.yaml .

# This is the port that our application will be listening on.
EXPOSE 1010

# This is the command that will be executed when the container is started.
ENTRYPOINT ["./main"]