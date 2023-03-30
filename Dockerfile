FROM golang:1.20.2-alpine3.17 as builder
RUN apk update && apk upgrade --available && sync
WORKDIR /RamEater
COPY . .
RUN go build -ldflags="-w -s" .
RUN rm -rf *.go && rm -rf go.* rm -rf *.c && rm -rf *.sh
FROM alpine:3.17.3
RUN apk update && apk upgrade --available && sync
COPY --from=builder /RamEater/RamEater /RamEater
CMD ["/RamEater"]
