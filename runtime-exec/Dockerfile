FROM ubuntu:rolling

WORKDIR /code-grader
COPY . .

VOLUME [ "/data" ]
WORKDIR /code-grader/runtime-exec

USER daemon
ENTRYPOINT [ "./run.sh" ]