FROM golang:1.20.2-alpine3.17 as builder
WORKDIR /app
COPY . .
RUN go build -ldflags="-w -s" .
FROM scratch
COPY --from=builder /app/RamEater /
CMD ["/RamEater"]
