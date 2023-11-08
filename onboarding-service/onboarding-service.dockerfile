FROM alpine:latest

RUN mkdir /app

COPY onboardingApp /app

CMD [ "/app/onboardingApp" ]