FROM ubuntu:rolling

RUN apt-get update
RUN apt-get -y upgrade
RUN apt-get -y install build-essential

WORKDIR /code-grader
COPY . .

VOLUME [ "/data" ]
WORKDIR /code-grader/runtime-compile

USER daemon
ENTRYPOINT [ "./run.sh" ]