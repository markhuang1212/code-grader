FROM ubuntu:21.04

RUN apt-get update
RUN apt-get -y upgrade

RUN apt-get -y install golang ca-certificates docker.io redis

RUN update-ca-certificates

WORKDIR /code-grader
COPY . .

WORKDIR /code-grader/backend
RUN go build

VOLUME [ "/var/run/docker.sock" ]
EXPOSE 8080
ENV GIN_MODE=release
ENV APP_ROOT=/code-grader

CMD [ "/code-grader/backend/backend" ]  