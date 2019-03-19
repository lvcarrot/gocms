FROM ubuntu
MAINTAINER Liping Wan <guyun_hy@163.com>
EXPOSE 8082
COPY ./gocms /data/bin/
COPY ./views /data/bin/views
COPY ./static /data/bin/static
ENV SERVICENAME="gocms"
ENTRYPOINT /data/bin/gocms
