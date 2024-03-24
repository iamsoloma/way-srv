FROM golang:1.22.1-alpine AS builder

WORKDIR /build

#Dependencies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

#Build
COPY . .
RUN go build -o ./bin/way-srv



FROM alpine as runner

#Copy
WORKDIR /way
COPY --from=builder /build/bin/way-srv /way/way-srv

#Config
COPY config.yaml .
ENV config=./config.yaml

EXPOSE 1436-1436

#Run
CMD ["./way-srv"]