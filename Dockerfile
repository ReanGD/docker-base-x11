FROM debian:jessie

MAINTAINER ReanGD

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
                libgcc1 libstdc++6 libx11-6 libxcursor1 libxext6 libxrender1 && \
    apt-get install -y --no-install-recommends \
                ca-certificates wget less nano && \
    groupadd --gid 9999 docker && \
    useradd --password='' --uid=9999 --gid=docker --shell=/bin/bash --create-home docker

CMD /bin/bash
