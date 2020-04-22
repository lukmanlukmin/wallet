FROM golang:1.12-alpine

RUN apk update && \
    apk upgrade && \
    apk add bash git && \
    apk add gcc && \
    apk add musl-dev && \
    apk add curl && \
    apk add --update make

# We create an /app directory within our
# image that will hold our application source
# files
RUN mkdir /go/src/wallet
# We copy everything in the root directory
# into our /app directory
COPY . /go/src/wallet
# We specify that we now wish to execute 
# any further commands inside our /app
# directory
WORKDIR /go/src/wallet
# change directory to our working directory
RUN cd /go/src/wallet
# add project maintainer
RUN go get -u github.com/golang/dep/cmd/dep
# RUN dep init
# RUN dep ensure
# we run go build to compile the binary
# executable of our Go program
# RUN go build -o main .
# Our start command which kicks off
# our newly created binary executable
# CMD ["make","run"]