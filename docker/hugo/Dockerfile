FROM alpine:3.7
LABEL description="Hugo container that builds and serve any hugo website."
LABEL maintainer="youenn@comptoirsecu.fr"
LABEL version="0.0.1"

ENV HUGO_VERSION=0.42.1
ENV HUGO_DOWNLOAD_URL=https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_${HUGO_VERSION}_Linux-64bit.tar.gz

## Install Hugo
RUN apk add --no-cache wget ca-certificates \
  && wget ${HUGO_DOWNLOAD_URL} \
  && tar xvzf ${HUGO_DOWNLOAD_URL##*/} \
  && rm ${HUGO_DOWNLOAD_URL##*/} \
  && mv hugo /usr/bin/hugo \
  && apk del wget ca-certificates

## Create user and file structure
RUN addgroup -g 1000 hugo \
  && adduser -u 1000 -G hugo -s /bin/sh -D hugo

USER hugo
WORKDIR /app
EXPOSE 1313/tcp

CMD ["hugo", "server", "--bind=0.0.0.0"]
