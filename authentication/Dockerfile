FROM golang:1.9.2-alpine3.6

RUN apk --update add postgresql-client && rm -rf /var/cache/apk/*

ADD ./build.sh /build.sh
RUN chmod +x /build.sh

CMD /build.sh
