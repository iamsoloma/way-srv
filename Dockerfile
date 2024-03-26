FROM golang:1.22.1-alpine AS builder

WORKDIR /build

#Dependencies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

#Build
COPY . .
RUN go build -o ./bin/way-srv



FROM alpine as runner

WORKDIR /way

#WorkSpace
RUN mkdir blockchains

#Config
RUN mkdir metadata

#Copy
COPY --from=builder /build/bin/way-srv /way/way-srv


#Run
EXPOSE 1436-1436
CMD ["./way-srv"]