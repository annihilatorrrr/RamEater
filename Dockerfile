FROM alpine:3.17.2
WORKDIR /app
COPY . .
CMD ["./main.sh"]
