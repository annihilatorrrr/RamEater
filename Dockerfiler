FROM golang:1.20.4-alpine3.18 as builder
RUN apk update && apk upgrade --available && sync
WORKDIR /RamEater
COPY . .
RUN go build -ldflags="-w -s" .
RUN rm -rf *.go && rm -rf go.* rm -rf *.c && rm -rf *.sh
FROM alpine:3.18.0
RUN apk update && apk upgrade --available && sync
COPY --from=builder /RamEater/RamEater /RamEater
CMD ["/RamEater"]
