FROM ubuntu:21.04

RUN apt-get update
RUN apt-get -y upgrade

RUN apt-get -y install build-essential curl golang ca-certificates docker.io

RUN update-ca-certificates

WORKDIR /code-grader
COPY . .

EXPOSE 8080
ENV GIN_MODE=release

ENTRYPOINT [ "/code-grader/backend/backend" ]