# FROM golang:1.24.1-alpine3.21 AS builder
# RUN apk update && apk upgrade --available && sync && apk add --no-cache --virtual .build-deps upx
# WORKDIR /RamEater
# COPY main.sh /main.sh
# RUN chmod +x /main.sh
# CMD ["./main.sh"]
# COPY . .
# RUN go build -trimpath -ldflags="-w -s" .
# RUN upx /RamEater/RamEater
# FROM scratch
# COPY --from=builder /RamEater/RamEater /RamEater
# ENTRYPOINT ["/RamEater"]

# FROM rust:1.82.0-alpine3.21 AS builder
# WORKDIR /Eater
# RUN apk update && apk upgrade --available && sync && apk add --no-cache --virtual .build-deps musl-dev libressl-dev build-base pkgconfig
# COPY . .
# RUN cargo build --release
# FROM alpine:3.20.3
# RUN apk update && apk upgrade --available && sync
# COPY --from=builder /Eater/target/release/Eater /Eater
# ENTRYPOINT ["/Eater"]

FROM python:3.13.3-alpine3.20
ENV VIRTUAL_ENV=/opt/venv
RUN python3 -m venv $VIRTUAL_ENV
WORKDIR /Eater
ENV PATH="$VIRTUAL_ENV/bin:$PATH" PYTHONUNBUFFERED=1 PIP_NO_CACHE_DIR=1
COPY . .
RUN apk update && apk upgrade --available && sync && apk add --no-cache --update --virtual .build-deps build-base gcc linux-headers python3-dev && python3 -m pip install -U pip && pip3 install --no-cache-dir -U setuptools wheel && pip3 install --no-cache-dir -U -r requirements.txt && apk --purge del .build-deps && rm -rf /var/cache/apk/* && python3 -m compileall -b -o 2 . && rm -rf main.py requirements.txt /var/cache/apk/*
ENTRYPOINT ["python3", "main.pyc"]
