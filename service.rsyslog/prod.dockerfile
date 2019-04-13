# Based on https://hub.docker.com/r/voxxit/rsyslog/dockerfile

FROM alpine

RUN  apk add --update rsyslog \
  && rm -rf /var/cache/apk/*

EXPOSE 514
COPY ./service.rsyslog/rsyslog.conf /etc/rsyslog.conf
COPY ./service.rsyslog/log-rotation.sh /log-rotation.sh

CMD [ "rsyslogd", "-n" ]