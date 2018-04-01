FROM centos:7
MAINTAINER Matthias Wessendorf <matzew@apache.org>
ARG BINARY=./ghello
EXPOSE 7000

COPY ${BINARY} /opt/ghello
ENTRYPOINT ["/opt/ghello"]
