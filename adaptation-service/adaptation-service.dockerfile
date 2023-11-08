FROM alpine:latest

RUN mkdir /app

COPY adaptationApp /app

CMD [ "/app/adaptationApp" ]