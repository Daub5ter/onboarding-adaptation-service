FROM node:16.20.2-alpine
#FROM node:16.20.2 AS node
#FROM php:7.4-fpm

#COPY --from=node /usr/local/lib/node_modules /usr/local/lib/node_modules
#COPY --from=node /usr/local/bin/node /usr/local/bin/node
#RUN ln -s /usr/local/lib/node_modules/npm/bin/npm-cli.js /usr/local/bin/npm

WORKDIR /app

COPY ./package.json .

RUN npm install

COPY . .

CMD [ "npm", "run", "dev" ]