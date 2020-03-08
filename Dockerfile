FROM alpine:3.10.1
LABEL maintainer="Haruai Tamada" \
      uniq2-version="1.0.1" \
      description="heatmap generator"

RUN    adduser -D heatman \
    && apk --no-cache add curl=7.66.0-r0 tar=1.32-r0 \
    && curl -s -L -O https://github.com/tamada/goheatman/releases/download/v1.0.1/heatman-1.0.1_linux_amd64.tar.gz \
    && tar xfz heatman-1.0.1_linux_amd64.tar.gz          \
    && mv heatman-1.0.1 /opt                             \
    && ln -s /opt/heatman-1.0.1 /opt/heatman             \
    && ln -s /opt/heatman/heatman /usr/local/bin/heatman \
    && rm heatman-1.0.1_linux_amd64.tar.gz

ENV HOME="home/heatman"

WORKDIR /home/heatman
USER    heatman

ENTRYPOINT [ "heatman" ]
