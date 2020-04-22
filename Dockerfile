FROM golang:1.14-alpine

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
RUN mkdir -p /home/go/src/wallet
# We copy everything in the root directory
# into our /app directory
COPY . /home/go/src/wallet
# We specify that we now wish to execute 
# any further commands inside our /app
# directory
WORKDIR /home/go/src/wallet

ENV GO111MODULE=on
# add project maintainer
# RUN go get -u github.com/golang/dep/cmd/dep
# RUN dep init
# RUN dep ensure -v

RUN ls -la
# use if it first time to register go module
# RUN go mod init github.com/lukmanlukmin/wallet

# install migrate-cli
RUN go get -tags 'postgres' -u github.com/golang-migrate/migrate/v4/cmd/migrate/
# we run go build to compile the binary
# executable of our Go program
# RUN go build -o main .
# Our start command which kicks off
# our newly created binary executable
# CMD ["make","run"]