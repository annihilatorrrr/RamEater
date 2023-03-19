FROM alpine3.17
WORKDIR /app
COPY . .
CMD ["/bin/bash", "./main.sh"]
