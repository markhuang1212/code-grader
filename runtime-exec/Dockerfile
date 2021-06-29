FROM ubuntu:21.04

RUN apt-get update
RUN apt-get -y upgrade
RUN apt-get -y install golang ca-certificates build-essential

RUN update-ca-certificates

WORKDIR /code-grader
COPY . .

WORKDIR /code-grader/runtime-exec
RUN go build
ENTRYPOINT [ "./runtime-exec" ]