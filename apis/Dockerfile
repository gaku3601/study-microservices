FROM alpine:3.6

RUN apk --no-cache add --virtual curl
ADD ./api.sh /api.sh

CMD sh api.sh
#CMD while sleep 3600; do :; done
