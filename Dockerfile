FROM alpine:latest

COPY ./go-fullstack ./app

ENTRYPOINT [ "./app" ]