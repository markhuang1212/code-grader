FROM ubuntu:21.04

RUN apt-get update
RUN apt-get upgrade
RUN apt-get install -y build-essential curl
RUN curl -sL https://deb.nodesource.com/setup_16.x | bash
RUN apt-get install nodejs

COPY . /runtime
RUN yarn install

CMD node /runtime/dist/main.js
