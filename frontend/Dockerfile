FROM node:20 AS base

WORKDIR /app
COPY package.json .
RUN npm install

FROM base AS build

WORKDIR /app
COPY . .
RUN npx tsc || npx sass ./source/default-theme.scss ./source/default-theme.css

FROM build AS release

WORKDIR /app
EXPOSE 2022
CMD npx http-server . -p 2022