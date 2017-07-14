FROM alpine:latest

WORKDIR "/opt"

ADD .docker_build/go-heroku-math /opt/bin/go-heroku-math

CMD
