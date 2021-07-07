FROM ubuntu:21.04

RUN apt-get update
RUN apt-get -y upgrade

WORKDIR /code-grader
COPY . .

VOLUME [ "/data" ]
WORKDIR /code-grader/runtime-exec

USER daemon
ENTRYPOINT [ "./run.sh" ]