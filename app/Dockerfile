FROM golang:1.15-alpine AS build
# Support CGO and SSL
RUN apk --no-cache add gcc g++ make
RUN apk add git
WORKDIR /go/src/app
COPY . .
# RUN go mod install
RUN GOOS=linux go build -ldflags="-s -w" -o ./main ./main.go

FROM alpine:3.10
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/app/main /usr/bin/main

CMD ./main