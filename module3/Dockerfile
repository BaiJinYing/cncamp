FROM golang:1.17.8-alpine3.15
MAINTAINER BaiJinYing <baijinying@263.net>

ENV SRC_DIR /module2

ENV TINI_VERSION v0.19.0
RUN set -x \
  && cd /tmp \
  && wget -q -O tini https://github.com/krallin/tini/releases/download/$TINI_VERSION/tini \
  && chmod +x tini

COPY . $SRC_DIR

RUN cd $SRC_DIR \
  && GOPROXY="https://goproxy.cn" go build -o httpserver main.go

FROM alpine:3.15
MAINTAINER BaiJinYing <baijinying@263.net>

ENV SRC_DIR /module2
COPY --from=0 $SRC_DIR/httpserver /usr/local/bin/httpserver
COPY --from=0 /tmp/tini /sbin/tini

EXPOSE 8080

CMD ["/sbin/tini", "--"]