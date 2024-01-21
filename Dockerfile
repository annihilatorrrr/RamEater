FROM alpine:3.19.0
# FROM golang:1.21.6-alpine3.19 as builder
# RUN apk update && apk upgrade --available && sync && apk add --no-cache --virtual .build-deps upx
# WORKDIR /RamEater
# COPY main.sh /main.sh
# RUN chmod +x /main.sh
# CMD ["./main.sh"]
# COPY . .
# RUN go build -ldflags="-w -s" .
# RUN upx /RamEater/RamEater
# FROM alpine:3.19.0
# RUN apk update && apk upgrade --available && sync
# COPY --from=builder /RamEater/RamEater /RamEater

FROM python:3.12.1-alpine3.19
ENV VIRTUAL_ENV=/opt/venv
RUN python3 -m venv $VIRTUAL_ENV
WORKDIR /Eater
ENV PATH="$VIRTUAL_ENV/bin:$PATH" PYTHONUNBUFFERED=1 PIP_NO_CACHE_DIR=1
COPY . .
RUN apk update && apk upgrade --available && sync && apk add --no-cache --update --virtual .build-deps build-base && python3 -m pip install -U pip && pip3 install --no-cache-dir -U setuptools wheel && pip3 install --no-cache-dir -U -r requirements.txt && apk --purge del .build-deps && rm -rf /var/cache/apk/* && python3 -m compileall -b -o 2 . && rm -rf main.py requirements.txt /var/cache/apk/*
ENTRYPOINT ["python3", "main.pyc"]
