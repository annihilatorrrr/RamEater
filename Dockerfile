FROM golang:1.20.6-alpine3.18 as builder
RUN apk update && apk upgrade --available && sync && apk add --no-cache --virtual .build-deps upx
WORKDIR /RamEater
COPY . .
RUN go build -ldflags="-w -s" .
RUN upx /RamEater/RamEater
FROM alpine:3.18.2
RUN apk update && apk upgrade --available && sync
COPY --from=builder /RamEater/RamEater /RamEater
ENTRYPOINT ["/RamEater"]
