FROM golang:1.15.12-alpine3.13 as build
# Add Maintainer Info
LABEL maintainer="Luan Au <luanau@gmail.com>"
# Set the Current Working Directory inside the container
WORKDIR /app
## We copy everything in the root directory into our /app directory
COPY . .
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
## we run go build to compile the binary executable of our Go program
RUN go build -o main

FROM alpine
COPY --from=build /app/main .
# Expose port 8080 to the outside world
EXPOSE 8080
## Our start command which kicks off
## our newly created binary executable
CMD ["./main"]