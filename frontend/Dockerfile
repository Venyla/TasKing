FROM node:19 as buildenv

WORKDIR /app

RUN npm install typescript -g
COPY . .

# todo convert .ts to .js
#RUN tsc source/scripts.ts

FROM caddy:2.6

WORKDIR /app
COPY --from=buildenv /app .

# use --browse to see all files in directory
CMD caddy file-server --browse --listen :2022