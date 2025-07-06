FROM node:24-alpine

RUN npm install -g @typespec/compiler

WORKDIR /app

COPY . /app
