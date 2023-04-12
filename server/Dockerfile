## We specify the base image we need for our
## go application
FROM golang:1.19.1-alpine3.16

## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /app

## We copy everything in the root directory
## into our /app directory
ADD . /app

## We specify that we now wish to execute
## any further commands inside our /app
## directory
WORKDIR /app

## Navigate to cmd folder
WORKDIR /app/cmd

## we run go build to compile the binary
## executable of our Go program
# Builds the application as a statically linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .

## Navigate back to the /app directory
WORKDIR /app

## Move the binary from the cmd folder to the /app directory
RUN mv /app/cmd/main /app/main

## Our start command which kicks off
## our newly created binary executable
CMD ["/app/main"]
