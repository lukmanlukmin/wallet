FROM golang:1.14-alpine

RUN apk update && \
    apk upgrade && \
    apk add bash git && \
    apk add gcc && \
    apk add musl-dev && \
    apk add curl && \
    apk add --update make

# We create an app directory within our
# image that will hold our application source
# files
RUN mkdir -p /home/go/src/wallet
# We copy everything in the root directory
# into our app directory
COPY . /home/go/src/wallet
# We specify that we now wish to execute 
# any further commands inside our /app
# directory
WORKDIR /home/go/src/wallet

ENV GO111MODULE=on

# use if it first time to register go module
# RUN go mod init github.com/lukmanlukmin/wallet

# install migrate-cli
RUN go get -tags 'postgres' -u github.com/golang-migrate/migrate/v4/cmd/migrate/

# run migration 
RUN migrate -verbose -source file://migration/postgresql -database postgres:'//wallet_user:wallet_password@174.21.210.11:5432/wallet_db?sslmode=disable' up

# build app to binary
# RUN go build -o main .

# RUN app on development mode
RUN go run main.go