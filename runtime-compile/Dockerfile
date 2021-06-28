FROM ubuntu:21.04

RUN apt-get update
RUN apt-get -y upgrade
RUN apt-get -y install golang

WORKDIR /code-grader
COPY ../ .

WORKDIR /code-grader/runtime-compile
RUN go build

ENTRYPOINT [ "./code-grader-runtime-compile" ]