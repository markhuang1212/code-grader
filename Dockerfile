FROM ubuntu:21.04

RUN apt-get update
RUN apt-get -y upgrade

RUN apt-get -y install build-essential curl golang ca-certificates docker.io
RUN systemctl enable docker.io
RUN systemctl start docker.io

RUN update-ca-certificates
RUN usermod -aG docker daemon

WORKDIR /code-grader
COPY . .

RUN docker build . -f runtime-compile/Dockerfile -t runtime-compile
RUN docker build . -f runtime-exec/Dockerfile -t runtime-exec

EXPOSE 8080

USER daemon
ENTRYPOINT [ "/code-grader/backend/code-grader" ]