FROM node:14-alpine AS build

COPY package.json yarn.lock /app/

WORKDIR /app

RUN yarn install --frozen-lockfile

COPY . .

RUN yarn build

FROM nginx:stable-alpine

COPY --from=build /app/build /usr/share/nginx/html

EXPOSE 80
