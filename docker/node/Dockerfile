FROM node:10-alpine
LABEL description="Node container to install dependencies and build assets."
LABEL maintainer="youenn@comptoirsecu.fr"
LABEL version="0.0.1"

RUN apk add --no-cache git python make g++
# Needed for sharp: https://github.com/lovell/sharp/issues/354
RUN apk add vips-dev fftw-dev --update-cache --repository https://dl-3.alpinelinux.org/alpine/edge/testing/
RUN yarn global add gulp-cli

WORKDIR /app
CMD ["yarn", "watch"]
