#build a golang image
FROM golang:1.20.1-alpine3.17 as builder
#set the working directory
WORKDIR /go/src/app
#copy the source code
COPY . .
#build the binary
RUN go build -o /go/bin/app

#build a scratch image
FROM scratch
#copy the binary from the builder image
COPY --from=builder /go/bin/app /go/bin/app
#set the entrypoint
ENTRYPOINT ["/go/bin/app"]