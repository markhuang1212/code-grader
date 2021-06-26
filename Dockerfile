FROM ubuntu:21.04

RUN apt-get update
RUN apt-get -y upgrade

RUN apt-get -y install build-essential curl golang

RUN curl -fsSL https://deb.nodesource.com/setup_current.x | bash -
RUN apt-get install -y nodejs
RUN npm i -g npm@latest
RUN npm i -g yarn

WORKDIR /code-grader
COPY . .
RUN cd backend && make
RUN cd runtime && yarn install && yarn run build

USER daemon
CMD /code-grader/runtime/bin/code-grader